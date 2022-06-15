package kickoff

import (
	"fmt"
	"github.com/gypsydave5/kickoff"
	"github.com/gypsydave5/kickoff/test/random"
	"strings"
	"testing"
)

func TestTextQuestioner_ReadingAndWriting(t *testing.T) {
	answer := random.String()
	question := random.String()

	out := new(strings.Builder)
	in := strings.NewReader(fmt.Sprintf("%s\n", answer))

	q := kickoff.NewTextQuestioner(in, out)

	gotAnswer, err := q.AskQuestion(kickoff.NewTextQuestion(question))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if gotAnswer.String() != answer {
		t.Errorf("Wanted %q got %q", answer, gotAnswer)
	}

	gotQuestion := out.String()

	if gotQuestion != question {
		t.Errorf("Wanted %q got %q", question, gotQuestion)
	}
}
