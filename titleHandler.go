package kickoff

type TitleHandler struct {
	question string
}

func NewTitleHandler(question string) *TitleHandler {
	return &TitleHandler{question: question}
}

func (t TitleHandler) Handle(questioner Questioner) (Kickoff, error) {
	title, err := questioner.AskQuestion(t.question)
	if err != nil {
		return &SimpleKickoff{}, err
	}
	return NewKickoff(title), nil
}

func NewKickoff(title string) *SimpleKickoff {
	return &SimpleKickoff{title}
}
