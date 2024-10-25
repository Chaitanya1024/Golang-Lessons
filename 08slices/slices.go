package main

import (
	"fmt"
	"slices"
)

//slice -> dynamic array
// most used construct in go
// + useful methods
func main() {
	//uninitialsed slice is nil
	// var nums []int

	// fmt.Println(nums==nil)
	// fmt.Println(len(nums ))


	// var nums = make([]int,0,5)
	//cap maximum element can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var nums = make([]int,0,5)
	// nums = append(nums, 2)
	// var nums2 = make([]int,len(nums))
	
	// //copy func
	// copy(nums2,nums)
	//  fmt.Println(nums,nums2)

	//slice operators

	// var nums = []int{1,2,3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:1])
	// fmt.Println(nums[1:])	

	//slice package
	var nums1 = []int{1,2}
	var nums2 = []int{1,3}
	fmt.Println(slices.Equal(nums1,nums2))
}
