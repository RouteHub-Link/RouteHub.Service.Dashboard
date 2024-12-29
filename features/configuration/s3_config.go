package configuration

import (
	"fmt"
)

type S3ClientConfig struct {
	API             string `env:"S3_API" envDefault:"" form:"api" validate:"required,url"`
	BucketDomain    string `env:"S3_BUCKET_DOMAIN" envDefault:"" form:"bucket_domain" validate:"required"`
	BucketName      string `env:"S3_BUCKET_NAME" envDefault:"" form:"bucket_name" validate:"required"`
	DefaultEndpoint string `env:"S3_DEFAULT_ENDPOINT" envDefault:"" form:"default_endpoint" validate:"required,url"`
	EUEndpoint      string `env:"S3_EU_ENDPOINT" envDefault:"" form:"eu_endpoint" validate:"required,url""`
	SecretAccessKey string `env:"S3_SECRET_ACCESS" envDefault:"" form:"secret_access_key" validate:"required"`
	PublicAccessKey string `env:"S3_PUBLIC_ACCESS" envDefault:"" form:"public_access_key" validate:"required"`
}

func (c *S3ClientConfig) String() string {
	return fmt.Sprintf("API: %s, BucketDomain: %s, BucketName: %s, DefaultEndpoint: %s, EUEndpoint: %s, SecretAccessKey: %s, PublicAccessKey: %s", c.API, c.BucketDomain, c.BucketName, c.DefaultEndpoint, c.EUEndpoint, c.SecretAccessKey, c.PublicAccessKey)
}

func (c *S3ClientConfig) Validate() (isValid bool, err error) {
	if c.API == "" {
		err = fmt.Errorf("S3 API is required")
	} else if c.BucketDomain == "" {
		err = fmt.Errorf("S3 BucketDomain is required")
	} else if c.BucketName == "" {
		err = fmt.Errorf("S3 BucketName is required")
	} else if c.DefaultEndpoint == "" {
		err = fmt.Errorf("S3 DefaultEndpoint is required")
	} else if c.EUEndpoint == "" {
		err = fmt.Errorf("S3 EUEndpoint is required")
	} else if c.SecretAccessKey == "" {
		err = fmt.Errorf("S3 SecretAccessKey is required")
	} else if c.PublicAccessKey == "" {
		err = fmt.Errorf("S3 PublicAccessKey is required")
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *S3ClientConfig) IsEmpty() bool {
	return c.API == "" && c.BucketDomain == "" && c.BucketName == "" && c.DefaultEndpoint == "" && c.EUEndpoint == "" && c.SecretAccessKey == "" && c.PublicAccessKey == ""
}
