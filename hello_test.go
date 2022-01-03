package main

import "testing"

func TestHello(t *testing.T) {
	asserCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello with name", func(t *testing.T) {
		got := Hello("Daniel", "")
		want := "Hello, Daniel"
		asserCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, Worl' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Antony", "Spanish")
		want := "Hola, Antony"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Johny", "French")
		want := "Bonjour, Johny"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in Ukrainian", func(t *testing.T) {
		got := Hello("Данило", "Ukrainian")
		want := "Вітаю, Данило"
		asserCorrectMessage(t, got, want)
	})
}
