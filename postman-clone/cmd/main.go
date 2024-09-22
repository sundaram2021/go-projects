package main

import (
	"net/http"
	"postman-clone/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLFiles("views/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/request", controllers.HandleRequest)

	r.Run(":8080") // listen and serve on localhost:8080
}
