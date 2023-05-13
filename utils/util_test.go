package utils

import "testing"

func TestExt(t *testing.T) {
	ext := Ext("/opt/blog/bg.jpeg")
	t.Logf("ext: %s\n", ext)
}
