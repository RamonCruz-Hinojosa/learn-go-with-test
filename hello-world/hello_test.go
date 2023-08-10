package hello_world

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

//writing a test requires the file to have the _test.go in the file name
//test func must start with the word Test
//test func takes only one argument t *testing.T
//must import testing functionality before the func will work
