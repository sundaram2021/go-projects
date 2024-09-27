package router

import (
	"bookstore-api/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/books", controller.GetBooks)
		v1.POST("/books", controller.CreateBook)
		v1.PUT("/books/:id", controller.UpdateBook)
		v1.DELETE("/books/:id", controller.DeleteBook)
	}
}
