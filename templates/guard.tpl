package guards

import (
    "github.com/gin-gonic/gin"
)

func {{.Name}}Guard(c *gin.Context) {
    // Implement guard logic here
    c.Next()
}
