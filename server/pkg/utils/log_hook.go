package utils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 18:18
 * @File: log_hook.go
 * @Description: //TODO $
 * @Version:
 */

func NewLfsHook(logLevel, logSplit string, maxRemainCnt uint) logrus.Hook {
	var (
		err      error
		duration time.Duration
		logName  = "logs/log"
	)

	duration, err = ParseDuration(logSplit)
	if err != nil {
		logrus.Infoln("请检查配置文件日志切割时间，正确格式为 1d | 1h | 1m | 1.5h等")
		duration = time.Hour
	}
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(duration),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	level, err := logrus.ParseLevel(logLevel)

	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	return lfsHook
}

type TraceIdHook struct {
	TraceId string
}

func NewTraceIdHook(traceId string) logrus.Hook {
	hook := TraceIdHook{
		TraceId: traceId,
	}
	return &hook
}

func (hook *TraceIdHook) Fire(entry *logrus.Entry) error {
	entry.Data[global.TraceID] = hook.TraceId
	return nil
}

func (hook *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
