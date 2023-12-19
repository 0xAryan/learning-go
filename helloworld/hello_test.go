package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Aryan", "")
		want := "Hello, Aryan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Aryan", "Spanish")
		want := "Hola, Aryan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Aryan", "French")
		want := "Bonjour, Aryan"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}