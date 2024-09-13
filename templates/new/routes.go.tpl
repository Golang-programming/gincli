package routes

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/controllers"
)

func SetupRoutes(r *gin.RouterGroup) {
    r.POST("/", controllers.Create)
    r.GET("/", controllers.GetAll)
    r.GET("/:id", controllers.GetById)
    r.PUT("/", controllers.Update)
    r.DELETE("/", controllers.Delete)
}