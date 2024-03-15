package main

import (
	"fmt"

	"github.com/dgageot/interfaces/go-style/format"
	"github.com/dgageot/interfaces/go-style/person"
)

func main() {
	g := person.New()
	fmt.Println(format.Name(g))
	fmt.Println(format.Age(g))
}
