package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/client/restclient"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/repositories"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}
func TestCreateRepoInvalidInputName(t *testing.T) {
	restclient.FlushMockups()
	request := repositories.CreateRepoRequest{}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGitub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message":"Requires authentication",
				"documentation_url":"https://developer.github.com/docs"
			}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())

}
func TestCreateRepoNoError(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	// assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	output := make(chan repositories.CreateRepositoriesResult)
	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message":"Requires authentication",
				"documentation_url":"https://developer.github.com/docs"
			}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}

	output := make(chan repositories.CreateRepositoriesResult)
	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{
			"id":123, 
			"name":"testing",
			"owner":{"login":"federicoleon"}}`))},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}

	output := make(chan repositories.CreateRepositoriesResult)
	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	// assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "testing", result.Response.Name)
}

func TestHandleRepoResults(t *testing.T) {
	var wg sync.WaitGroup
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	service := reposService{}
	go service.handleRepoResults(&wg, input, output)
	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()
	wg.Wait()
	close(input)
	result := <-output
	assert.NotNil(t, result)

}

func TestCreateReposInvalidsRequest(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: ""},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)

	assert.Nil(t, result.Results[0].Response)
	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.Message())
}

func TestCreateReposOneSuccessOneFail(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{
			"id":123, 
			"name":"testing",
			"owner":{"login":"federicoleon"}}`))},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)

	//we can make sure which one is the first because it run on go routine.
	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			assert.EqualValues(t, "invalid repository name", result.Error.Message())
		} else {
			// assert.EqualValues(t, 123, result.Response.Id)
			assert.EqualValues(t, "testing", result.Response.Name)
			assert.EqualValues(t, "federicoleon", result.Response.Owner)
		}
	}
}

func TestCreateReposAlreadyExistsFailure(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{
			"id":123, 
			"name":"testing",
			"owner":{"login":"federicoleon"}}`))},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)

	//we can make sure which one is the first because it run on go routine.
	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusInternalServerError, result.Error.Status())
			assert.EqualValues(t, "error unmarshal github create repo response", result.Error.Message())
		} else {
			// assert.EqualValues(t, 123, result.Response.Id)
			assert.EqualValues(t, "testing", result.Response.Name)
			assert.EqualValues(t, "federicoleon", result.Response.Owner)
		}
	}
}
