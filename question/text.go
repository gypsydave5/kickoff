package question

import (
	"bufio"
	"fmt"
	"io"
)

type CommandLine struct {
	input  io.Reader
	output io.Writer
}

func (t CommandLine) Ask(question Question) (Answer, error) {
	_, err := fmt.Fprint(t.output, question)
	if err != nil {
		return NewTextAnswer(""), err
	}
	scanner := bufio.NewScanner(t.input)
	scanner.Scan()
	text := scanner.Text()
	return NewTextAnswer(text), nil
}

func NewCommandLine(input io.Reader, output io.Writer) *CommandLine {
	return &CommandLine{input: input, output: output}
}

type TextAnswer struct {
	answer string
}

func NewTextAnswer(answer string) *TextAnswer {
	return &TextAnswer{answer: answer}
}

func (t *TextAnswer) String() string {
	return t.answer
}

type TextQuestion struct {
	question string
}

func NewTextQuestion(question string) *TextQuestion {
	return &TextQuestion{question: question}
}

func (t *TextQuestion) String() string {
	return t.question
}
