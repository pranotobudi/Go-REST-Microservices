package oauth

import (
	"fmt"
	"strings"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() errors.ApiError {
	r.Username = strings.TrimSpace(r.Username)
	fmt.Println("username: ", r.Username)
	fmt.Println("password: ", r.Password)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}

	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
