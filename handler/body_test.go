package handler_test

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/test/test_double"
	"testing"
)

func TestBodyHandler(t *testing.T) {
	h := handler.NewBody("body")
	sq := test_double.NewSpyQuestioner()
	initalKickOff := kickoff.NewKickoff("")

	ko, _ := h.Handle(initalKickOff, sq)

	if len(sq.Questions) != 0 {
		t.Errorf("Expected no questions to be asked, but questions were asked...")
	}

	if ko.Body != "body" {
		t.Errorf("Expected body %q but got %q", "body", ko.Body)
	}
}
