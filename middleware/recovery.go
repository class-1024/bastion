package middleware

import (
	"bastion/pkg/response"
	"errors"
	"fmt"
	"github.com/georgehao/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"runtime/debug"
)

// Recovery 捕获所有panic，并且返回错误信息
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))

				log.Errorw("RecoveryMiddleware", "err", fmt.Sprint(err), "stack", string(debug.Stack()))

				if viper.GetString("base.debug_mode") != "debug" {
					response.Error(c, http.StatusInternalServerError, errors.New("内部错误"))
					return
				} else {
					response.Error(c, http.StatusInternalServerError, errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
		c.Next()
	}
}
