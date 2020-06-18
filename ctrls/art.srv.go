package ctrls

import (
	"github.com/gin-gonic/gin"
	"github.com/stetsd/blo-go/renderer"
)

func Forbidden(c *gin.Context) {
	renderer.Render(c, gin.H{}, "forbidden.html")
}
