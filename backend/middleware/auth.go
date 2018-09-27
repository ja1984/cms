package middleware

import (
	"github.com/gin-gonic/gin"
)

func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
