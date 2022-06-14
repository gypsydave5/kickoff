package handler

import (
	"github.com/gypsydave5/kickoff"
)

type BodyHandler struct {
	body string
}

func (b BodyHandler) Handle(kickoff *kickoff.Kickoff, _ kickoff.Questioner) (*kickoff.Kickoff, error) {
	kickoff.Body = b.body
	return kickoff, nil
}

func NewBodyHandler(body string) *BodyHandler {
	return &BodyHandler{body: body}
}
