package handler_test

import (
	"github.com/gypsydave5/kickoff/handler"
	"github.com/gypsydave5/kickoff/questioner"
	"github.com/gypsydave5/kickoff/test/random"
	"github.com/gypsydave5/kickoff/test/test_double"

	"testing"
)

func TestTitleHandler_Handle(t *testing.T) {
	question := random.String16()
	expectedTitle := random.String16()

	h := handler.NewTitle(questioner.NewTextQuestion(question))
	sq := test_double.NewSpyQuestioner(expectedTitle)

	ko, _ := h.Handle(sq)
	if ko.Title != expectedTitle {
		t.Errorf("Expected title %q but got %q", expectedTitle, ko.Title)
	}
}
