package main

import (
	"fmt"
	"github.com/alehpineda/go-bootcamp/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	var speed int
	fmt.Println(speed) // prints 0, go assignes a value of 0, python falsy
}