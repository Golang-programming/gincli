package middleware

import (
	"github.com/gin-gonic/gin"
)

func {{.CapitalizeMiddlewareName}}Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your middleware logic here
		c.Next()
	}
}
