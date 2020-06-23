package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Noop() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("---------------middleware--Noop--before----------------")
		c.Next()
		fmt.Println("---------------middleware--Noop--after-----------------")
	}
}
