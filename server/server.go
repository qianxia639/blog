package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/qianxia/blog/global"
)

func Server(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", global.QX_CONFIG.Http.Host, global.QX_CONFIG.Http.Port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
