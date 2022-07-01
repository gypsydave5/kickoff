package kickoff

import questioner2 "github.com/gypsydave5/kickoff/question"

// Persistence interface represents some external service where your kickoffs are stored once generated.
type Persistence interface {
	Save(kickoff *Kickoff) error
}

type Engine struct {
	persistence Persistence
	handler     InitialHandler
	questioner  questioner2.Input
}

func (e Engine) Start() error {
	ko, err := e.handler.Handle(e.questioner)
	if err != nil {
		return err
	}

	err = e.persistence.Save(ko)
	return err
}

func NewEngine(client Persistence, handler InitialHandler, questioner questioner2.Input) *Engine {
	return &Engine{client, handler, questioner}
}

type InitialHandler interface {
	Handle(questioner questioner2.Input) (*Kickoff, error)
}

type Handler interface {
	Handle(kickoff *Kickoff, questioner questioner2.Input) (*Kickoff, error)
}

// Kickoff represents a kickoff
type Kickoff struct {
	Title string
	Body  string
}

func (k *Kickoff) AddSection(line string) {
	k.Body = k.Body + "\n\n\n" + line
}

func NewKickoff(title string) *Kickoff {
	return &Kickoff{Title: title}
}
