package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-REST-Microservices/oauth-api/src/api/domain/oauth"
	"github.com/pranotobudi/Go-REST-Microservices/oauth-api/src/api/services"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
	}
	token, err := services.OAuthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")
	token, err := services.OAuthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.Status(), err)

	}
	c.JSON(http.StatusOK, token)
}
