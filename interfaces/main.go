package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	
	printGreeting(eb)
	printGreeting(sb)
}


func (eb englishBot) getGreeting() string {
	//  very custom logic for generating an English greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}