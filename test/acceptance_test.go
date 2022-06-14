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
	spyQuestioner := NewSpyQuestioner(answer)

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

type SpyQuesitoner struct {
	Answers   []kickoff.Answer
	Questions []kickoff.Question
}

func (sq *SpyQuesitoner) AskQuestion(question kickoff.Question) (kickoff.Answer, error) {
	sq.Questions = append(sq.Questions, question)
	answer := sq.Answers[0]
	sq.Answers = sq.Answers[1:]
	return answer, nil
}

func NewSpyQuestioner(answers ...string) *SpyQuesitoner {
	q := &SpyQuesitoner{}
	for _, answer := range answers {
		q.Answers = append(q.Answers, kickoff.NewTextAnswer(answer))
	}
	return q
}

func RandomString() string {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(rando.Intn(1000))
}
