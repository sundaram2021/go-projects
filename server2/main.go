package main 

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server2/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("something_new_secret"))
	router.Use(sessions.Sessions("recipes_api", store))

	router.GET("/recipes", handlers.GetAllRecipes)
	router.POST("/signin", handlers.SignInHandler)
	router.POST("/refresh", handlers.RefreshHandler)

	authorized := router.Group("/")

	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.GET("/", handlers.RootHandler)
		authorized.GET("/recipes/:Id", handlers.GetRecipeById)
		authorized.DELETE("/recipes/:Id", handlers.DeleteRecipe)
		authorized.PUT("/recipes/:Id", handlers.UpdateRecipe)
		authorized.POST("/recipes", handlers.CreateRecipe)
	}


	fmt.Println("server is running on port 8080")
	router.Run()
}