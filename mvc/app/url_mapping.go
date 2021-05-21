package app

import (
	"github.com/pranotobudi/Go-REST-Microservices/mvc/controllers"
)

func mapUrls() {
	// http.HandleFunc("/users/:user_id", controllers.GetUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
