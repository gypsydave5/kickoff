package kickoff

type TextQuestion struct {
	question string
}

func NewTextQuestion(question string) *TextQuestion {
	return &TextQuestion{question: question}
}

func (t *TextQuestion) String() string {
	return t.question
}

type TitleHandler struct {
	question Question
}

func NewTitleHandler(question string) *TitleHandler {
	return &TitleHandler{question: NewTextQuestion(question)}
}

func (t TitleHandler) Handle(questioner Questioner) (Kickoff, error) {
	answer, err := questioner.AskQuestion(t.question)
	if err != nil {
		return &SimpleKickoff{}, err
	}
	return NewKickoff(answer.String()), nil
}

func NewKickoff(title string) *SimpleKickoff {
	return &SimpleKickoff{title}
}
