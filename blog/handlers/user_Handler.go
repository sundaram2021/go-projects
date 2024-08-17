package handlers

import (
	"blog/utils"
	"net/http"
	"time"

	// "blog/routes"
	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type User struct {
	Id   string `json:id`
	Name string `json:name`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("blog_api")

		if token == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "sign in please",
			})
			c.Abort()
		}
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("admin_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Add claims to context, so they can be accessed in handlers
		c.Set("username", claims.Username)

		c.Next()
	}
}

func SignInHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json binding error while signin",
		})
		return
	}

	// if user.Name == "user22" {
	// 	store := cookie.NewStore([]byte("blogsecret"))
	// 	store.Options(sessions.Options{
	// 		Path:     "/",
	// 		Domain:   "http://localhost:8080", // Adjust to your domain
	// 		MaxAge:   3600 * 2,          // 8 hours
	// 		HttpOnly: true,
	// 		SameSite: http.SameSiteLaxMode,
	// 	})
	// 	routes.Use(sessions.Sessions("blog_api", store))
	// }

	// Generate JWT
	if user.Name == "sundaram" {
		tokenStr, err := utils.GenerateJWT(user.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}

		// Set token in response cookie
		c.SetCookie("admin_token", tokenStr, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)

	}

	tokenString := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Name)
	session.Set("token", tokenString)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "user signed in sucessfully",
	})

}

func SignOutHandler(c *gin.Context) {
	session := sessions.Default(c)

	// Check if the user is logged in by checking the session token
	tokenValue := session.Get("token")

	if tokenValue == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "user is not logged in",
		})
		return
	}

	// Delete specific session values or clear the session entirely
	session.Delete("username")
	session.Delete("token")
	session.Clear() // This will clear all session values
	session.Save()  // Don't forget to save the session to apply changes

	c.JSON(http.StatusOK, gin.H{
		"message": "user signed out successfully",
	})
}
