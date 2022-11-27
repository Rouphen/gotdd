package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chirs")

	got := buffer.String()
	want := "Hello, Chirs"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
