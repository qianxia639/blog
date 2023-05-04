package api

import (
	"Blog/core/config"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"Blog/middleware"

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

func NewServer(opts ...ServerOptions) *Server {
	server := &Server{}
	for _, opt := range opts {
		opt(server)
	}

	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(middleware.CORS()).Use(middleware.LogFuncExecTime())

	router.POST("/user", server.createUser)
	router.POST("/login", server.login)

	router.PUT("/article/incr/:id", server.incrViews)
	router.GET("/article", server.listBlogs)
	router.GET("/article/:id", server.getArticle)
	router.GET("/article/search", server.searchArticle)

	router.POST("/comment", server.createComment)
	router.GET("/comment", server.getComments)

	authRouter := router.Group("/").Use(middleware.Authorization(server.maker, server.rdb))
	{
		authRouter.PUT("/user", server.updateUser)

		authRouter.POST("/article", server.insertArticle)
		authRouter.DELETE("/article", server.deleteArticle)
		authRouter.PUT("/article", server.updateArticle)
	}

	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func (server Server) GetRoutr() *gin.Engine {
	return server.router
}

func (server Server) GetMaker() token.Maker {
	return server.maker
}
