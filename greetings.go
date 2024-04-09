package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("Nombre Vacio")
	}
	message := fmt.Sprintf(randomGreetings(), name)
	return message, nil
}

func randomGreetings() string {

	formats := []string{
		"Hola %v! Bienvenido!",
		"Que bueno verte, %v",
		"Quien sos? %v",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
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
