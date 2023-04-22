package api

import (
	"Blog/core/config"
	"Blog/core/logs"
	db "Blog/db/sqlc"
	"Blog/utils"
	"database/sql"
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

	server.router.Use(CORS(), LogFuncExecTime())

	logs.Logs = logs.InitZap(&conf)

	return server
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

type TestServer struct {
	store db.Store
}

type Option func(*TestServer)

func InitOption(opts ...Option) *TestServer {
	ts := &TestServer{}
	for _, opt := range opts {
		opt(ts)
	}
	return ts
}

func WithStore(store db.Store) Option {
	return func(ts *TestServer) {
		ts.store = store
	}
}

// func WithCache() Option {
// 	return func(ts *TestServer) {

// 	}
// }
