package oauth

import (
	"fmt"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var users = map[string]*User{
	"fede": {Id: 123, Username: "fede"},
}

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotfoundApiError(fmt.Sprint("no user found with given parameters"))
	}
	return user, nil
}
