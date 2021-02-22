package pointers

import (
	"fmt"
	"strconv"
)

func main() {
	//https://medium.com/codezillas/uh-ohs-in-go-slice-of-pointers-c0a30669feee
	// Declare a list/slice of string pointers
	listOfNumberStrings := []*string{}

	// Forward declaration of the variable we will store the data in before adding to slice
	var numberString string

	// Loop from 0 to 9
	for i := 0; i < 10; i++ {
		// Format the string with '#' prefixed to the number
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
		fmt.Printf("Adding number %s to the slice\n", numberString)
		// Add number string to the slice
		listOfNumberStrings = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings {
		// this will just print out #9. why? read above post. In sjhort, when we add the address of &numberString, it's the address of the variable
		// so the listOfNumberStrings just has 9 addresses of the same address. And the last iteration points 9 at that address.
		fmt.Printf("%s\n", *n)
	}


	//this is the fix
	listOfNumberStrings1 := []*string{}

	for i := 0; i < 10; i++ {
		// we now declare the variable inside so it gets a new address for every number
		// What this accomplishes is, that at every iteration through the loop we are forced to re-declare the variable numberString on the stack, thereby giving it a new unique memory address.
		var numberString string
		// this is a format string
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
		listOfNumberStrings1 = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings1 {
		fmt.Printf("%s\n", *n)
	}

	return
}
