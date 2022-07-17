package main

import "testing"

func TestPart01Sample(t *testing.T) {
	expected := "fgij"
	have, _ := part02("input_sample.txt")
	if have != expected {
		t.Fatalf("Expected '%s', but got '%s'", expected, have)
	}
}

func TestPart01(t *testing.T) {
	expected := "hhvsdkatysmiqjxunezgwcdpr"
	have, _ := part02("input.txt")
	if have != expected {
		t.Fatalf("Expected %s, but got %s", expected, have)
	}
}
