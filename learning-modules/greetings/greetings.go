package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greet(name string) (string, error) {

	if name == "" {
		err := errors.New("name cannot be empty")

		return "", err
	}

	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func SayHelloToMultiplePeople(names []string) (map[string]string, error) {
	// If no names were received, return an error.
	if len(names) == 0 {
		return nil, errors.New("must pass at least one name")
	}

	// Create a map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received names, calling the function
	// to get a message for each name.
	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
