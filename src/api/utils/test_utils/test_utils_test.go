package test_utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockContext(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:123/repositories", strings.NewReader(`{"name":"testing"}`))
	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockContext(request, response)

	assert.NotNil(t, c)
	assert.EqualValues(t, http.MethodPost, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/repositories", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
	assert.EqualValues(t, "true", c.GetHeader("x-mock"))
}
