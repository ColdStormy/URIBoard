package usermanager

import (
	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID
	Name           string
	PasswordHashed string
}

var users = [2]User{
	User{uuid.New(), "alfred", "pass"},
	User{uuid.New(), "admin", "secure"},
}

func ValidateUser(username string, password string) bool {

	for _, user := range users {
		if user.Name == username && user.PasswordHashed == password {
			return true
		}
	}

	return false
}

func NameToId(username string) uuid.UUID {
	for _, user := range users {
		if user.Name == username {
			return user.Id
		}
	}

	return uuid.UUID{}
}

func GetUserById(uid uuid.UUID) (User, bool) {
	for _, user := range users {
		if user.Id == uid {
			return user, true
		}
	}

	return User{}, false
}
