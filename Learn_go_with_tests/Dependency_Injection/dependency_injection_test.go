package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sonia")

	got := buffer.String()
	want := "Hello, Sonia"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
