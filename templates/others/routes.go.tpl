package {{.ResourceName}}

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/{{.ResourceName}}")

	groupRouter.POST("/")
	groupRouter.GET("/:id")
	groupRouter.PUT("/:id")
	groupRouter.DELETE("/:id")
	groupRouter.GET("/")
}
