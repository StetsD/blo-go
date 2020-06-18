package validators

import (
	"errors"
	"github.com/stetsd/blo-go/models"
	"github.com/stetsd/blo-go/store"
	"strings"
)

func IsUserValid(username, password string) bool {
	for _, u := range store.GetUserList() {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func RegisterNewUser(username, password string) (*models.User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	u, _ := store.CreateNewUser(username, password)

	return &u, nil
}

func IsUsernameAvailable(username string) bool {
	for _, u := range store.GetUserList() {
		if u.Username == username {
			return false
		}
	}
	return true
}