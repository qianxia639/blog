package utils

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func De(input []byte) string {
	output := blackfriday.Run(input, blackfriday.WithNoExtensions())
	return string(bluemonday.UGCPolicy().SanitizeBytes(output))
}
