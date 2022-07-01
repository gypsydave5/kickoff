package handler

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/question"
)

type InitSequence struct {
	first kickoff.InitialHandler
	rest  []kickoff.Handler
}

func (s InitSequence) Handle(questioner question.Input) (*kickoff.Kickoff, error) {
	ko, _ := s.first.Handle(questioner)
	for _, h := range s.rest {
		ko, _ = h.Handle(ko, questioner)
	}

	return ko, nil
}

func NewSequence(first kickoff.InitialHandler, rest ...kickoff.Handler) *InitSequence {
	return &InitSequence{first: first, rest: rest}
}
