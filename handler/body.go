package handler

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/question"
)

type Body struct {
	body string
}

func (b Body) Handle(kickoff *kickoff.Kickoff, _ question.Input) (*kickoff.Kickoff, error) {
	kickoff.AddSection(b.body)
	return kickoff, nil
}

func NewBody(body string) *Body {
	return &Body{body: body}
}
