package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create a map
	var mpp = make(map[int]int)
	mpp[1] = 10
	mpp[2] = 20
	fmt.Println(mpp)
	fmt.Println(len(mpp))

	// map without make function
	var mp map[int]string = map[int]string{1:"kanishk", 2:"kumar"}
	fmt.Println(mp)
	val, exists := mp[2]		// val will contain the value of key 2 and exists will be true if key exists
	if exists {
		fmt.Println(val)
	} else {
		fmt.Println("Key not found")
	}

	// checking equality of maps
	m1 := make(map[int]int)
	fmt.Println(maps.Equal(mpp, m1))
}