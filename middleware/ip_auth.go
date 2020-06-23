package middleware

import (
	"bastion/pkg/errno"
	"bastion/pkg/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IPAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range viper.GetStringSlice("http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			response.Fail(c, errno.ErrorIpNotAllow, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
