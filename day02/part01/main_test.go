package main

import "testing"

func TestPart01(t *testing.T) {
	expected := 8398
	have := part01("input.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", have, expected)
	}
}

func TestPart01Sample(t *testing.T) {
	expected := 12 // 4 * 3 = 12
	have := part01("input_sample.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", have, expected)
	}
}
