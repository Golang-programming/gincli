    touch main.go.tpl
    touch .env.tpl
    touch loadEnv.go.tpl
    mkdir app && cd app
    mkdir pkg && mkdir pkg/database
    touch pkg/database/database.go.tpl
    mkdir utils
    touch utils/sum-to-numbers.go.tpl
    mkdir controllers
    mkdir services
    touch controllers/controller.go.tpl
    touch services/service.go.tpl
    touch routes.go.tpl
    mkdir middleware
    mkdir modules


here are files content
// .env
DB_HOST=localhost
DB_USER={{.DBUsername}}
DB_PASSWORD={{.DBPassword}}
DB_NAME={{.DBName}}
DB_PORT={{.DBPort}}


// main.go
package main

import (
    "{{.Module}}/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()

    router := app.Group("/api")
    routes.SetupRoutes(router)

    app.Run("specify port from .env")
}


// database.go
package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/{{.DBDriver}}"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open({{.DBDriver}}.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
}


// loadEnv.go
package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables from a configuration file
	fmt.Println("Loading environment variables from a configuration file")
	godotenv.Load()
}


// routes.go
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


// controller.go
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


// service.go
package service

func Create() string {
    return fmt.Sprintf("add some meaningful message")
}
func GetAll() string {
    return fmt.Sprintf("add some meaningful message")
}
func GetById(id) string {
    return fmt.Sprintf("add some meaningful message: %s", id)
}
func Update() string {
    return fmt.Sprintf("add some meaningful message")
}
func Delete() string {
    return fmt.Sprintf("add some meaningful message")
}

