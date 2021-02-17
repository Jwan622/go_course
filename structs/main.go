package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
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

	jim := person{
		firstName: "Jim",
		lastName: "Wan",
		contactInfo: contactInfo{
			email: "someEmailForJim@gmail.com",
			zipCode: 90210,
		},
	}

	fmt.Println("")
	fmt.Printf("%+v", jim)
}