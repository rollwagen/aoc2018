package main

import "testing"

func TestPart01Sample(t *testing.T) {
	expected := 4
	have := part01("input_sample.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", expected, have)
	}
}

func TestPart01(t *testing.T) {
	expected := 114946
	have := part01("input.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", have, expected)
	}
}
