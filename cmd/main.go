package main

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/handler/jira"
	"github.com/gypsydave5/kickoff/persistence/github"
	"github.com/gypsydave5/kickoff/question"
	"log"
	"os"
)

func main() {
	username := os.Getenv("JIRA_USER")
	key := os.Getenv("JIRA_KEY")
	jiraClient := jira.NewClient("https://saltpayco.atlassian.net", jira.NewBasicAuthHTTP(username, key))
	engine := kickoff.NewEngine(
		github.NewPersistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient()),
		handler.NewSequence(
			jira.NewInitHandler(jiraClient),
			handler.NewBody(body),
		),

		question.NewCommandLine(os.Stdin, os.Stdout),
	)

	err := engine.Start()
	if err != nil {
		log.Fatal(err)
	}
}
