package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHostname(t *testing.T) {
	hostname := getHostname()
	require.NotEmpty(t, hostname)
}
