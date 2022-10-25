package initialize

import (
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/liuzw3018/otto/server/pkg/zaplog"
	"go.uber.org/zap"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 11:09
 * @File: sys_zap.go
 * @Description: //TODO $
 * @Version:
 */

func initZap() {
	data := &zaplog.Options{
		LogFileDir: "/Users/liuzhineng/GolandProjects/otto/logs",
		AppName:    "otto",
		MaxSize:    30,
		MaxBackups: 7,
		MaxAge:     7,
		Config:     zap.Config{},
	}
	// dev模式
	data.Development = true
	// prod模式
	//data.Development = false

	global.OttoZap = zaplog.NewZapLogger(data)
}
