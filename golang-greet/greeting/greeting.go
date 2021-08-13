package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hail, %v! Well met!",
		"Great to see you, %v!",
	}

	return formats[rand.Intn(len(formats))]
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Error: empty name")
	}

	// TODO: Force first character in the name to upper case.

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
