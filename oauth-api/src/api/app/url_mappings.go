package app

import (
	"github.com/pranotobudi/Go-REST-Microservices/oauth-api/src/api/controller/oauth"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
