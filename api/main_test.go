package api

import (
	"Blog/core/cache"
	"Blog/core/config"
	"Blog/core/logs"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"Blog/middleware"
	"Blog/utils"
	"database/sql"
	"net/http"
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

	conf.Redis.Address = "localhost:6379"

	cache := cache.InitRedis(conf)

	maker := newTokenMaker(t, conf.Token.TokenSymmetricKey)

	opts := []ServerOptions{
		WithConfig(conf),
		WithStor(store),
		WithCache(cache),
		WithMaker(maker),
	}

	server := NewServer(opts...)

	server.router.Use(middleware.CORS(), middleware.LogFuncExecTime())

	logs.Logs = logs.InitZap(&conf)

	return server
}

func newTokenMaker(t *testing.T, symmetriKey string) token.Maker {
	maker, err := token.NewPasetoMaker(symmetriKey)
	require.NoError(t, err)

	return maker
}

func addAuthorizatin(t *testing.T, req *http.Request, tokenMaker token.Maker,
	username string, duration time.Duration) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	req.Header.Set("Authorization", token)
}

func newTestDB(t *testing.T) db.Store {

	conf, err := config.LoadConfig("..")
	require.NoError(t, err)

	conn, err := sql.Open(conf.Postgres.Driver, conf.Postgres.Source)
	require.NoError(t, err)

	return db.NewStore(conn)

}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
