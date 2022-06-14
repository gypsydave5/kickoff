package kickoff_test

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/test/random"
	"github.com/gypsydave5/kickoff/test/test_double"
	"testing"
)

func TestTitleHandler_Handle(t *testing.T) {
	question := random.RandomString()
	expectedTitle := random.RandomString()

	h := handler.NewTitleHandler(kickoff.NewTextQuestion(question))
	sq := test_double.NewSpyQuestioner(expectedTitle)

	ko, _ := h.Handle(sq)
	if ko.Title != expectedTitle {
		t.Errorf("Expected title %q but got %q", expectedTitle, ko.Title)
	}
}
