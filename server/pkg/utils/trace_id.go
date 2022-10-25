package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 19:12
 * @File: trace_id.go
 * @Description: //TODO $
 * @Version:
 */

func GetTraceId(c *gin.Context) string {
	traceId := uuid.NewV4().String()
	c.Set("traceId", traceId)
	return traceId
}
