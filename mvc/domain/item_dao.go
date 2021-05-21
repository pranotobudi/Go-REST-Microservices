package domain

import "github.com/pranotobudi/Go-REST-Microservices/mvc/utils"

func GetItem(userID int64) (*Item, *utils.ApplicationError) {

	// if user := users[userID]; user != nil {
	// 	return user, nil
	// }
	// return nil, &utils.ApplicationError{
	// 	Message:    fmt.Sprintf("user %v not found", userID),
	// 	StatusCode: http.StatusNotFound,
	// 	Code:       "not found",
	// }
	return nil, nil
}
