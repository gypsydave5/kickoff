package kickoff

// Persistence interface represents some external service where your kickoffs are stored once generated.
type Persistence interface {
	Add(kickoff *Kickoff) error
}

type Engine struct {
	persistence Persistence
	handler     InitialHandler
	questioner  Questioner
}

func (e Engine) Start() error {
	ko, err := e.handler.Handle(e.questioner)
	if err != nil {
		return err
	}

	err = e.persistence.Add(ko)
	return err
}

func NewEngine(client Persistence, handler InitialHandler, questioner Questioner) *Engine {
	return &Engine{client, handler, questioner}
}

type InitialHandler interface {
	Handle(questioner Questioner) (*Kickoff, error)
}

type Handler interface {
	Handle(kickoff *Kickoff, questioner Questioner) (*Kickoff, error)
}

// Kickoff represents a kickoff
type Kickoff struct {
	Title string
	Body  string
}

func NewKickoff(title string) *Kickoff {
	return &Kickoff{Title: title}
}
