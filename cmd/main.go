package main

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	"log"
	"os"
)

func main() {
	engine := kickoff.NewEngine(
		kickoff.NewGitHubPersistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient()),
		handler.NewSequenceHandler(
			handler.NewTitleHandler(kickoff.NewTextQuestion("Title: ")),
			handler.NewBodyHandler(body),
		),

		kickoff.NewTextQuestioner(os.Stdin, os.Stdout),
	)

	err := engine.Start()
	if err != nil {
		log.Fatal(err)
	}
}
