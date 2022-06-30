package github

import (
	"context"
	"github.com/google/go-github/v45/github"
	"github.com/gypsydave5/kickoff"
	"net/http"
)

// Persistence creates a new GitHubPersistence object
func Persistence(owner string, repo string, client *http.Client) *GitHubPersistence {
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
func (c GitHubPersistence) Save(ko *kickoff.Kickoff) error {
	_, _, err := c.gitHubClient.Issues.Create(context.Background(), c.owner, c.repo, &github.IssueRequest{
		Title: github.String(ko.Title),
		Body:  github.String(ko.Body),
	})
	return err
}
