package test_double

import (
	"github.com/gypsydave5/kickoff/questioner"
)

type SpyQuestioner struct {
	Answers   []questioner.Answer
	Questions []questioner.Question
}

func (sq *SpyQuestioner) AskQuestion(question questioner.Question) (questioner.Answer, error) {
	sq.Questions = append(sq.Questions, question)
	answer := sq.Answers[0]
	sq.Answers = sq.Answers[1:]
	return answer, nil
}

func NewSpyQuestioner(answers ...string) *SpyQuestioner {
	q := &SpyQuestioner{}
	for _, answer := range answers {
		q.Answers = append(q.Answers, questioner.NewTextAnswer(answer))
	}
	return q
}
