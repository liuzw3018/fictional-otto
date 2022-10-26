package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/otto/server/api"
	"github.com/liuzw3018/otto/server/pkg/response"
	"net/http"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:40
 * @File: sys_admin.go
 * @Description: //TODO $
 * @Version:
 */

type sysAdmin struct{}

func NewSysAdmin() api.BaseApiHandle {
	return &sysAdmin{}
}

func (sa *sysAdmin) GetAll(c *gin.Context) {
	value, _ := c.Get("claims")
	response.Success(c, http.StatusOK, "OK", value)
}

func (sa *sysAdmin) GetOne(c *gin.Context) {

}

func (sa *sysAdmin) UpdateOne(c *gin.Context) {

}

func (sa *sysAdmin) UpdateAll(c *gin.Context) {

}

func (sa *sysAdmin) Add(c *gin.Context) {

}

func (sa *sysAdmin) DeleteOne(c *gin.Context) {

}

func (sa *sysAdmin) DeleteMore(c *gin.Context) {

}
