package controller

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/app/modules/{{.ResourceName}}/dtos"
    "{{.Module}}/app/modules/{{.ResourceName}}/services"
)

func Create{{.CapitalizeResourceName}}(c *gin.Context) {
    var input dtos.Create{{.CapitalizeResourceName}}Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    message := service.Create{{.CapitalizeResourceName}}(input)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func GetAll{{.CapitalizeResourceName}}s(c *gin.Context) {
    message := service.GetAll{{.CapitalizeResourceName}}s()
    c.JSON(200, gin.H{
        "message": message,
    })
}

func Get{{.CapitalizeResourceName}}ById(c *gin.Context) {
    ID := c.Param("id")

    message := service.Get{{.CapitalizeResourceName}}ById(ID)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func Update{{.CapitalizeResourceName}}(c *gin.Context) {
    ID := c.Param("id")
    var input dtos.Update{{.CapitalizeResourceName}}Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    message := service.Update{{.CapitalizeResourceName}}(ID, input)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func Delete{{.CapitalizeResourceName}}(c *gin.Context) {
    ID := c.Param("id")

    message := service.Delete{{.CapitalizeResourceName}}(ID)
    c.JSON(200, gin.H{
        "message": message,
    })
}