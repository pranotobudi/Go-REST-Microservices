package oauth

import "github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"

var tokens = make(map[string]*AccessToken, 0)

func (at *AccessToken) Save() errors.ApiError {
	at.Token = "2343242"
	tokens[at.Token] = at
	return nil
}
func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil || token.IsExpired() {
		return nil, errors.NewNotfoundApiError("no access token found with given parameters")
	}
	return token, nil
}
