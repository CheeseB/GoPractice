package main

import "testing"

func TestHello(t *testing.T) { // number of subtest: 3

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sonia", "")
		want := "Hello, Sonia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("just say 'Hello, World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("French Man", "French")
		want := "Bonjour, French Man"
		assertCorrectMessage(t, got, want)
	})
}
