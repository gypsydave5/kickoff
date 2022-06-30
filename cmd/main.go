package main

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/handler/jira"
	"github.com/gypsydave5/kickoff/persistence/github"
	"github.com/gypsydave5/kickoff/questioner"
	"log"
	"os"
)

func main() {
	username := os.Getenv("JIRA_USER")
	key := os.Getenv("JIRA_KEY")
	jiraClient := jira.NewClient(jira.NewBasicAuthHTTP(username, key))
	engine := kickoff.NewEngine(
		github.Persistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient()),
		handler.NewSequence(
			jira.NewInitHandler(jiraClient),
			handler.NewBody(body),
		),

		questioner.NewTextQuestioner(os.Stdin, os.Stdout),
	)

	err := engine.Start()
	if err != nil {
		log.Fatal(err)
	}
}
