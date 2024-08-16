package handlers

import (
	// "crypto/sha256"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		if sessionToken == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not logged",
			})
			c.Abort()
		}
		c.Next()
	}
}

func SignInHandler(c *gin.Context) {
	var user User 
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// h := sha256.New()
	// cur := handler.collection.FindOne(handler.ctx, bson.M{
	// 	"username": user.Username,
	// 	"password": string(h.Sum([]byte(user.Password))),
	// })
	// if cur.Err() != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	// 	return
	// }

	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "User signed in"})
}