package extensions

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/features/fileUpload"
	"github.com/labstack/echo/v4"
)

func ProcessFile(ctx context.Context, filePart *multipart.FileHeader, path string, fileName string) (uploadedBucketPath string, hasFile bool, err error) {
	if filePart != nil && filePart.Size > 0 {
		hasFile = true

		s3ClientService, err := fileUpload.GetS3ClientService()
		if err != nil {
			return "", hasFile, err
		}

		staticConfig := configuration.Get().GetStaticConfig()
		if staticConfig == nil {
			return "", hasFile, fmt.Errorf("S3 configuration is not set")
		}

		fileLink := strings.Join([]string{path, fileName}, "/")

		uploadedPath, err := s3ClientService.UploadFormFileThroughS3(ctx, *filePart, fileLink)
		if err != nil {
			return "", hasFile, err
		}

		return uploadedPath, hasFile, nil
	}

	return "", hasFile, nil
}

func ProcessFileFromEchoContext(c echo.Context, field *string, formKey string, path string, fileName string) error {
	if file, err := c.FormFile(formKey); err == nil {
		ctx := c.Request().Context()

		if fileName == "" {
			fileName = file.Filename
			ext := filepath.Ext(fileName)

			timeStamp := time.Now().Unix()
			fileName = fmt.Sprintf("file_%d%s", timeStamp, ext)
		} else {
			if filepath.Ext(fileName) == "" {
				fileName = fmt.Sprintf("%s%s", fileName, filepath.Ext(file.Filename))
			}
		}

		filePath, hasFile, err := ProcessFile(ctx, file, path, fileName)
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
