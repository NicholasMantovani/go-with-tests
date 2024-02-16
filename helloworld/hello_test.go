package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nicholas", "")
		want := "Hello, Nicholas"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
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
		got := Hello("Nicholas", "French")
		want := "Bonjour, Nicholas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in italian", func(t *testing.T) {
		got := Hello("Nicholas", "Italian")
		want := "Ciao, Nicholas"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // If this is present it prints in the fatal log the go line where this function is called instead of printing where the test fail (t.Errorf...)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
