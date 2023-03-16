package api

import (
	db "Blog/db/sqlc"
	"Blog/utils"
	"Blog/utils/config"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	conf := config.Config{}

	conf.Token.TokenSymmetricKey = utils.RandomString(32)
	conf.Token.AccessTokenDuration = time.Minute

	server, err := NewServer(conf, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
