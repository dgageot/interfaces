package person

type Person interface {
	Name() string
	Age() int
}

type person struct{}

func New() Person {
	return &person{}
}

func (g *person) Name() string {
	return "Bob"
}

func (g *person) Age() int {
	return 52
}
