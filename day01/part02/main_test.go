package main

import "testing"

func TestPart01(t *testing.T) {
	expected := 563
	have := part02("input.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", expected, have)
	}
}

func TestPart01Sample(t *testing.T) {
	expected := 10
	have := part02("input_sample.txt")
	if have != expected {
		t.Fatalf("Expected %d, but got %d", expected, have)
	}
}
