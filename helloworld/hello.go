package main

import "fmt"

const enlishHelloPrefix = "Hello, "

// Hello for testing TDD
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return enlishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
