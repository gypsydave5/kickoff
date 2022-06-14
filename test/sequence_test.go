package kickoff_test

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/test/test_double"
	"testing"
)

func TestNewSequence(t *testing.T) {
	seq := handler.NewSequenceHandler(
		handler.NewTitleHandler(kickoff.NewTextQuestion("hello")),
		handler.NewBodyHandler("body"),
	)
	sq := test_double.NewSpyQuestioner("but this")

	ko, _ := seq.Handle(sq)

	if ko.Title != "but this" {
		t.Errorf("Expected title %q but got %q", "but this", ko.Title)
	}

	if ko.Body != "body" {
		t.Errorf("Expected body %q but got %q", "body", ko.Body)
	}
}
