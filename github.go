package kickoff

import (
	"context"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

func NewGitHubOAuthHTTPClient() *http.Client {
	ctx := context.Background()
	githubAccessToken, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Fatal("GITHUB_TOKEN not set in environment")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	return oauth2.NewClient(ctx, ts)
}

func NewGitHubKickoffClient(owner string, repo string, client *http.Client) *GitHubKickoffClient {
	return &GitHubKickoffClient{
		github.NewClient(client),
		owner,
		repo,
	}
}

type GitHubKickoffClient struct {
	gitHubClient *github.Client
	owner        string
	repo         string
}

func (c GitHubKickoffClient) Add(ko *Kickoff) error {
	_, _, err := c.gitHubClient.Issues.Create(context.Background(), c.owner, c.repo, &github.IssueRequest{
		Title: github.String(ko.Title()),
	})
	return err
}
