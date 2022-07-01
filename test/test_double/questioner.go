package test_double

import (
	"github.com/gypsydave5/kickoff/question"
)

type SpyQuestioner struct {
	Answers   []question.Answer
	Questions []question.Question
}

func (sq *SpyQuestioner) Ask(question question.Question) (question.Answer, error) {
	sq.Questions = append(sq.Questions, question)
	answer := sq.Answers[0]
	sq.Answers = sq.Answers[1:]
	return answer, nil
}

func NewSpyQuestioner(answers ...string) *SpyQuestioner {
	q := &SpyQuestioner{}
	for _, answer := range answers {
		q.Answers = append(q.Answers, question.NewTextAnswer(answer))
	}
	return q
}
