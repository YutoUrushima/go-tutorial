package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
 	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Taking the long way, you might have written this as:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}