package main

import (
	"bookstore-api/router"

	"github.com/gin-gonic/gin"

	// Swagger-related imports
	"bookstore-api/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// @title BookStore API
// @version 1.0
// @description This is a sample Bookstore API using Gin framework.

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// Initialize routes
	router.InitializeRoutes(r)

	docs.SwaggerInfo.Title = "Swagger Bookstore API"
	docs.SwaggerInfo.Description = "This is a sample Bookstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the Gin server
	r.Run(":8080")
}
