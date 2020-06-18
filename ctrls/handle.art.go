package ctrls

import (
	"github.com/gin-gonic/gin"
	"github.com/stetsd/blo-go/renderer"
	"github.com/stetsd/blo-go/store"
	"net/http"
	"strconv"
)

func ShowIndexPage(c *gin.Context) {
	articles := store.GetAllArticles()

	renderer.Render(c, gin.H{
		"title": "Home Page",
		"payload": articles}, "index.html")
}


func ShowArticleCreationPage(c *gin.Context) {
	renderer.Render(c, gin.H{
		"title": "Create New Article",
	}, "create-article.html")
}

func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	article, err := store.GetArticleByID(articleID)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	renderer.Render(c, gin.H{
		"title": article.Title,
		"payload": article,
	}, "article.html")
}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	a, err := store.CreateNewArticle(title, content)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	renderer.Render(c, gin.H{
		"title": "Submission Successful",
		"payload": a,
	}, "submission-successful.html")
}