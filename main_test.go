package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPullRequests(t *testing.T) {
	token, _ := getToken()
	owner, name, _ := getOnwerRepo()
	pulls, err := getPullRequests(token, owner, name)
	assert.Nil(t, err)
	assert.NotNil(t, pulls)
	assert.Equal(t, len(pulls), 0, "should get pull requests")
	fmt.Println(pulls)
}

func TestGetOwnerRepo(t *testing.T) {
	owner, name, err := getOnwerRepo()
	assert.Nil(t, err)
	assert.Equal(t, owner, "cgcgbcbc", "owner should be correct")
	assert.Equal(t, name, "git-show-pr", "repo should be correct")
}

func TestGetToken(t *testing.T) {
	token, err := getToken()
	assert.Nil(t, err)
	assert.NotEqual(t, token, "", "token should not be empty")
}
