package handler

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/questioner"
)

type Body struct {
	body string
}

func (b Body) Handle(kickoff *kickoff.Kickoff, _ questioner.Questioner) (*kickoff.Kickoff, error) {
	kickoff.AddSection(b.body)
	return kickoff, nil
}

func NewBody(body string) *Body {
	return &Body{body: body}
}
