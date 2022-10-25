package global

import (
	"github.com/liuzw3018/otto/server/config"
	"github.com/liuzw3018/otto/server/pkg/zaplog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:55
 * @File: global.go
 * @Description: //TODO $
 * @Version:
 */

var (
	OttoConfig *config.Config
	OttoViper  *viper.Viper
	OttoLogrus *logrus.Entry
	OttoZap    *zaplog.Logger

	OttoLogger *FactoryLogger
)

const (
	TraceID = "traceId"
)
