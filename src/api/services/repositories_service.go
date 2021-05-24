package services

import (
	"net/http"
	"sync"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/config"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/github"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/repositories"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/provider/github_provider"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos([]repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (r *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (r *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)
	var wg sync.WaitGroup
	go r.handleRepoResults(&wg, input, output)
	for _, current_req := range requests {
		wg.Add(1)
		go r.createRepoConcurrent(current_req, input)
	}
	wg.Wait()
	close(input)
	result := <-output

	successCreation := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreation++
		}
	}

	if successCreation == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreation == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{Response: incomingEvent.Response, Error: incomingEvent.Error}
		results.Results = append(results.Results, repoResult)

		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
