package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/otto/server/api"
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

func (s *sysAdmin) GetAll(c *gin.Context) {
	traceId, _ := c.Get("traceId")
	c.String(http.StatusOK, traceId.(string))
}

func (s *sysAdmin) GetOne(c *gin.Context) {

}

func (s *sysAdmin) UpdateOne(c *gin.Context) {

}

func (s *sysAdmin) UpdateAll(c *gin.Context) {

}

func (s *sysAdmin) Add(c *gin.Context) {

}

func (s *sysAdmin) DeleteOne(c *gin.Context) {

}

func (s *sysAdmin) DeleteMore(c *gin.Context) {

}
