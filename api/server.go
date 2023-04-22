package api

import (
	"Blog/core/config"
	"Blog/core/logs"
	db "Blog/db/sqlc"
	"Blog/token"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	store  db.Store
	conf   config.Config
	router *gin.Engine
	maker  token.Maker
	rdb    *redis.Client
}

type ServerOptions func(*Server)

func WithStor(store db.Store) ServerOptions {
	return func(s *Server) {
		s.store = store
	}
}

func WithConfig(conf config.Config) ServerOptions {
	return func(s *Server) {
		s.conf = conf
	}
}

func WithMaker(maker token.Maker) ServerOptions {
	return func(s *Server) {
		s.maker = maker
	}
}

func WithCache(rdb *redis.Client) ServerOptions {
	return func(s *Server) {
		s.rdb = rdb
	}
}

func NewServerV2(opts ...ServerOptions) *Server {
	server := &Server{}
	for _, opt := range opts {
		opt(server)
	}

	server.setupRouter()

	return server
}

func NewServer(conf config.Config, store db.Store) (*Server, error) {

	maker, err := token.NewPasetoMaker(conf.Token.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		conf:  conf,
		store: store,
		maker: maker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(CORS()).Use(LogFuncExecTime())

	router.POST("/user", server.createUser)
	router.POST("/login", server.login)

	router.PUT("/blog/incr/:id", server.incrViews)
	router.GET("/blog", server.listBlogs)
	router.GET("/blog/:id", server.getBlog)
	router.GET("/blog/search", server.searchBlog)

	router.POST("/comment", server.createComment)
	router.GET("/comment", server.getComments)

	authRouter := router.Group("/").Use(server.authMiddlware())
	{
		authRouter.PUT("/user", server.updateUser)

		authRouter.POST("/blog", server.insertBlog)
		authRouter.DELETE("/blog/:id", server.deleteBlog)
		authRouter.PUT("/blog", server.updateBlog)
	}

	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func LogFuncExecTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		method := ctx.Request.Method
		ip := ctx.ClientIP()

		if raw != "" {
			path += raw
		}

		statusCode := ctx.Writer.Status()

		latency := time.Since(start).Milliseconds()

		// time | statusCode | timeSub | ip | method | path
		logs.Logs.Infof("%s | %d | %5dms | %s | %s | %s",
			start.Format("2006/01/02 15:04:05"), statusCode, latency, ip, method, path,
		)
	}
}
