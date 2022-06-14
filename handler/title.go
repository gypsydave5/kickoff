package handler

import "github.com/gypsydave5/kickoff"

type TitleHandler struct {
	question kickoff.Question
}

func NewTitleHandler(question kickoff.Question) *TitleHandler {
	return &TitleHandler{question}
}

func (t TitleHandler) Handle(questioner kickoff.Questioner) (*kickoff.Kickoff, error) {
	answer, err := questioner.AskQuestion(t.question)
	if err != nil {
		return &kickoff.Kickoff{}, err
	}
	return kickoff.NewKickoff(answer.String()), nil
}
