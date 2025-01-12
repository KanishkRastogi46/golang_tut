package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("x=",x)

	var y int32 = 10000
	fmt.Println("y=",y)

	var firstName string = "Kanishk"
	var lastName = "Rastogi"
	fmt.Println(firstName, lastName)

	var marks float32 = 98.5
	var percentage float64 = 98.5345
	fmt.Println("float32=",marks, "float64=",percentage)

	var truth bool = true
	var lie bool = false
	fmt.Println(truth, lie)

	// shorthand declaration of variables
	name := "Golang"
	fmt.Println(name)

	// contants
	const pi = 3.14159
	fmt.Println("const=",pi)
	const editor string = "VS Code"
	fmt.Println(editor)

	// multiple constant variables
	const (
		lang1 = "Go"
		lang2 = "Python"
		lang3 = "Java"
		lang4 = "JavaScript"
	)

	fmt.Println(lang1, lang2, lang3, lang4)
}