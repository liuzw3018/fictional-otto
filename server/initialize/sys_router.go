package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/otto/server/middleware"
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/liuzw3018/otto/server/router"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:50
 * @File: sys_router.go
 * @Description: //TODO $
 * @Version:
 */

func Routers() *gin.Engine {
	gin.SetMode(global.OttoConfig.Server.Mode)
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())
	r.Use(middleware.GinLogger())

	// 路由注册
	baseApiGroup := r.Group("/api")
	v1ApiGroup := baseApiGroup.Group("/v1")
	{
		router.RegisterSysAdminRouter(v1ApiGroup)
	}
	return r
}
