package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age int
}

func main() {
	// this has a pointer-based bug
	jackie := Dog{
		Name: "Jeffyyyyyy",
		Age: 19,
	}
	fmt.Printf("Jackie Addr: %p\n", &jackie)

	sammy := Dog{
		Name: "Sammyyyyyy",
		Age: 10,
	}
	fmt.Printf("Sammy Addr: %p\n", &sammy)

	dogs := []Dog{jackie, sammy}
	fmt.Println("")

	for _, dog := range dogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		// so this points to the address of the dog variable, which is the same, but the value of the variable is changing.
		fmt.Printf("Addr: %p\n", &dog)
		fmt.Printf("Value of Dog: %+v\n", dog)

		fmt.Println("")
	}

	// this is the bug
	allDogs := []*Dog{}

	for _, dog := range dogs {
		allDogs = append(allDogs, &dog)
	}

	for _, dog := range allDogs {
		// this will print out the dog that was last iterated over because we're appending the address of the dog that was last iterated over
		// I end up with a slice where every element has the same address. This address is pointing to a copy of the last value that we iterated over. Yikes!!
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)

		//above returns this:
		//Name: Sammy Age: 10
		//Name: Sammy Age: 10
	}


	// lets do this using pointers

	fmt.Println("Lets do this the right way...")
	jackiePointer := &Dog{
		Name: "Jackie",
		Age: 19,
	}

	fmt.Printf("Jackie Addr: %p\n", jackiePointer)

	sammyPointer := &Dog{
		Name: "Sammy",
		Age: 10,
	}

	fmt.Printf("Sammy Addr: %p\n\n", sammyPointer)

	// this is an array of pointers, same as above
	dogs1 := []*Dog{jackiePointer, sammyPointer}

	for _, dog := range dogs1 {
		// This time we create a slice of pointers to Dog values. When we iterate over this slice, the value of the dog
		// variable is the address of each Dog value we stored in the slice. Instead of creating two extra copies of each
		// Dog value, we are using the same initial Dog value we created with the composite literal.
		// Special note: When the slice is a collection of Dog values or a collection of pointers to Dog values, the range loop is the same.
		// Go handles access to the Dog value regardless of whether we are using a pointer or not. This is awesome but can sometimes lead to a bit of confusion. At least it was for me in the beginning.
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n\n", dog)
	}
}