package {{.ResourceName}}

import (
	"github.com/gin-gonic/gin"
	"{{.Module}}/app/modules/{{.ResourceName}}/controllers"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/{{.ResourceName}}")

	groupRouter.POST("/", controllers.Create{{.CapitalizeResourceName}})
	groupRouter.GET("/:id", controllers.Get{{.CapitalizeResourceName}}ById)
	groupRouter.PUT("/:id", controllers.Update{{.CapitalizeResourceName}})
	groupRouter.DELETE("/:id", controllers.Delete{{.CapitalizeResourceName}})
	groupRouter.GET("/", controllers.GetAll{{.CapitalizeResourceName}}s)
}
