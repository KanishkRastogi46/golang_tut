package main

import "fmt"

func total(args ...int) int {
	sum := 0
	for _ , val := range args {
		sum += val
	}
	return sum
}

func variadic(slc ...interface{}) {
	for _, val := range slc {
		fmt.Println(val)
	}
}

func main() {
	nums := []int{1,2,3,4,5}
	res := total(nums...)
	fmt.Println("Sum is", res)

	variadic(1, "Hello", 3.14)
}