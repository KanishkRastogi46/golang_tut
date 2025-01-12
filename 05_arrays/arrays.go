package main

import "fmt"

func main() {
	var name[4] string
	name[0] = "John"
	name[1] = "Paul"
	name[2] = "George"
	name[3] = "Ringo"
	fmt.Println(name)
	fmt.Println("Length of name array is", len(name))

	for i:= range len(name) {
		fmt.Println("Name", i, "is", name[i])
	}

	var num[3] int = [3]int{1, 2, 3}
	fmt.Println(num)
}