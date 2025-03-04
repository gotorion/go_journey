package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFommat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func Hello(name string) (string, error) {
	if name == " " {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFommat(), name)
	return message, nil
}