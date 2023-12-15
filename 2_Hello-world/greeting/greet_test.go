package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aryan")
	want := "Hello, Aryan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}