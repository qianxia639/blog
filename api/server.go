package api

import (
	"Blog/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	maker  token.Maker
}

const key = ""

func NewServer() (*Server, error) {

	maker, err := token.NewPasetoMaker(key)
	if err != nil {
		return nil, err
	}

	server := &Server{
		maker: maker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/usre", server.createUser)
	router.POST("/login", server.login)

	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
