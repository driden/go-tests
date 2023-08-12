package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nacho", "")
		want := "Hello, Nacho"

		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		assertCorrectMsg(
			t,
			Hello("Elodie", "Spanish"),
			"Hola, Elodie")
	})

	t.Run("Say hello in French", func(t *testing.T) {
		assertCorrectMsg(
			t,
			Hello("Elodie", "French"),
			"Bonjour, Elodie")
	})
}

func assertCorrectMsg(t testing.TB, got, wanted string) {
	t.Helper()

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}

}
