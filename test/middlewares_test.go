package test

import (
	"github.com/stetsd/blo-go/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", setLoggedIn(false), middlewares.EnsureLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func TestEnsureLoggedInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), middlewares.EnsureLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoggedInAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", setLoggedIn(true), middlewares.EnsureNotLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func TestEnsureNotLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), middlewares.EnsureNotLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestSetUserStatusAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", middlewares.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}

	r.ServeHTTP(w, req)
}

func TestSetUserStatusUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", middlewares.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if exists && loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)
}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}