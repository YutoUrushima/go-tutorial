package greetings

import "fmt"

// Hello returns a greeting for named person
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Taking the long way, you might have written this as:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}