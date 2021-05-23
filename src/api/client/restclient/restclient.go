package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var enabledMocks = false
var mocks = make(map[string]*Mock)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func StartMockups() {
	enabledMocks = true
}
func StopMockups() {
	enabledMocks = false
}
func FlushMockups() {
	mocks = make(map[string]*Mock)
}
func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}
func AddMockups(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		//return local mocks without calling external resource!
		mock := mocks[getMockId(http.MethodPost, url)]
		// fmt.Println("mocks:", mocks)
		// fmt.Println("mock:", mock)
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers
	client := http.Client{}
	return client.Do(request)
}
