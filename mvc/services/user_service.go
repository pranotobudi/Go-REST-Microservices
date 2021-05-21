package services

import (
	"github.com/pranotobudi/Go-REST-Microservices/mvc/domain"
	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

type usersService struct{}

var UsersService usersService

func (u *usersService) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(id)
}
