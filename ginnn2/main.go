package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context){
	c.File("assets/templates/index.html")
}

func main() {
	router := gin.Default()

	router.GET("/", IndexHandler)
	router.Static("/assets", "./assets")

	fmt.Println("gin server is running on port 8080")
	router.Run()
}