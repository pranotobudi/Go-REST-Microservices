package github

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a go introduction directory",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}
	fmt.Println("Bismillah")
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))
	assert.EqualValues(t, `{"name":"golang introduction","description":"a go introduction directory","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`, string(bytes))
}
func TestCreateRepoRequestDynamicAttributeAsJson(t *testing.T) {
	request := CreateRepoRequestDynamicAttrib{
		Name:        "golang introduction",
		Description: "a go introduction directory",
		Homepage:    "https://github.com",
		Private:     true,
		CreatedDate: time.Now(),
		EditedDate:  time.Now(),
		HasWiki:     false,
	}
	fmt.Println("Bismillah")
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))
	assert.EqualValues(t, `{"name":"golang introduction","description":"a go introduction directory","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`, string(bytes))

}
