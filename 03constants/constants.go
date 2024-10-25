package main

import "fmt"

// name:="goland" shorthand not allowed outside main
// var name1 string = "golang" with var allowed

func main()  {
	const name string = "hello"
	// const name1 = "hello"
	const (
		port = 5000
		host = "localhost"
	)
	var port1 int = 5001
	port1 = 5002
	fmt.Println(port1)

}