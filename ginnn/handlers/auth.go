package handlers

import (
	"fmt"
	// "log"
	"net/http"
	"time"

	"ginnn/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "github.com/gorilla/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/rs/xid"
)

type AuthHandler struct{}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (handler *AuthHandler) SignInHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Hello aaaa5")
	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	fmt.Printf("session token :%v", session)
	session.Save()

	fmt.Println("Hello aaaa43 : ")

	c.JSON(http.StatusOK, gin.H{"message": "User signed in"})
}

///----------hashing password in go in just two lines using gin
// h := sha256.New()
// "password": string(h.Sum([]byte(password))),
