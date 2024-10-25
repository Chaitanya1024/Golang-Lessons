package main

import "fmt"

//number sequence of specified length
func main()  {
	//initialsed with zero values
	//int -> 0, string -> "", bool -> false

	// var num[4]int
	//array length
	// fmt.Println(len(num))
	// num[0]=1
	// fmt.Println(num[0])
	// fmt.Println(num)

	//initialsed with false
	// var vals[4]bool
	// fmt.Println(vals)

	//initialised with empty string
	// var names[3]string
	// names[0]= "golang"
	// fmt.Println(names)

	//to decalre in single line
	// nums:=[3]int{1,2,3}

	// fmt.Println(nums)

	//2d array
	nums:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nums)

	//fixed size that are predictable
	//Memory optimizarion
	//Constant time access
}