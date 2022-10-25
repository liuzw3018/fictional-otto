package core

import (
	"context"
	"fmt"
	"github.com/liuzw3018/otto/server/initialize"
	"github.com/liuzw3018/otto/server/pkg/global"
	_ "github.com/liuzw3018/otto/server/pkg/global"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/9/22 11:48
 * @File: server.go
 * @Description: //TODO $
 * @Version:
 */

func RunServer() {
	app := initialize.Routers()
	address := fmt.Sprintf("%s", global.OttoConfig.Server.Listen)

	s := initHttpServer(address, app)
	go func() {
		global.OttoLogger.Infof("服务器已启动，监听地址 %s", address)
		global.OttoLogger.Error(s.ListenAndServe())
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.OttoLogger.Warningln("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.OttoLogger.Fatalln("Server Shutdown:", err)
	}
	global.OttoLogger.Warningln("Server exiting")
}
