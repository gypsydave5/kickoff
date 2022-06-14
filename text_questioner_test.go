package kickoff

import (
	"strings"
	"testing"
)

func TestTextQuestioner_ReadingAndWriting(t *testing.T) {
	out := new(strings.Builder)
	in := strings.NewReader("answer\n")

	q := NewTextQuestioner(in, out)
	gotAnswer, err := q.AskQuestion(NewTextQuestion("question"))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if gotAnswer.String() != "answer" {
		t.Errorf("Wanted %q got %q", "answer", gotAnswer)
	}

	gotQuestion := out.String()

	if gotQuestion != "question" {
		t.Errorf("Wanted %q got %q", "question", gotQuestion)
	}
}
