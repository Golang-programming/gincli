package main

import (
    "github.com/gin-gonic/gin"
    "os"
)

func main() {
    app := gin.Default()

    // Load port from .env
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified in .env
    }

    router := app.Group("/api")
    SetupRoutes(router)

    app.Run(":" + port)
}
