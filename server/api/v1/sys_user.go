package v1

import (
	"github.com/gin-gonic/gin"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 18:20
 * @File: sys_user.go
 * @Description: //TODO $
 * @Version:
 */

type sysUser struct{}

func NewSysUser() *sysUser {
	return &sysUser{}
}

func (su *sysUser) Login(c *gin.Context) {

}

func (su *sysUser) Logout(c *gin.Context) {

}
