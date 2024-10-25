package main

import (
	"fmt"
	// "time"
)

func main() {
	//simpe switch
	// i := 1

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")

	// }
	
	//Multiple condition switch

	// switch time.Now().Month(){
	// case time.April, time.May:
	// 	fmt.Println("Summer")
	// default:
	// 	fmt.Println("Workday")
	// }

	//type switch 
	whoAmI:= func(i interface{}){
		switch t:= i.(type){
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("other",t)
		}
	}

	whoAmI(60.2)


}