package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"{{.Module}}/app/resolver"
)

func main() {
	app := gin.Default()

	// Load port from .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}

	router := app.Group("/api")

	srv := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{}}))
	router.POST("/graphql", gin.WrapH(srv))
	router.GET("/playground", gin.WrapH(playground.Handler("GraphQL Playground", "/graphql")))

	app.Run(":" + port)
}
