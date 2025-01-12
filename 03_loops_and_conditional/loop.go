package main

import "fmt"

func main() {
	// for loop with break
	fmt.Println("for loop:")
	for i:= 0; i < 5; i++ {
		if i == 4 { break }
		fmt.Println(i)
	}

	fmt.Println()

	// while loop with continue
	fmt.Println("while loop:")
	var j = 0
	for j < 5 {
		fmt.Println(j)
		j++
		if j == 3 { continue }
	}

	fmt.Println()

	// for using range keyword from 0 to 9
	// introduced in go 1.22
	fmt.Println("for using range keyword:")
	for i:= range 10 {
		fmt.Println(i)
	}
}

// in go there is only one loop construct, the for loop