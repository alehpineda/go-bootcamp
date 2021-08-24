package pathseparator

import (
	"fmt"
	"path"
)

// uppercase if you need it to be exported to main
func Pathseparator() {
	var dir, file string

	dir, file = path.Split("css/main.css")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}

func Blankpathseparator() {
	var file string

	_, file = path.Split("css/main.css")

	fmt.Println("file: ", file)
}