package main

import (
	"fmt"

	"github.com/dgageot/interfaces/java-style/format"
	"github.com/dgageot/interfaces/java-style/person"
)

func main() {
	g := person.New()
	fmt.Println(format.Name(g))
	fmt.Println(format.Age(g))
}
