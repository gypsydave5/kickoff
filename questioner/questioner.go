package questioner

type Question interface {
	String() string
}

type Answer interface {
	String() string
}

type Questioner interface {
	Ask(Question) (Answer, error)
}
