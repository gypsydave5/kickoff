package kickoff

import (
	"bufio"
	"fmt"
	"io"
)

type TextQuestioner struct {
	input  io.Reader
	output io.Writer
}

func (t TextQuestioner) AskQuestion(question string) (string, error) {
	_, err := fmt.Fprint(t.output, question)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(t.input)
	scanner.Scan()
	text := scanner.Text()
	return text, nil
}

func NewTextQuestioner(input io.Reader, output io.Writer) *TextQuestioner {
	return &TextQuestioner{input: input, output: output}
}
