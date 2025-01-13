package main

import (
	"fmt"
	"slices"
)

func main() {
	// making a slices
	var num []int
	num = append(num, 1)
	fmt.Println(num)

	var nums = make([]int, 2, 5)
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	nums[0] = 1
	nums[1] = 2
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 7)
	fmt.Println("nums: ",nums)
	fmt.Println(cap(nums))

	// copy slice
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println("nums2:", nums2)

	// slice operator
	fmt.Println(nums[2:5])

	// check equality of slice
	fmt.Println(slices.Equal(nums, nums2))

	// 2d slice
	num3 := make([][]int, 2)
	num3[0] = []int{1, 2, 3}
	fmt.Println(num3)
}