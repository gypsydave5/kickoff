package github

import (
	"context"
	"github.com/google/go-github/v45/github"
	"github.com/gypsydave5/kickoff"
	"net/http"
)

// NewPersistence creates a new Persistence object
func NewPersistence(owner string, repo string, client *http.Client) *Persistence {
	return &Persistence{
		gitHubGateway: Gateway{
			gitHubClient: github.NewClient(client),
			owner:        owner,
			repo:         repo,
		},
	}
}

// Persistence persists kickoffs on a GitHub repository's Issues
type Persistence struct {
	gitHubGateway Gateway
}

// Save saves a Kickoff as a GitHub issue
func (c Persistence) Save(ko *kickoff.Kickoff) error {
	return c.gitHubGateway.CreateIssue(&github.IssueRequest{
		Title: github.String(ko.Title),
		Body:  github.String(ko.Body),
	})
}

type Gateway struct {
	owner        string
	repo         string
	gitHubClient *github.Client
}

func (g Gateway) CreateIssue(i *github.IssueRequest) error {
	_, _, err := g.gitHubClient.Issues.Create(context.Background(), g.owner, g.repo, i)
	return err
}
