package store

import (
	"errors"
	"github.com/stetsd/blo-go/models"
)

var articleList = []models.Article{
	models.Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	models.Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []models.Article {
	return articleList
}

func GetArticleByID(id int) (*models.Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}

func CreateNewArticle(title, content string) (*models.Article, error) {
	a := models.Article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}