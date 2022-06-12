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
	ghKOClient := kickoff.NewGitHubKickoffClient("gypsydave5", "kickoff", kickoff.NewGitHubOAuthHTTPClient())
	question := RandomString()
	answer := RandomString()

	var ko = kickoff.NewEngine(
		ghKOClient,
		kickoff.NewTitleHandler(question),
		NewMockInput(answer),
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

type MockInput struct {
	Answers   []string
	Questions []string
}

func (mi MockInput) AskQuestion(question string) string {
	mi.Questions = append(mi.Questions, question)
	answer := mi.Answers[0]
	mi.Answers = mi.Answers[1:]
	return answer
}

func NewMockInput(answers ...string) *MockInput {
	return &MockInput{
		Answers: answers,
	}
}

func RandomString() string {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(rando.Intn(1000))
}

func main() {
	fmt.Println(RandomString())
}
