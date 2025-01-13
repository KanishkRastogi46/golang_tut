package main

import "fmt"

// closure function to generate a sequence of integers
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// recursive function to calculate factorial of a number
func factorial(n int) int {
	if n == 0 { 
		return 1 
	}
	return n * factorial(n-1)
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())

	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)
	fmt.Println(factorial(num))
}