package kickoff

func NewInitialHandler() {}

type Client interface {
	Add(kickoff *Kickoff) error
}

type Engine struct {
	client     Client
	handler    InitialHandler
	questioner Questioner
}

func (e Engine) Start() error {
	ko, err := e.handler.Handle(e.questioner)
	if err != nil {
		return err
	}

	err = e.client.Add(ko)
	return err
}

func NewEngine(client Client, handler InitialHandler, questioner Questioner) *Engine {
	return &Engine{client, handler, questioner}
}

type InitialHandler interface {
	Handle(questioner Questioner) (*Kickoff, error)
}

type Questioner interface {
	AskQuestion(string) string
}

type Kickoff struct {
	title string
}

func (k *Kickoff) Title() string {
	return k.title
}
