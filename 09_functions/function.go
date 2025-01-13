package main

import "fmt"

// function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello Kanishk")
}

// normal function
func add(a int, b int) int {
	return a + b
}

// return multiple values
func swap(str1 string, str2 string) (string, string) {
	return str2, str1
}

// in the return type we specify the varibales that we want to return
func namedReturn(a int) (x int, y int) {
	x = a * 10
	y = a * 20
	return
}

func funcAsRaram(fn func(a int, b int)int) int {
	return fn(10, 20)
}

func main() {
	// function with no parameters and no return value
	sayHello()

	// function with parameters and return value
	sum := add(1, 2)
	fmt.Println("Sum:", sum)

	// function with multiple return values
	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	// function with named return values
	x, y := namedReturn(10)
	fmt.Println(x, y)

	ans := funcAsRaram(add)
	fmt.Println("funcAsRaram(add):", ans)

	// variable of function type
	function := func(a int, b int, c int) int {
		return a + b + c
	}
	fmt.Println(function(1, 2, 3))
}