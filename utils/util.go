package utils

import (
	"path/filepath"
)

func Ext(localfile string) string {
	return filepath.Ext(localfile)
}
