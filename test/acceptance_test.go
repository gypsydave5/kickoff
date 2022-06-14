package kickoff_test

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/gypsydave5/kickoff"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCreatingAnIssueWithATitle(t *testing.T) {
	ghKOClient := kickoff.NewGitHubPersistence("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient())
	question := RandomString()
	answer := RandomString()

	var ko = kickoff.NewEngine(
		ghKOClient,
		kickoff.NewTitleHandler(question),
		NewMockQuestioner(answer),
	)

	err := ko.Start()
	if err != nil {
		t.Errorf("error: %s", err)
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
}

type MockQuestioner struct {
	Answers   []kickoff.Answer
	Questions []kickoff.Question
}

func (mi MockQuestioner) AskQuestion(question kickoff.Question) (kickoff.Answer, error) {
	mi.Questions = append(mi.Questions, question)
	answer := mi.Answers[0]
	mi.Answers = mi.Answers[1:]
	return answer, nil
}

func NewMockQuestioner(answers ...string) *MockQuestioner {
	q := &MockQuestioner{}
	for _, answer := range answers {
		q.Answers = append(q.Answers, kickoff.NewTextAnswer(answer))
	}
	return q
}

func RandomString() string {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(rando.Intn(1000))
}
