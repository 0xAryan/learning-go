package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hi there ! \n"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}