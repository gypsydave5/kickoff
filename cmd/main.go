package main

import (
	"github.com/gypsydave5/kickoff"
	"log"
	"os"
)

func main() {
	engine := kickoff.NewEngine(
		kickoff.NewGitHubPersistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient()),
		kickoff.NewTitleHandler("Title: "),
		kickoff.NewTextQuestioner(os.Stdin, os.Stdout),
	)

	err := engine.Start()
	if err != nil {
		log.Fatal(err)
	}
}
