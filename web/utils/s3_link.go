package utils

import (
	"strings"

	"RouteHub.Service.Dashboard/features/configuration"
)

func linkToS3Path(link string) string {
	staticConfig := configuration.Get().GetStaticConfig()
	if staticConfig == nil {
		return link
	}
	s3Config := staticConfig.S3
	return strings.Join([]string{"https:/", s3Config.BucketDomain, link}, "/")
}

func S3AppendAble(link string) bool {
	// check it if starts with 'http://' or 'https://'
	if link == "" {
		return false
	} else if len(link) < 8 {
		return true
	}

	return link[:7] != "http://" && link[:8] != "https://"
}

func LinkToS3Path(link string) string {
	if S3AppendAble(link) {
		return linkToS3Path(link)
	}
	return link
}
