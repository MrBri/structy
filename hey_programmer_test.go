package main

import "testing"

func TestGreet(t *testing.T) {
	phrase := Greet("Teddy")
	if phrase != "hey Teddy" {
		t.Errorf("Greet(\"Teddy\") = %s; want \"hey Teddy\"", phrase)
	}
	// t.Fatal("not implemented")
}
