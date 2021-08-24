package main

import (
	"fmt"

	pathseparator "github.com/alehpineda/go-bootcamp/01-basics/03-path-separator"
	changevalues "github.com/alehpineda/go-bootcamp/01-basics/05-change-values"
	shadowing "github.com/alehpineda/go-bootcamp/01-basics/04-shadowing"
	"github.com/alehpineda/go-bootcamp/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	var speed int
	fmt.Println(speed) // prints 0, go assignes a value of 0, python falsy
	pathseparator.Pathseparator()
	pathseparator.Blankpathseparator()
	shadowing.Shadowing()
	changevalues.TypeConversion()
}