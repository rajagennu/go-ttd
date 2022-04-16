package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Raja", "english")
		want := "Hello, Raja"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Salut, Raja' in french, when lang passed as french", func(t *testing.T) {
		got := Hello("Raja", "french")
		want := "Salut, Raja"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hola, Raja' in spansih, when lang passed as french", func(t *testing.T) {
		got := Hello("Raja", "spanish")
		want := "Hola, Raja"

		assertCorrectMessage(t, got, want)
	})

}
