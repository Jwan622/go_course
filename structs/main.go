package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Wan"}
	fmt.Println(alex)
}