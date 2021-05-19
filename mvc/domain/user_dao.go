package domain

import (
	"fmt"
	"net/http"

	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Bud", LastName: "Sason", Email: "bud.sason@gmail.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {

	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
