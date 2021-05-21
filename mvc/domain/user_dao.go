package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Bud", LastName: "Sason", Email: "bud.sason@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("we're accesing the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
