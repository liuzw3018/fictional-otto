package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/liuzw3018/otto/server/pkg/utils"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 19:11
 * @File: logger.go
 * @Description: //TODO $
 * @Version:
 */

// AddTraceId 在每个处理器前添加该处理函数，为每个请求添加traceId
func AddTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := utils.GetTraceId(c)
		global.OttoLogger.Logger.AddHook(utils.NewTraceIdHook(traceId))
	}
}

func LoggerForGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		//不记录OPTION方法访问
		switch reqMethod {
		case "OPTIONS":
			return
		}

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		// 若被nginx代理，则应在X-Real-IP中获取真实访问IP
		clientIP := c.GetHeader("X-Real-IP")
		if clientIP == "" {
			clientIP = c.ClientIP()
		}
		//traceId := c.GetHeader("traceId")
		// clientIP := c.ClientIP()
		// realIP := c.GetHeader("X-Real-IP")

		//自带互斥锁不需要加锁
		global.OttoLogger.Infof("| %3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)

	}
}
