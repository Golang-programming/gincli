package controller

import (
    "github.com/gin-gonic/gin"
    "{{.Module}}/app/modules/{{.ResourceName}}/dtos"
    "{{.Module}}/app/modules/{{.ResourceName}}/services"
)

func CreateCapitalizeResourceName(c *gin.Context) {
    var input dtos.CreateCapitalizeResourceNameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    message := service.CreateCapitalizeResourceName(input)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func GetAllCapitalizeResourceNames(c *gin.Context) {
    message := service.GetAllCapitalizeResourceNames()
    c.JSON(200, gin.H{
        "message": message,
    })
}

func GetCapitalizeResourceNameById(c *gin.Context) {
    ID := c.Param("id")

    message := service.GetCapitalizeResourceNameById(ID)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func UpdateCapitalizeResourceName(c *gin.Context) {
    ID := c.Param("id")
    var input dtos.UpdateCapitalizeResourceNameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    message := service.UpdateCapitalizeResourceName(ID, input)
    c.JSON(200, gin.H{
        "message": message,
    })
}

func DeleteCapitalizeResourceName(c *gin.Context) {
    ID := c.Param("id")

    message := service.DeleteCapitalizeResourceName(ID)
    c.JSON(200, gin.H{
        "message": message,
    })
}