package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/domain/repositories"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/services"
	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/errors"
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)
	file, err := os.Open("requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}
	return result

}
func main() {
	requests := getRequests()
	fmt.Printf("processing %d line \n", len(requests))
	outputResult := make(chan createRepoResult)
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResult(&wg, outputResult)
	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, request, outputResult)
	}
	wg.Wait()
	close(outputResult)

	//Now you can process the success and failed result to disk, or email, or anything
}

func handleResult(wg *sync.WaitGroup, input chan createRepoResult) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Request.Name
		}
		wg.Done()
	}
}
func createRepo(buffer chan bool, request repositories.CreateRepoRequest, output chan createRepoResult) {
	result, err := services.RepositoryService.CreateRepo(request)
	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buffer
}
