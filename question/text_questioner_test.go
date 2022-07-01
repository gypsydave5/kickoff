package question_test

import (
	"fmt"
	"github.com/gypsydave5/kickoff/question"
	"github.com/gypsydave5/kickoff/test/random"
	"strings"
	"testing"
)

func TestTextQuestioner_ReadingAndWriting(t *testing.T) {
	answer := random.String16()
	q := random.String16()

	out := new(strings.Builder)
	in := strings.NewReader(fmt.Sprintf("%s\n", answer))

	input := question.NewCommandLine(in, out)

	gotAnswer, err := input.Ask(question.NewTextQuestion(q))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if gotAnswer.String() != answer {
		t.Errorf("Wanted %q got %q", answer, gotAnswer)
	}

	gotQuestion := out.String()

	if gotQuestion != q {
		t.Errorf("Wanted %q got %q", q, gotQuestion)
	}
}
