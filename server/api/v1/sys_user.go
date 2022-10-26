package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/otto/server/model/dto"
	"github.com/liuzw3018/otto/server/pkg/public"
	"github.com/liuzw3018/otto/server/pkg/response"
	"net/http"
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

// Login
// @Tags 用户相关
// @Summary 用户登录
// @accept application/json
// @Produce application/json
// @Param data body dto.SysUserLoginInput true
// @Success 200 {object} response.Response{msg=response.Response}
// @Router /api/v1/login [post]
func (su *sysUser) Login(c *gin.Context) {
	var (
		code  public.ResponseCode
		token string
		out   = &dto.SysUserLoginOutput{}

		err error
	)
	params := &dto.SysUserLoginInput{}
	if err = params.BindingValidParams(c); err != nil {
		code = 2001
		goto ERR
	}

	if !(params.Username == "admin" && params.Password == "123456") {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err = public.GenToken(public.Users{Username: "admin", Password: "123456"})
	if err != nil {
		code = 2002
		goto ERR
	}

	out.Token = token
	out.Role = "admin"

	response.Success(c, public.SuccessCode, "ok", out)
	return

ERR:
	response.Error(c, code, err)
	return
}

func (su *sysUser) Logout(c *gin.Context) {

}