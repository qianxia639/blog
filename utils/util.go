package utils

import (
	"net/url"
	"path/filepath"
)

func Ext(localfile string) string {
	return filepath.Ext(localfile)
}

func InValidUrl(u string) bool {
	ul, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	ur, err := ul.Parse(u)
	if err != nil || ur.Scheme == "" || ur.Host == "" {
		return false
	}

	if ur.Scheme != "http" && ur.Scheme != "https" {
		return false
	}

	return true
}

const (
	// image max 3M
	ImageMaxSize = 3072 * 1024 // 单位: byte
)
