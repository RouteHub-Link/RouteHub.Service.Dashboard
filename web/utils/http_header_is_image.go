package utils

import (
	"errors"
	"strings"
)

// List of web-embeddable image MIME types and their extensions
var WebEmbeddableImages = map[string][]string{
	"image/avif":    {".avif"},
	"image/bmp":     {".bmp"},
	"image/gif":     {".gif"},
	"image/jpeg":    {".jpe", ".jpeg", ".jpg"},
	"image/png":     {".png"},
	"image/svg+xml": {".svg", ".svgz"},
	"image/tiff":    {".tif", ".tiff"},
	"image/webp":    {".webp"},
	"image/x-icon":  {".ico"},
}

// CheckHeaderIsImage checks if the given Content-Type header indicates a web-embeddable image
func CheckHeaderIsImage(contentType string) (bool, error) {
	if contentType == "" {
		return false, errors.New("empty content type")
	}

	// Normalize Content-Type (remove parameters like charset)
	contentType = strings.Split(contentType, ";")[0]

	// Check if the Content-Type exists in the list of embeddable images
	if _, exists := WebEmbeddableImages[contentType]; exists {
		return true, nil
	}

	return false, nil
}
