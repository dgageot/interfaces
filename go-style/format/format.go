package format

import "fmt"

type Ager interface {
	Age() int
}

func Age(p Ager) string {
	return fmt.Sprintf("%d years old", p.Age())
}

type Namer interface {
	Name() string
}

func Name(p Namer) string {
	return fmt.Sprintf(" Mr %s", p.Name())
}
