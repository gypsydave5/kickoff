package handler

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/questioner"
)

type TitleHandler struct {
	question questioner.Question
}

func NewTitle(question questioner.Question) *TitleHandler {
	return &TitleHandler{question}
}

func (t TitleHandler) Handle(questioner questioner.Questioner) (*kickoff.Kickoff, error) {
	answer, err := questioner.AskQuestion(t.question)
	if err != nil {
		return &kickoff.Kickoff{}, err
	}
	return kickoff.NewKickoff(answer.String()), nil
}
