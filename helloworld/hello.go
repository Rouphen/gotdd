package main

import "fmt"

// function Hello() for testing TDD
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return "Hello, " + name
}

func main() {
	fmt.Println(Hello(""))
}
