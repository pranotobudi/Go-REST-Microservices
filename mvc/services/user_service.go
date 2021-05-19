package services

import (
	"github.com/pranotobudi/Go-REST-Microservices/mvc/domain"
	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

func GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(id)
}
