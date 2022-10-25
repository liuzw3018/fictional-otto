package initialize

import (
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/liuzw3018/otto/server/pkg/utils"
	"github.com/sirupsen/logrus"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 18:35
 * @File: sys_logger.go
 * @Description: //TODO $
 * @Version:
 */

func initLogrus() {
	l := logrus.New()
	l.AddHook(utils.NewLfsHook(
		global.OttoConfig.Log.Level,
		global.OttoConfig.Log.Split,
		global.OttoConfig.Log.MaxRemainCnt,
	))

	l.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	global.OttoLogrus = logrus.NewEntry(l).WithField("Server", "otto")
}
