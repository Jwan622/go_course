package main

import "fmt"

type bot interface {
 	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}
	printGreeting(englishBot)
	printGreeting(spanishBot)
}

//instead of two methods with a diff argument type, use an interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting())
//}

func (englishBot) getGreeting() string {
  	//very custom logic for english greeting, imagine
	return "hi there"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}

// we can clean this duplication with interfaces
