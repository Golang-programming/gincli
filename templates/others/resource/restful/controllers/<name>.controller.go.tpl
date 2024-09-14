package controller

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/app/modules/{{.ResourceName}}/dtos"
    "{{.Module}}/app/modules/{{.ResourceName}}/services"
)

func Create{{Capitalize .ResourceName}}(c *gin.Context) {
    var input dtos.Create{{Capitalize .ResourceName}}Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    message := service.Create{{Capitalize .ResourceName}}(input)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func GetAll{{Capitalize .ResourceName}}s(c *gin.Context) {
    message := service.GetAll{{Capitalize .ResourceName}}s()
    c.JSON(200, gin.H{
        "message": message,
    })
}

func Get{{Capitalize .ResourceName}}ById(c *gin.Context) {
    ID := c.Param("id")

    message := service.Get{{Capitalize .ResourceName}}ById(ID)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func Update{{Capitalize .ResourceName}}(c *gin.Context) {
    ID := c.Param("id")
    var input dtos.Update{{Capitalize .ResourceName}}Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    message := service.Update{{Capitalize .ResourceName}}(ID, input)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func Delete{{Capitalize .ResourceName}}(c *gin.Context) {
    ID := c.Param("id")

    message := service.Delete{{Capitalize .ResourceName}}(ID)
    c.JSON(200, gin.H{
        "message": message,
    })
}