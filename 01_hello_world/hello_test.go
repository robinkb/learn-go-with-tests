package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		receiver string
		want     string
	}{
		{"say hello to a person",
			"Robin",
			"Hello, Robin"},

		{"say hello to the world when the receiver is empty",
			"",
			"Hello, world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.receiver); got != tt.want {
				t.Errorf("Hello() = %q, want %q", got, tt.want)
			}
		})
	}
}
