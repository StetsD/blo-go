package main

import "github.com/stetsd/blo-go/ctrls"

func initializeRoutes() {
	router.GET("/", ctrls.ShowIndexPage)
}