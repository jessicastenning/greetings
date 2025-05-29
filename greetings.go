package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name included then return an error
	if name == "" {
		return "", errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hail %v, we meet!",
	}
	return formats[rand.Intn(len(formats))]
}
