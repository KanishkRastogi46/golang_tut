package main

import "fmt"

// pass by reference
func swap (a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func changeVal (x *int) *int {
	*x = 100
	return x
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

	y := 101
	fmt.Println("y:", y)
	fmt.Println("Address of x is:",changeVal(&y))
	fmt.Println("y:", y)
	
	x := &y
	fmt.Println("Address of x is:", x)
	fmt.Println("Value of x is:", *x)
	fmt.Println((*x + 1))

	// pointer to pointer
	var z **int = &x
	fmt.Println(**z)
}