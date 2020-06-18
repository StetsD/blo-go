package main

import (
	"github.com/stetsd/blo-go/ctrls"
	"github.com/stetsd/blo-go/middlewares"
)

func initializeRoutes() {

	router.Use(middlewares.SetUserStatus())

	router.GET("/", ctrls.ShowIndexPage)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/login", middlewares.EnsureNotLoggedIn(), ctrls.ShowLoginPage)
		userRoutes.GET("/login", middlewares.EnsureNotLoggedIn(), ctrls.PerformLogin)
		userRoutes.GET("/logout", middlewares.EnsureLoggedIn(), ctrls.Logout)
		userRoutes.GET("/login", middlewares.EnsureNotLoggedIn(), ctrls.ShowLoginPage)

	}
}