package main

import "fmt"

// pass by reference
func swap (a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 5
	b := 10
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	swap(&a, &b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Address of a and b
	var ptr1 *int = &a
	ptr2 := &b
	fmt.Println("Address of a:", ptr1)
	fmt.Println("Address of b:", ptr2)
}