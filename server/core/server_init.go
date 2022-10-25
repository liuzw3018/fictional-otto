package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/9/22 11:54
 * @File: server_init.go
 * @Description: //TODO $
 * @Version:
 */

func initHttpServer(address string, handlerFunc *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        handlerFunc,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
