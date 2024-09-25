package guards

import (
	"github.com/gin-gonic/gin"
)

// ExampleGuardMiddleware example middleware
func ExampleGuardMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your guard logic here
		c.Next()
	}
}
