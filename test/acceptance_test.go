package kickoff_test

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/test/doubles"
	"github.com/gypsydave5/kickoff/test/random"
	"testing"
	"time"
)

func TestCreatingAnIssueWithATitle(t *testing.T) {
	ghKOClient := kickoff.NewGitHubPersistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient())
	question := random.RandomString()
	answer := random.RandomString()
	spyQuestioner := doubles.NewSpyQuestioner(answer)

	var ko = kickoff.NewEngine(
		ghKOClient,
		kickoff.NewTitleHandler(kickoff.NewTextQuestion(question)),
		spyQuestioner,
	)

	err := ko.Start()
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if len(spyQuestioner.Questions) != 1 {
		t.Errorf("Wrong number of questions asked: %d", len(spyQuestioner.Questions))
	}
	askedQuestion := spyQuestioner.Questions[0].String()
	if askedQuestion != question {
		t.Errorf("Wanted question %q to be asked, but it was %q", question, askedQuestion)
	}

	time.Sleep(time.Second * 5) // takes a while for the change to happen in GH

	ghc := github.NewClient(kickoff.NewGitHubOAuthHTTPClient())
	issues, _, err := ghc.Search.Issues(context.Background(), fmt.Sprintf("%s in:title is:issue repo:gypsydave5/kickoff", answer), nil)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if len(issues.Issues) != 1 {
		t.Log(issues.Issues)
		t.Error("Kickoff issue not created")
	}
	title := *(issues.Issues[0].Title)
	if title != answer {
		t.Errorf("expected title of %q but got %q", answer, title)
	}

}
