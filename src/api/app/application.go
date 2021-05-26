package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/log"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}
func StartApp() {
	mapUrls()
	// http.HandleFunc("/users", controllers.GetUser)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }
	router.Run()
}
