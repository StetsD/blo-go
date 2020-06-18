package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stetsd/blo-go/ctrls"
	"github.com/stetsd/blo-go/middlewares"
)

func Init(router *gin.Engine) {

	router.Use(middlewares.SetUserStatus())

	router.GET("/", ctrls.ShowIndexPage)

	router.GET("/forbidden", ctrls.Forbidden)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/login", middlewares.EnsureNotLoggedIn(), ctrls.ShowLoginPage)
		userRoutes.POST("/login", middlewares.EnsureNotLoggedIn(), ctrls.PerformLogin)
		userRoutes.GET("/logout", middlewares.EnsureLoggedIn(), ctrls.Logout)
		userRoutes.GET("/register", middlewares.EnsureNotLoggedIn(), ctrls.ShowRegistrationPage)
		userRoutes.POST("/register", middlewares.EnsureNotLoggedIn(), ctrls.Register)

	}

	artRoutes := router.Group("/article")
	{
		artRoutes.GET("/", ctrls.ShowArticles)
		artRoutes.GET("/view/:article_id", ctrls.GetArticle)
		artRoutes.GET("/create", middlewares.EnsureLoggedIn(), ctrls.ShowArticleCreationPage)
		artRoutes.POST("/create", middlewares.EnsureLoggedIn(), ctrls.CreateArticle)
	}
}
