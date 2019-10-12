package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	name := "Chris"
	Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Greet(%v, %q) = %q, want %q", buffer, name, got, want)
	}
}
