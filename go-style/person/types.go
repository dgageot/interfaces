package person

type Person struct{}

func New() *Person {
	return &Person{}
}

func (g *Person) Name() string {
	return "Bob"
}

func (g *Person) Age() int {
	return 52
}
