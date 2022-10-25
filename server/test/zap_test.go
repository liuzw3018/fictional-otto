package test

import (
	"fmt"
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/liuzw3018/otto/server/pkg/zaplog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 11:02
 * @File: zap_test.go
 * @Description: //TODO $
 * @Version:
 */

func TestZap(t *testing.T) {
	data := &zaplog.Options{
		LogFileDir: "/Users/liuzhineng/GolandProjects/otto/logs",
		AppName:    "otto",
		MaxSize:    30,
		MaxBackups: 7,
		MaxAge:     7,
		Config:     zap.Config{},
	}
	data.Development = true
	//zaplog.InitLogger(data)

	global.OttoZap = zaplog.NewZapLogger(data)
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		global.OttoZap.Debug(fmt.Sprint("debug log ", i), zap.Int("line", 999))
		global.OttoZap.Info(fmt.Sprint("Info log ", i), zap.Any("level", "1231231231"))
		global.OttoZap.Warn(fmt.Sprint("warn log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		global.OttoZap.Error(fmt.Sprint("err log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}

}

func TestLogger(t *testing.T) {
	data := &zaplog.Options{
		LogFileDir: "/Users/liuzhineng/GolandProjects/otto/logs",
		AppName:    "otto",
		MaxSize:    30,
		MaxBackups: 7,
		MaxAge:     7,
		Config:     zap.Config{},
	}
	data.Development = true
	//zaplog.InitLogger(data)
	global.OttoZap = zaplog.NewZapLogger(data)
	global.OttoZap.Opts.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l := global.NewLogger(global.OttoZap)
	l.Info("12345")
}
