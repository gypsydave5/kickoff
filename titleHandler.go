package kickoff

type TitleHandler struct {
	question string
}

func NewTitleHandler(question string) *TitleHandler {
	return &TitleHandler{question: question}
}

func (t TitleHandler) Handle(questioner Questioner) (*Kickoff, error) {
	title := questioner.AskQuestion(t.question)
	return NewKickoff(title), nil
}

func NewKickoff(title string) *Kickoff {
	return &Kickoff{title}
}
