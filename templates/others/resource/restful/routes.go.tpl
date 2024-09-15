package food

import (
	"github.com/gin-gonic/gin"
	"{{.Module}}/app/modules/{{.ResourceName}}/controllers"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/{{.ResourceName}}")

	groupRouter.POST("/", controllers.CreateCapitalizeResourceName)
	groupRouter.GET("/:id", controllers.GetCapitalizeResourceNameByID)
	groupRouter.PUT("/:id", controllers.UpdateCapitalizeResourceName)
	groupRouter.DELETE("/:id", controllers.DeleteCapitalizeResourceName)
	groupRouter.GET("/", controllers.ListCapitalizeResourceNames)
}
