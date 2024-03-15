package format

import (
	"fmt"

	"github.com/dgageot/interfaces/java-style/person"
)

func Age(p person.Person) string {
	return fmt.Sprintf("%d years old", p.Age())
}

func Name(p person.Person) string {
	return fmt.Sprintf(" Mr %s", p.Name())
}
