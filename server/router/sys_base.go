package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liuzw3018/otto/server/api/v1"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 18:21
 * @File: sys_user.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterSysUserRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("/")

	sysBase := v1.NewSysUser()
	{
		baseRouter.POST("login", sysBase.Login)
		baseRouter.POST("logout", sysBase.Logout)
	}
}
