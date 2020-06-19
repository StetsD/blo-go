package test

import (
	"github.com/stetsd/blo-go/store"
	"github.com/stetsd/blo-go/validators"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := store.GetAllArticles()

	if len(alist) != len(store.GetAllArticles()) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != store.GetAllArticles()[i].Content ||
			v.ID != store.GetAllArticles()[i].ID ||
			v.Title != store.GetAllArticles()[i].Title {

			t.Fail()
			break
		}
	}
}

func TestGetArticleByID(t *testing.T) {
	a, err := store.GetArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Turbo Power!!!" {
		t.Fail()
	}
}

func TestCreateNewArticle(t *testing.T) {
	originalLength := len(store.GetAllArticles())

	a, err := store.CreateNewArticle("New test title", "New test content")

	allArticles := store.GetAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}

func TestUserValidity(t *testing.T) {
	saveLists()

	if !validators.IsUserValid("boris", "britva") {
		t.Fail()
	}

	if validators.IsUserValid("drug", "britva") {
		t.Fail()
	}

	if validators.IsUserValid("boris", "") {
		t.Fail()
	}

	if validators.IsUserValid("", "britva") {
		t.Fail()
	}
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := validators.RegisterNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	u, err := validators.RegisterNewUser("boris", "britva")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = validators.RegisterNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUsernameAvailability(t *testing.T) {
	saveLists()

	if validators.IsUsernameAvailable("boris") {
		t.Fail()
	}

	validators.RegisterNewUser("newuser", "newpass")

	if validators.IsUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}