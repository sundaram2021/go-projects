package handlers

import (
	// "blog/routes"
	// "blog/utils"
	// "fmt"
	"blog/middlewares"
	"net/http"
	_ "time"

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



func SignInHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json binding error while signin",
		})
		return
	}

	if user.Name == "sundaram" {
		middlewares.SetAdminToken()
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
