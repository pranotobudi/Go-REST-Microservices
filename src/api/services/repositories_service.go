package services

import (
	"strings"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/config"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/github"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/repositories"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/provider/github_provider"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (r *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
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
