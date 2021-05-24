package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
	assert.EqualValues(t, "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", githubAccessToken) // using go test terminal
	// assert.EqualValues(t, "", githubAccessToken) //using run test output
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "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", GetGithubAccessToken()) // using go test terminal
	// assert.EqualValues(t, "", GetGithubAccessToken()) // using run test output

}
