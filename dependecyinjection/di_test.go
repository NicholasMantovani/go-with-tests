package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Nic")

	got := buffer.String()
	want := "Hello, Nic"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
