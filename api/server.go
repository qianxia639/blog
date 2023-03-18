package api

import (
	db "Blog/db/sqlc"
	"Blog/token"
	"Blog/utils/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	conf   config.Config
	router *gin.Engine
	maker  token.Maker
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

	router.POST("/user", server.createUser)
	router.POST("/login", server.login)

	router.PUT("/blog/inrc", server.incrViews)
	router.GET("/blog", server.listBlogs)
	router.GET("/blog/:id", server.getBlog)

	authRouter := router.Group("/").Use(server.authMiddlware(server.maker))
	{
		authRouter.PUT("/user", server.updateUser)

		authRouter.POST("/type", server.insertType)

		authRouter.POST("/blog", server.insertBlog)
		authRouter.DELETE("/blog/:id", server.deleteBlog)
	}

	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
