package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Define a struct to bind form data
type MyForm struct {
	Colors []string `form:"colors[]"` // Specify the form tag for checkbox binding
}

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Serve the HTML form
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	// Handle form submission
	router.POST("/submit", func(c *gin.Context) {
		var form MyForm
		// Bind the form data to the struct
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond with the received data
		c.JSON(http.StatusOK, gin.H{"selected_colors": form.Colors})
	})

	// Load HTML templates
	router.LoadHTMLFiles("form.html")

	// Run the server
	router.Run(":8080")
}
