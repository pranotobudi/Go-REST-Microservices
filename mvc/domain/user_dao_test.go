package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	userPtr, err := UserDao.GetUser(0)
	assert.Nil(t, userPtr, "we're not expecting a user with id==0")
	assert.NotNil(t, err)
	assert.EqualValues(t, "not found", err.Code)
	assert.EqualValues(t, "user 0 not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	// if userPtr != nil {
	// 	t.Error("we're not expecting a user with id==0")
	// }
	// if err == nil {
	// 	t.Error("we're expecting error with id==0")
	// }
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when user is not found")
	// }
}

func TestGetUserNoError(t *testing.T) {
	userPtr, err := UserDao.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, userPtr)
	assert.EqualValues(t, 123, userPtr.Id)
	assert.EqualValues(t, "Bud", userPtr.FirstName)
	assert.EqualValues(t, "Sason", userPtr.LastName)
	assert.EqualValues(t, "bud.sason@gmail.com", userPtr.Email)

}
