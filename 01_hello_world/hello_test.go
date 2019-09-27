package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"say hello to a person",
			args{"Robin", ""},
			"Hello, Robin"},

		{"say hello to the world when the receiver is empty",
			args{"", ""},
			"Hello, world"},

		{"say hello in Spanish",
			args{"Elodie", "Spanish"},
			"Hola, Elodie"},

		{"say hello in French",
			args{"François", "French"},
			"Bonjour, François"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.lang); got != tt.want {
				t.Errorf("Hello() = %q, want %q", got, tt.want)
			}
		})
	}
}
