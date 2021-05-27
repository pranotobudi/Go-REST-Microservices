package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/repositories"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/services"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var (
	funcCreateRepo  func(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func([]repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct{}

func (r *repoServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(request)
}
func (r *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(request)
}

//THIS IS UNIT TESTING (only for this layer)
func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    123,
			Name:  "mocked services",
			Owner: "golang",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)
	// c, _ := gin.CreateTestContext(response)
	// c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "mocked services", result.Name)
	assert.EqualValues(t, "golang", result.Owner)
}

//THIS IS UNIT TESTING (only for this layer)
func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)
	// c, _ := gin.CreateTestContext(response)
	// c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	// var result repositories.CreateRepoResponse
	// err := json.Unmarshal(response.Body.Bytes(), &result)
	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}
