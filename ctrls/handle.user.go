package ctrls

import (
	"github.com/gin-gonic/gin"
	"github.com/stetsd/blo-go/renderer"
	"github.com/stetsd/blo-go/store"
	"github.com/stetsd/blo-go/validators"
	"math/rand"
	"net/http"
	"strconv"
)

func ShowLoginPage(c *gin.Context) {
	renderer.Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func PerformLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if validators.IsUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", true, false)
		c.Set("is_logged_in", true)

		renderer.Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", true, false)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ShowRegistrationPage(c *gin.Context) {
	renderer.Render(c, gin.H{
		"title": "Register"}, "register.html")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := store.CreateNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", true, false)
		c.Set("is_logged_in", true)

		renderer.Render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}