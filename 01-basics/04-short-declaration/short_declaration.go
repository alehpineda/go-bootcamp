package shortdeclaration

func Declaration() {
	// when to use a declaration?
	// if you don't know the initial value
	var score int
	// when you need a package scoped variables
	var version string
	// when you want to group variables together for greater readability
	var (
		// related:
		video string
		// closely related:
		duration int
		current int
	)
}

func ShortDeclaration() {
	// when to use short declaration
	// you know the initial value
	ancho, altura := 100, 100
	// to keep the code consice
	// for redeclaration
	ancho, color := 50, "red"
}