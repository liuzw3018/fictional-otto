package main

import (
	"github.com/liuzw3018/otto/server/core"
	"github.com/liuzw3018/saber/lib"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author: liu zw
 * @Date: 2022/9/22 11:48
 * @File: main.go
 * @Description: //TODO $
 * @Version:
 */

func main() {
	if err := lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"}); err != nil {
		panic(err)
	}
	defer lib.Destroy()

	core.RunServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	core.StopServer()
}
