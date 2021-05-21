package services

import (
	"github.com/pranotobudi/Go-REST-Microservices/mvc/domain"
	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

type itemsService struct{}

var ItemsService itemsService

func (s *itemsService) GetItem(id int64) (*domain.Item, *utils.ApplicationError) {
	return domain.GetItem(id)
}
