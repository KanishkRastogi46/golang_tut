package main

import "fmt"

func main() {
	// range function on array/slice provides both index and value
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	// range function on map provides key and value
	m := make(map[int]string)
	m[1] = "Kanishk"
	m[2] = "Rastogi"
	for key, value := range m {
		fmt.Println(key, value)
	}

	// range function on string provides index and unicode value
	str := "Go is a beautiful language"
	for i, c := range str {
		fmt.Println(i, string(c))
	}
}