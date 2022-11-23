package main

import "fmt"

// Hello for testing TDD
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return "Hello, " + name
}

func main() {
	fmt.Println(Hello(""))
}
