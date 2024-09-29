package controllers

import (
	"net/http"

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

	message := services.Create{{.CapitalizeResourceName}}(input)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func GetAll{{.CapitalizeResourceName}}s(c *gin.Context) {
	message := services.GetAll{{.CapitalizeResourceName}}s()
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func Get{{.CapitalizeResourceName}}ById(c *gin.Context) {
	ID := c.Param("id")

	message := services.Get{{.CapitalizeResourceName}}ById(ID)
	c.JSON(http.StatusOK, gin.H{
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

	message := services.Update{{.CapitalizeResourceName}}(ID, input)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func Delete{{.CapitalizeResourceName}}(c *gin.Context) {
	ID := c.Param("id")

	message := services.Delete{{.CapitalizeResourceName}}(ID)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
