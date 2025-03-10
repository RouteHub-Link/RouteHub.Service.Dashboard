package fileUpload

import (
	"context"
	"errors"
	"mime/multipart"
	"sync"

	"RouteHub.Service.Dashboard/features/configuration"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	client             *s3.Client
	once               sync.Once
	ConfigurationError error
)

type S3ClientService struct{ configuration configuration.S3ClientConfig }

func GetS3ClientService() (service *S3ClientService, err error) {
	sc, err := configuration.Get().GetStaticConfigX()
	if err != nil {
		return nil, err
	}

	cc := sc.S3

	if isValid, err := cc.Validate(); !isValid {
		return nil, err
	}

	once.Do(func() {
		resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: cc.DefaultEndpoint,
			}, nil
		})

		cfg, err := awsConfig.LoadDefaultConfig(
			context.Background(),
			awsConfig.WithClientLogMode(aws.LogRequestWithBody|aws.LogResponseWithBody),
			awsConfig.WithRegion("auto"),
			awsConfig.WithEndpointResolverWithOptions(resolver),
			awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cc.PublicAccessKey, cc.SecretAccessKey, "")))

		if err != nil {
			ConfigurationError = err
			return
		}

		client = s3.NewFromConfig(cfg)
		service = &S3ClientService{
			configuration: cc,
		}
	})

	return service, ConfigurationError

}

func (cs S3ClientService) GetClient() (*s3.Client, error) {
	if ConfigurationError != nil {
		return nil, ConfigurationError
	}

	if client == nil {
		return nil, errors.New("S3 client is not initialized")
	}

	return client, nil
}

func (cs S3ClientService) UploadFormFileThroughS3(ctx context.Context, fileHeader multipart.FileHeader, objectPath string) (filePath string, err error) {
	client, err := cs.GetClient()

	if err != nil {
		return "", err
	}

	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(cs.configuration.BucketName),
		Key:         aws.String(objectPath),
		Body:        src,
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})

	if err != nil {
		return "", err
	}

	return objectPath, nil
}
