package food

import (
	"github.com/gin-gonic/gin"
	"{{.Module}}/app/modules/{{.ResourceName}}/controllers"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/{{.ResourceName}}")

	groupRouter.POST("/", controllers.Create{{Capitalize .ResourceName}})
	groupRouter.GET("/:id", controllers.Get{{Capitalize .ResourceName}}ByID)
	groupRouter.PUT("/:id", controllers.Update{{Capitalize .ResourceName}})
	groupRouter.DELETE("/:id", controllers.Delete{{Capitalize .ResourceName}})
	groupRouter.GET("/", controllers.List{{Capitalize .ResourceName}}s)
}
