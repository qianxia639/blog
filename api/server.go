package api

import (
	"Blog/core/config"
	"Blog/core/task"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"Blog/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	store           db.Store
	conf            *config.Config
	router          *gin.Engine
	maker           token.Maker
	rdb             *redis.Client
	taskDistributor task.TaskDistributor
}

func NewServer() *Server {
	server := &Server{}

	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.CORS())

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

	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func (server *Server) WithStore(store db.Store) *Server {
	server.store = store
	return server
}

func (server *Server) WithConfig(conf *config.Config) *Server {
	server.conf = conf
	return server
}

func (server *Server) WithMaker(maker token.Maker) *Server {
	server.maker = maker
	return server
}

func (server *Server) WithCache(rdb *redis.Client) *Server {
	server.rdb = rdb
	return server
}

func (server *Server) WithTaskDistributor(taskDistributor task.TaskDistributor) *Server {
	server.taskDistributor = taskDistributor
	return server
}
