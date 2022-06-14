package handler

import "github.com/gypsydave5/kickoff"

type InitSequenceHandler struct {
	first kickoff.InitialHandler
	rest  []kickoff.Handler
}

func (s InitSequenceHandler) Handle(questioner kickoff.Questioner) (*kickoff.Kickoff, error) {
	ko, _ := s.first.Handle(questioner)
	for _, h := range s.rest {
		ko, _ = h.Handle(ko, questioner)
	}

	return ko, nil
}

func NewSequenceHandler(first kickoff.InitialHandler, rest ...kickoff.Handler) *InitSequenceHandler {
	return &InitSequenceHandler{first: first, rest: rest}
}
