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
		fmt.Printf("Addr: %p\n", &dog)

		fmt.Println("")
	}
}