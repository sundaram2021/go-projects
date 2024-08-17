package routes

import (
	"blog/handlers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// SetupRouter initializes and returns a Gin engine instance
func SetupRouter() *gin.Engine {
	router = gin.Default()

	store := cookie.NewStore([]byte("blogsecret"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost", // Use just "localhost" for local development
		MaxAge:   3600 * 2,    // 2 hours
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	router.Use(sessions.Sessions("blog_api", store))

	router.GET("/", handlers.RootHandler)
	router.POST("/signin", handlers.SignInHandler)
	router.GET("/signout", handlers.SignOutHandler)

	authorize := router.Group("/")
	authorize.Use(handlers.AuthMiddleware())
	{
		authorize.GET("/blogs", handlers.GetAllBlogs)
		authorize.GET("/blogs/:id", handlers.GetBlogById)
		authorize.POST("blogs", handlers.CreateBlog)
		authorize.PUT("/blogs/:id", handlers.UpadateBlog)
		authorize.DELETE("/blogs/:id",handlers.AdminMiddleware(), handlers.DeleteBlog)
	}

	return router
}
