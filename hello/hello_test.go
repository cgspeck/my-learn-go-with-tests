package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, a, e string) {
		t.Helper()
		if a != e {
			t.Errorf("expected: %q, actual: %q", e, a)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		e := "Hello, Chris"
		a := Hello("Chris", "")
		assertCorrectMessage(t, a, e)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		e := "Hola, Elodie"
		a := Hello("Elodie", "es")
		assertCorrectMessage(t, a, e)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		e := "Bonjour, Jean Paul"
		a := Hello("Jean Paul", "fr")
		assertCorrectMessage(t, a, e)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		e := "Hello, World"
		a := Hello("", "")
		assertCorrectMessage(t, a, e)
	})
}
