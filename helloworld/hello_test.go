package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessaeg := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessaeg(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessaeg(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessaeg(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Alex", "French")
		want := "Bonjour, Alex"

		assertCorrectMessaeg(t, got, want)
	})

}
