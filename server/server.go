package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/routers"
)

func Run() {
	router := routers.Router()

	addr := fmt.Sprintf("%s:%d", global.QX_CONFIG.Http.Host, global.QX_CONFIG.Http.Port)
	s := server(addr, router)
	global.QX_LOG.Infof("server run success [%s]", addr)
	global.QX_LOG.Error(s.ListenAndServe().Error())

}

func server(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
