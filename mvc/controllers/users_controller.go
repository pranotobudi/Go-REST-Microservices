package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pranotobudi/Go-REST-Microservices/mvc/services"
	"github.com/pranotobudi/Go-REST-Microservices/mvc/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("user_id")
	log.Printf("About to process userID: %v", userIDParam)
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonVal, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonVal)
		return
	}

	userPtr, apiErr := services.GetUser(userID)
	if apiErr != nil {
		//handle the error
		jsonVal, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonVal)
		return
	}
	jsonValue, _ := json.Marshal(userPtr)

	w.Write(jsonValue)
}
