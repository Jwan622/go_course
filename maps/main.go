package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#745someHex",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	colors["blue"] = "someOtherHexValue"
	fmt.Println(colors)
	delete(colors,"blue")
	fmt.Println(colors)

	// iterate through a map
	fmt.Println("about to iterate through colors")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color:", color, "hex:", hex)
	}
}