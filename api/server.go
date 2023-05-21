package api

import (
	"Blog/core/config"
	"Blog/core/task"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"Blog/middleware"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	store           db.Store
	conf            config.Config
	router          *gin.Engine
	maker           token.Maker
	rdb             *redis.Client
	taskDistributor task.TaskDistributor
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

func WithTaskDistributor(taskDistributor task.TaskDistributor) ServerOptions {
	return func(s *Server) {
		s.taskDistributor = taskDistributor
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	limiter := redis_rate.NewLimiter(server.rdb)

	router.Use(middleware.CORS()).Use(middleware.Limit(limiter))

	router.POST("/user", server.createUser)
	router.POST("/login", server.login)
	router.POST("/refresh", server.refreshToken)

	router.PUT("/article/incr/:id", server.incrViews)
	router.GET("/article", server.listArticles)
	router.GET("/article/:id", server.getArticle)

	router.POST("/critique", server.createCritique)
	router.GET("/critique", server.getCritiques)

	authRouter := router.Group("/").Use(middleware.Authorization(server.maker, server.rdb))
	{
		authRouter.GET("/user", server.getUser)
		authRouter.PUT("/user", server.updateUser)
		authRouter.POST("/logout", server.logout)

		authRouter.POST("/article", server.insertArticle)
		authRouter.DELETE("/article", server.deleteArticle)
		authRouter.PUT("/article", server.updateArticle)
	}

	ginpprof.Wrap(router)

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
