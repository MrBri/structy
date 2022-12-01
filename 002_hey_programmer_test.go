package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		out  string
	}{
		{name: "Teddy", out: "hey Teddy"}, {name: "Joey", out: "hey Joey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phrase := Greet(tt.name)
			if phrase != tt.out {
				t.Errorf("got = %s, want = %s", phrase, tt.out)
			}
		})
	}
}
