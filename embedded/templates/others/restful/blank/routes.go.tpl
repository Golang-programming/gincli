package {{.ResourceName}}

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	_groupRouter := router.Group("/{{.ResourceName}}")
	// Define your routes here
}
