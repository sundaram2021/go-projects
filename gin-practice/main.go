package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	route := gin.Default()
	route.LoadHTMLFiles("form.html") // Load the HTML file
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})
	route.POST("/submit", getName)
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func getName(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println("-----binding html-------")
		log.Println("name : " + person.Name)
	}
	c.String(http.StatusOK, "success")
}

func startPage(c *gin.Context) {
	var person Person
	if c.BindQuery(&person) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(http.StatusOK, "Success")
}
