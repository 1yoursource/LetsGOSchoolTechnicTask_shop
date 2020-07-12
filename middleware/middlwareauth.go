package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webServer/services"
)

func TokenLogin(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil {
		_, ok := services.Tokens[token]
		if ok {
			//c.JSON(http.StatusOK, user)
			return
		}
	}
	c.Redirect(http.StatusPermanentRedirect, "/signup")
	c.Abort()
}

func EnsureLoggedIn(c *gin.Context) {
	// If there's an error or if the token is empty
	// the user is not logged in
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if !loggedIn {
		//if token, err := c.Cookie("token"); err != nil || token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func EnsureNotLoggedIn(c *gin.Context) {
	// If there's no error or if the token is not empty
	// the user is already logged in
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if loggedIn {
		// if token, err := c.Cookie("token"); err == nil || token != "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// This middleware sets whether the user is logged in or not
func SetUserStatus(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil || token != "" {
		c.Set("is_logged_in", true)
	} else {
		c.Set("is_logged_in", false)
	}
}
