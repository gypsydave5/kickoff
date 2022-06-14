package kickoff

type TitleHandler struct {
	question Question
}

func NewTitleHandler(question Question) *TitleHandler {
	return &TitleHandler{question}
}

func (t TitleHandler) Handle(questioner Questioner) (*Kickoff, error) {
	answer, err := questioner.AskQuestion(t.question)
	if err != nil {
		return &Kickoff{}, err
	}
	return NewKickoff(answer.String()), nil
}
