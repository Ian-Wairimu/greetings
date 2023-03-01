package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in the message
	if name == "" {
		return "", errors.New("empty name")
	}
	//creating message using random greetings
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//Hellos return a map that associates each of the named people
//with a greeting message
func Hellos(names []string) (map[string]string, error) {
	//a map associate name with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

//init set initial values of variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of a set of greetings message
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	//return random selected message format by specifying
	// a random index from the slice of formats
	return formats[rand.Intn(len(formats))]
}
func Name() string {
	return "Ian Moon is my name"
}
