package changevalues


import "fmt"

// type(value)
func TypeConversion() {
	speed := 100
	force := 2.5

	// error mismatched types
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
