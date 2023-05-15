package utils

import "testing"

func TestExt(t *testing.T) {
	ext := Ext("/opt/blog/sqlc.yaml")
	t.Logf("ext: %s\n", ext)
}
