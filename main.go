package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stetsd/blo-go/routes"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	routes.Init(router)

	router.Run()
}