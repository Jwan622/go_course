package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// one way to use structs
	alex := person{firstName: "Alex", lastName: "Wan"}
	fmt.Println(alex)

	// second way to deal with structs
	var alex2 person
	alex2.firstName = "Alex"
	alex2.lastName = "Wan22"
	fmt.Println(alex2)
	fmt.Printf("%+v", alex2)
}