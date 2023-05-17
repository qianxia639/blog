package utils

import (
	"sync"
)

// Image Type
var imageTypes = map[string]string{
	"image/avif":    ".avif",
	"image/jpeg":    ".jpeg",
	"image/jpg":     ".jpg",
	"image/png":     ".png",
	"image/svg+xml": ".svg",
	"image/webp":    ".webp",
}

var once sync.Once

var imageMime = make(map[string]string, len(imageTypes))

func GetInstance() map[string]string {
	once.Do(func() {
		imageMime = imageTypes
	})
	return imageMime
}
