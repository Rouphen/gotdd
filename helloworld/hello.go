package main

import "fmt"

const spanish = "Spanish"

const enlishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello for testing TDD
func Hello(name, language string) string {

	prefix := enlishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	default:
		prefix = enlishHelloPrefix
	}

	if name == "" {
		name = "world"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
