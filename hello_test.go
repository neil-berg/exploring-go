package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("handles english prefix", func(t *testing.T) {
		got := Hello("english", "Neil")
		want := "Hello, Neil"
		assertMessage(t, got, want)
	})

	t.Run("handles french prefix", func(t *testing.T) {
		got := Hello("french", "Neil")
		want := "Bonjour, Neil"
		assertMessage(t, got, want)
	})

	t.Run("handles spanish prefix", func(t *testing.T) {
		got := Hello("spanish", "Neil")
		want := "Hola, Neil"
		assertMessage(t, got, want)
	})
}
