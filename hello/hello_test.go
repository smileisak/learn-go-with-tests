package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ismail")
	want := "Hello, Ismail"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
