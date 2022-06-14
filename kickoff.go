package kickoff

type Client interface {
	Add(kickoff Kickoff) error
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
	Handle(questioner Questioner) (Kickoff, error)
}

type Kickoff interface {
	Title() string
	AddTitle(string)
}

type SimpleKickoff struct {
	title string
}

func (k *SimpleKickoff) AddTitle(s string) {
	k.title = s
}

func (k *SimpleKickoff) Title() string {
	return k.title
}
