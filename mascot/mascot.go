package mascot

import (
	"fmt"

	"rsc.io/quote"
)

// best mascot
func BestMascot() string {
	return "Gopher"
}

func runBestMascot() {
	fmt.Println(BestMascot())
	fmt.Println(quote.Go())
	var speed int
	fmt.Println(speed) // prints 0, go assignes a value of 0, python falsy
}