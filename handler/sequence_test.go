package handler_test

import (
	h "github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/questioner"
	"github.com/gypsydave5/kickoff/test/test_double"
	"testing"
)

func TestNewSequence(t *testing.T) {
	body := "body"
	seq := h.NewSequence(
		h.NewTitle(questioner.NewTextQuestion("hello")),
		h.NewBody(body),
	)
	sq := test_double.NewSpyQuestioner("but this")

	ko, _ := seq.Handle(sq)

	if ko.Title != "but this" {
		t.Errorf("Expected title %q but got %q", "but this", ko.Title)
	}

	expectedBody := "\n\n\nbody"
	if ko.Body != expectedBody {
		t.Errorf("Expected body %q but got %q", expectedBody, ko.Body)
	}
}
