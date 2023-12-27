package main

import "testing"

func assertTestCase(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World!"
		assertTestCase(t, got, want)
	})

	t.Run("saying hello to Josh in English", func(t *testing.T) {
		var got string = Hello("Josh", "English")
		var want string = "Hello Josh!"
		assertTestCase(t, got, want)
	})

	t.Run("saying hello Bella in Spanish", func(t *testing.T) {
		got := Hello("Bella", "Spanish")
		want := "Hola Bella!"
		assertTestCase(t, got, want)
	})
}