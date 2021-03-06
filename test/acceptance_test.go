package kickoff_test

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	github2 "github.com/gypsydave5/kickoff/persistence/github"
	"github.com/gypsydave5/kickoff/question"
	"github.com/gypsydave5/kickoff/test/random"
	"github.com/gypsydave5/kickoff/test/test_double"
	"testing"
	"time"
)

func TestCreatingAKickoffOnGitHub(t *testing.T) {
	ghKOClient := github2.NewPersistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient())
	q := random.String16()
	answer := random.String16()
	spyQuestioner := test_double.NewSpyQuestioner(answer)

	expectedBody := random.String16()
	seq := handler.NewSequence(
		handler.NewTitle(question.NewTextQuestion(q)),
		handler.NewBody(expectedBody),
	)

	var ko = kickoff.NewEngine(
		ghKOClient,
		seq,
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
	if askedQuestion != q {
		t.Errorf("Wanted question %q to be asked, but it was %q", q, askedQuestion)
	}

	time.Sleep(time.Second * 2) // takes a while for the change to happen in GH

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

	if issues.Issues[0].Body == nil {
		t.Errorf("No body")
	}
	body := *(issues.Issues[0].Body)
	if title != answer {
		t.Errorf("expected body of %q but got %q", expectedBody, body)
	}

}
