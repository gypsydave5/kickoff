package question

type Question interface {
	String() string
}

type Answer interface {
	String() string
}

type Input interface {
	Ask(Question) (Answer, error)
}
