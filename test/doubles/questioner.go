package doubles

import "github.com/gypsydave5/kickoff"

type SpyQuestioner struct {
	Answers   []kickoff.Answer
	Questions []kickoff.Question
}

func (sq *SpyQuestioner) AskQuestion(question kickoff.Question) (kickoff.Answer, error) {
	sq.Questions = append(sq.Questions, question)
	answer := sq.Answers[0]
	sq.Answers = sq.Answers[1:]
	return answer, nil
}

func NewSpyQuestioner(answers ...string) *SpyQuestioner {
	q := &SpyQuestioner{}
	for _, answer := range answers {
		q.Answers = append(q.Answers, kickoff.NewTextAnswer(answer))
	}
	return q
}
