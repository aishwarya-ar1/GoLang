//GO does not support overloading and hence we use interfaces

package main

import "fmt"

//1 - defining a new custom type called bot; Other bots share implementation from this bot
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

//3 - if getGreeting becomes a type of bot, we are allowed to call it using the function printGreeting()
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//2 - any func names getGreeting and returing a string becomes type 'bot'
func (englishBot) getGreeting() string {
	return "Hey!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
