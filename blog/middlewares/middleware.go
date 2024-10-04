package middlewares

import (
	"blog/utils"
	"fmt"

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")

		if token == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "sign in please",
			})
			c.Abort()
		}
		c.Next()
	}
}

var jwtKey = []byte("my_secret_key") // Replace with your own secret key

func SetAdminToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define the claims for the token
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
			Issuer:    "admin",                               // Add relevant data (e.g., username or user ID)
		}

		// Create the token with the specified algorithm and claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign the token with your secret key
		tokenStr, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(500, gin.H{"error": "could not generate token"})
			return
		}

		// Set the JWT token as a cookie
		c.SetCookie("admin_token", tokenStr, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)
	}
}

func AdminMiddleware(c *gin.Context) {

	tokenString, err := c.Cookie("admin_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		// c.Abort()
		return
	}

	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		// c.Abort()
		return
	}

	fmt.Println("using adminn middleware")

	// Add claims to context, so they can be accessed in handlers
	c.Set("username", claims.Username)
}
