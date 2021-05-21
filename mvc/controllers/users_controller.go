package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-REST-Microservices/mvc/services"
	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	// userIDParam := r.URL.Query().Get("user_id")
	userIDParam := c.Param("user_id")
	log.Printf("About to process userID: %v", userIDParam)
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		// jsonVal, _ := json.Marshal(apiErr)
		// w.WriteHeader(apiErr.StatusCode)
		// w.Write(jsonVal)
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	userPtr, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		//handle the error
		// jsonVal, _ := json.Marshal(apiErr)
		// w.WriteHeader(apiErr.StatusCode)
		// w.Write(jsonVal)
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	// jsonValue, _ := json.Marshal(userPtr)

	// w.Write(jsonValue)
	utils.Respond(c, http.StatusOK, userPtr)
}
