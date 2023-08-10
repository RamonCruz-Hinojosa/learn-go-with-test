package hello_world

//writing a test requires the file to have the _test.go in the file name
//test func must start with the word Test
//test func takes only one argument t *testing.T
//must import testing functionality before the func will work

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("chris", "French")
		want := "Bonjour, chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run()
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
