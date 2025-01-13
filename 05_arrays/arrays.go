package main

import "fmt"

func main() {
	// array declaration and initialization
	var name [4]string
	name[0] = "John"
	name[1] = "Paul"
	name[2] = "George"
	name[3] = "Ringo"
	fmt.Println(name)
	fmt.Println("Length of name array is", len(name))

	// iterate over array
	for i:= range len(name) {
		fmt.Println("Name", i, "is", name[i])
	}

	// array declaration and initialization
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println(nums)

	// iterate over array and insert value
	var num	[5]int
	fmt.Println(num)
	for i:= range(len(num)) {
		fmt.Println("Insert", i, "number in array num: ")
		fmt.Scan(&num[i])
	}
	fmt.Println(num)

	// 2d array
	var arr2d [2][2]int
	fmt.Println(arr2d)
}