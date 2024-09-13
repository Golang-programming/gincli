package controller

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/service"
)

func Create(c *gin.Context) {

    message = service.Create()

    c.JSON(200, gin.H{
        "message",
    })
}

func GetAll(c *gin.Context) {

    message = service.GetAll()

    c.JSON(200, gin.H{
        "message",
    })
}

func GetById(c *gin.Context) {

    // get ID from params

    message = service.GetById(id)

    c.JSON(200, gin.H{
        "message",
    })
}

func Update(c *gin.Context) {

    message = service.Update()

    c.JSON(200, gin.H{
        "message",
    })
}

func Delete(c *gin.Context) {

    message = service.Delete()

    c.JSON(200, gin.H{
        "message",
    })
}
