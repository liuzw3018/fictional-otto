package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liuzw3018/otto/server/api/v1"
	"github.com/liuzw3018/otto/server/middleware"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:49
 * @File: sys_admin.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterSysAdminRouter(router *gin.RouterGroup) {
	adminRouter := router.Group("/admin")
	adminRouter.Use(middleware.AddTraceId())
	sysAdmin := v1.NewSysAdmin()
	{
		adminRouter.GET("/:id", sysAdmin.GetOne)
		adminRouter.GET("", sysAdmin.GetAll)
		adminRouter.POST("", sysAdmin.Add)
		adminRouter.PUT("/:id", sysAdmin.UpdateOne)
		adminRouter.PUT("", sysAdmin.UpdateAll)
		adminRouter.DELETE("/:id", sysAdmin.DeleteOne)
		adminRouter.DELETE("", sysAdmin.DeleteMore)

	}
}
