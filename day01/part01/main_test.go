package main

import "testing"

func TestPart01(t *testing.T) {
	expected := 561
	have := part01("input.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", have, expected)
	}
}

func TestPart01Sample(t *testing.T) {
	expected := 0
	have := part01("input_sample.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", have, expected)
	}
}
