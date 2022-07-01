package handler

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/question"
)

type TitleHandler struct {
	question question.Question
}

func NewTitle(question question.Question) *TitleHandler {
	return &TitleHandler{question}
}

func (t TitleHandler) Handle(questioner question.Input) (*kickoff.Kickoff, error) {
	answer, err := questioner.Ask(t.question)
	if err != nil {
		return &kickoff.Kickoff{}, err
	}
	return kickoff.NewKickoff(answer.String()), nil
}
