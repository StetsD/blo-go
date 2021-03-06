package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.HTML(http.StatusUnauthorized, "forbidden.html", gin.H{})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.HTML(http.StatusUnauthorized, "forbidden.html", gin.H{})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}