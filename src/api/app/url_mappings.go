package app

import (
	"github.com/pranotobudi/Go-REST-Microservices/src/api/controllers/polo"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/controllers/repositories"
)

func mapUrls() {
	// http.HandleFunc("/users/:user_id", controllers.GetUser)
	router.POST("/repositories", repositories.CreateRepo)
	router.GET("/marco", polo.Polo)
}
