package kickoff_test

import (
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/test/doubles"
	"github.com/gypsydave5/kickoff/test/random"
	"testing"
)

func TestTitleHandler_Handle(t *testing.T) {
	question := random.RandomString()
	expectedTitle := random.RandomString()

	h := kickoff.NewTitleHandler(kickoff.NewTextQuestion(question))
	sq := doubles.NewSpyQuestioner(expectedTitle)

	ko, _ := h.Handle(sq)
	if ko.Title != expectedTitle {
		t.Errorf("Expected title %q but got %q", expectedTitle, ko.Title)
	}
}
