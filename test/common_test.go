package test

import (
	"github.com/stetsd/blo-go/middlewares"
	"github.com/stetsd/blo-go/models"
	"github.com/stetsd/blo-go/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpUserList []models.User
var tmpArticleList []models.Article

var userList []models.User
var articleList []models.Article

func init() {
	userList = append(userList, store.GetUserList()...)
	articleList = append(articleList, store.GetAllArticles()...)

	tmpUserList = store.GetUserList()
	tmpArticleList = store.GetAllArticles()
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../templates/*")
		r.Use(middlewares.SetUserStatus())
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}

func saveLists() {
	copy(tmpUserList, store.GetUserList())
	copy(tmpArticleList, store.GetAllArticles())
}

func restoreLists() {
	tmpUserList = append([]models.User{}, userList...)
	tmpArticleList = append([]models.Article{}, articleList...)
}