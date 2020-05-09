package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ismail")
	got := buffer.String()
	want := "Hello, Ismail"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
