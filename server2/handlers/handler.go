package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "welcome to the root handler",
	})
}

func RefreshHandler(c *gin.Context){
	session := sessions.Default(c)
	sessionToken := session.Get("token")
	sessionUser := session.Get("username")

	if sessionToken == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message" : "seesion token is empty",
		})
		return
	}

	sessionToken = xid.New().String()
	session.Set("username", sessionUser.(string))
	session.Set("token", sessionToken)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "New session issued"})
}