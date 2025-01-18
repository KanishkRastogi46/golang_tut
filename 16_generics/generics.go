package main

import "fmt"

// using generics with function
func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// using generics with struct that can accept id as int or string
type person[T int | string] struct {
	id T
	name string
	age int
}

type Stack[T int | string] struct {
	items []T
}
	

func main() {
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}
	truthTable := []bool{true, false, true, false}
	print(nums)
	print(names)
	print(truthTable)

	person1 := person[int]{id: 1, name: "Alice", age: 25}
	fmt.Println(person1)

	person2 := person[string]{id: "72751d", name: "Bob", age: 30}
	fmt.Println(person2)

	stack1 := Stack[int]{items: []int{1, 2, 3, 4, 5}}
	fmt.Println("Elements of stack are:",stack1.items)
	 
	stack2 := Stack[string]{items: []string{"Alice", "Bob", "Charlie"}}
	fmt.Println("Elements of stack are:",stack2.items)
}