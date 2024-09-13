package routes

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/controller"
)

func SetupRoutes(r *gin.RouterGroup) {
    r.POST("/", controller.Create)
    r.GET("/", controller.Read)
    r.PUT("/", controller.Update)
    r.DELETE("/", controller.Delete)
}
