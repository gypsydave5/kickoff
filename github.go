package kickoff

import (
	"context"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

// NewGitHubOAuthHTTPClient makes an HttpClient authorized with GitHub using a OAuth2 token. The token must be set in the execution environment as GITHUB_TOKEN.
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

// NewGitHubPersistence creates a new GitHubPersistence object
func NewGitHubPersistence(owner string, repo string, client *http.Client) *GitHubPersistence {
	return &GitHubPersistence{
		github.NewClient(client),
		owner,
		repo,
	}
}

// GitHubPersistence persists kickoffs on a GitHub repository's Issues
type GitHubPersistence struct {
	gitHubClient *github.Client
	owner        string
	repo         string
}

// Add saves a Kickoff as a GitHub issue
func (c GitHubPersistence) Add(ko *Kickoff) error {
	_, _, err := c.gitHubClient.Issues.Create(context.Background(), c.owner, c.repo, &github.IssueRequest{
		Title: github.String(ko.Title),
	})
	return err
}
