package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Sam", "")
		want := "Hello, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {

		got := Hello("Sam", "Spanish")
		want := "Hola, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		
		got := Hello("Sam", "French")
		want := "Bonjour, Sam"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
