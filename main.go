package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	githubAccessToken, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Fatal("GITHUB_TOKEN not set in environment")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)
	issue, response, err := c.Issues.Create(context.Background(), "gypsydave5", "kickoff", &github.IssueRequest{
		Title:     github.String("Test"),
		Body:      github.String("TEST"),
		Labels:    nil,
		Assignee:  nil,
		State:     nil,
		Milestone: nil,
		Assignees: nil,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(issue)
	fmt.Println(response)
}
