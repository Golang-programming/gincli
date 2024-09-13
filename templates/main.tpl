package main

import (
    "{{.Module}}/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()

    router := app.Group("/api")
    routes.SetupRoutes(router)

    app.Run("")
}
