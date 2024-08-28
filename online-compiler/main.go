package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default();

	// router.LoadHTMLFiles("index.html");
	
	router.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/?language=c"
    	router.HandleContext(ctx)
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"message": "file is loaded",
		})
	})
	// router.GET("/?language=c", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "index.html", gin.H{
	// 		"message": "file is loaded",
	// 	})
	// })

	fmt.Println("server is running on port 8080")
	router.Run()
}