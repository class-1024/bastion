package middleware

import (
	"bastion/pkg/constant"
	"bastion/pkg/errno"
	"bastion/pkg/response"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MiAppAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(constant.SessionKeyWeApp)

		fmt.Printf("session: %#v \n", user)
		if user == nil {
			response.Fail(c, errno.ErrorUserNotLogin, errors.New("未登录"))
			c.Abort()
			return
		}
		c.Next()
	}
}

func StatAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(constant.SessionKeyStatAdmin)

		fmt.Printf("session: %#v \n", user)
		if user == nil {
			response.Fail(c, errno.ErrorUserNotLogin, errors.New("未登录"))
			c.Abort()
			return
		}
		c.Next()
	}
}
