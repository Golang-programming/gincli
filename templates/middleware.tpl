package middlewares

import (
    "github.com/gin-gonic/gin"
)

func {{.Name}}Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Implement middleware logic here
        c.Next()
    }
}
