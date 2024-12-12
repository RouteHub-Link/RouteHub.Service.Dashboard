package extensions

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func ProcessFile(ctx context.Context, filePart *multipart.FileHeader, bucket string, path string, fileName string) (uploadedBucketPath string, hasFile bool, err error) {
	if filePart != nil && filePart.Size > 0 {
		hasFile = true
		favLink := FileBucketLocation(bucket, path, fileName)

		uploadedPath, err := cr2Client.UploadFormFileThroughR2(ctx, *filePart, favLink)
		if err != nil {
			return "", hasFile, err
		}

		return uploadedPath, hasFile, nil
	}

	return "", hasFile, nil
}

func ProcessFileFromEchoContext(c echo.Context, field *string, formKey string, bucket string, path string, fileName string) error {
	if file, err := c.FormFile(formKey); err == nil {
		ctx := c.Request().Context()
		filePath, hasFile, err := ProcessFile(ctx, file, bucket, path, fileName)
		if filePath != "" && err == nil && hasFile {
			log.Printf("Uploaded successfully: %s", filePath)
			*field = filePath
			return nil
		} else if err != nil && hasFile {
			return err
		}
	}

	return nil
}

func FileBucketLocation(bucket string, path string, fileName string) string {
	filePath := filepath.Ext(fileName)

	return fmt.Sprintf("%s/%s/%s%s", bucket, path, fileName, filePath)
}
