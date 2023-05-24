package utils

import (
	"sync"
)

type M map[string]string

// Image Type
var imageTypes = M{
	"image/avif":    ".avif",
	"image/jpeg":    ".jpeg",
	"image/jpg":     ".jpg",
	"image/png":     ".png",
	"image/svg+xml": ".svg",
	"image/webp":    ".webp",
}

var once sync.Once

var imageMime = make(M, len(imageTypes))

func GetInstance() M {
	once.Do(func() {
		imageMime = imageTypes
	})
	return imageMime
}

func GetInstance2() sync.Map {
	once.Do(initImageMime)

	return extensions
}

func GetInstance3() sync.Map {
	once.Do(func() {
		setImageMime(imageTypes)
	})

	return extensions
}

var testInitMime func()

func initImageMime() {
	if fn := testInitMime; fn != nil {
		fn()
	} else {
		setImageMime(imageTypes)
	}
}

var extensions sync.Map
var extensionsMutex sync.Mutex

func clearSyncMap(m *sync.Map) {
	m.Range(func(key, _ any) bool {
		m.Delete(key)
		return true
	})
}

func setImageMime(imaggeTypes M) {
	clearSyncMap(&extensions)

	extensionsMutex.Lock()
	defer extensionsMutex.Unlock()
	for k, v := range imaggeTypes {
		extensions.Store(k, v)
	}
}
