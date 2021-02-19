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
	fmt.Println("printing out struct")
	fmt.Printf("%+v", jim)

	// using function below, same as above but using a receiver
	fmt.Println("")
	fmt.Println("test code using pointers")
	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	// these two are equivalent. instead of using explicit pointers, go converts the person type into a pointer because the argument is *person in updateName
	jim.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// need to use a pointer here because go passes by value (and not by reference) so it'd be a copy of the value if we didn't use a pointer.
	// then below we get the value at that pointer address
	(*pointerToPerson).firstName = newFirstName
}