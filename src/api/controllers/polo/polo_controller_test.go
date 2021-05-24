package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pranotobudi/Go-REST-Microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}
func TestPolo(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	Polo(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
