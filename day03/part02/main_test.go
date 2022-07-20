package main

import "testing"

func TestPart01Sample(t *testing.T) {
	expected := 3 // claim with id = 3
	have := part02("input_sample.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", expected, have)
	}
}

func TestPart01(t *testing.T) {
	expected := 877
	have := part02("input.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", have, expected)
	}
}
