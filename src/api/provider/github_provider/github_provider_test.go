package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/client/restclient"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

// func TestMain(t *testing.T) {
// 	restclient.StartMockups()
// }
func TestAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	//mock initialization and execution
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	//mock validation
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)
}
func TestCreateRepoInvalidResponseBody(t *testing.T) {
	//mock initialization and execution
	restclient.StartMockups()
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	//mock validation
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	//mock initialization and execution
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
		},
	})
	//mock validation
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	//mock initialization and execution
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message": "Requires authentication",
				"documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"
			}`)),
		},
	})
	//mock validation
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

// func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
// 	//mock initialization and execution
// 	restclient.StartMockups()
// 	restclient.FlushMockups()
// 	restclient.AddMockups(restclient.Mock{
// 		Url:        "https://api.github.com/user/repos",
// 		HttpMethod: http.MethodPost,
// 		Response: &http.Response{
// 			StatusCode: http.StatusCreated,
// 			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
// 		},
// 	})
// 	//mock validation
// 	response, err := CreateRepo("", github.CreateRepoRequest{})

// 	assert.Nil(t, response)
// 	assert.NotNil(t, err)
// 	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
// 	assert.EqualValues(t, "error unmarshal github create repo response", err.Message)
// }

// func TestCreateRepoNoError(t *testing.T) {
// 	//mock initialization and execution
// 	restclient.StartMockups()
// 	restclient.FlushMockups()
// 	jsonResponse, _ := json.Marshal(github.CreateRepoResponse{
// 		Id:       123,
// 		Name:     "repo-name",
// 		FullName: "pranotobudi/repo-name",
// 	})
// 	restclient.AddMockups(restclient.Mock{
// 		Url:        "https://api.github.com/user/repos",
// 		HttpMethod: http.MethodPost,
// 		Response: &http.Response{
// 			StatusCode: http.StatusCreated,
// 			Body:       ioutil.NopCloser(bytes.NewReader(jsonResponse)),
// 		},
// 	})
// 	//mock validation
// 	response, err := CreateRepo("", github.CreateRepoRequest{})

// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	assert.EqualValues(t, http.StatusCreated, err.StatusCode)
// 	assert.EqualValues(t, 123, response.Id)
// 	assert.EqualValues(t, "budi", response.Name)
// 	assert.EqualValues(t, "pranoto budi", response.FullName)
// }
