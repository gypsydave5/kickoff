package kickoff

import (
	"context"
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
