package main

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/app/controller"
)

func SetupRoutes(r *gin.RouterGroup) {
    r.POST("/", controller.Create)
    r.GET("/", controller.GetAll)
    r.GET("/:id", controller.GetById)
    r.PUT("/", controller.Update)
    r.DELETE("/", controller.Delete)
}
