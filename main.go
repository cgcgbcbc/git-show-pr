package main

import (
	"fmt"
	"github.com/github/hub/git"
	"github.com/github/hub/github"
	githubapi "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	var owner, name string
	var err error
	if len(os.Args) == 3 {
		owner = os.Args[1]
		name = os.Args[2]
	} else {
		owner, name, err = getOnwerRepo()
		handleError(err)
	}

	fmt.Printf("repo %s/%s ", owner, name)

	token, err := getToken()
	handleError(err)

	pulls, err := getPullRequests(token, owner, name)
	handleError(err)

	fmt.Printf("pull request count %d:\n", len(pulls))

	showPullRequest(pulls)
}

func getOnwerRepo() (owner string, name string, err error) {
	repo, err := github.LocalRepo()
	if err != nil {
		return
	}
	project, err := repo.CurrentProject()
	if err != nil {
		return
	}
	return project.Owner, project.Name, err
}

func getToken() (token string, err error) {
	return getConfig("github.token")
}

func getClient(token string) *githubapi.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return githubapi.NewClient(tc)
}

func getPullRequests(token string, owner string, repo string) (pulls []githubapi.PullRequest, err error) {
	client := getClient(token)
	pulls, _, err = client.PullRequests.List(owner, repo, nil)
	return
}

func showPullRequest(pulls []githubapi.PullRequest) {
	for _, pull := range pulls {
		fmt.Printf("#%d %d/%02d/%02d [\x1b[34m%s\x1b[0m] %s\n", *pull.Number, pull.CreatedAt.Year(), pull.CreatedAt.Month(), pull.CreatedAt.Day(), *pull.Head.Label, *pull.Title)
	}
}

func getConfig(key string) (value string, err error) {
	value, err = git.Config(key)
	if err != nil {
		return
	}
	if value != "" {
		return
	}
	value, err = git.GlobalConfig(key)
	return
}
