package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part01(inputFile string) int {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open file %s: %v", inputFile, err)
	}
	defer f.Close()

	var doubles, triples int

	s := bufio.NewScanner(f)
	for s.Scan() {
		count := countRunes(s.Text())

		var hasDouble, hasTriple bool
		for _, c := range count {
			switch c {
			case 2:
				hasDouble = true
			case 3:
				hasTriple = true
			}
		}
		if hasDouble {
			doubles++
		}
		if hasTriple {
			triples++
		}
	}
	if err = s.Err(); err != nil {
		log.Fatal("Error reading file")
	}

	return doubles * triples
}

func countRunes(text string) map[rune]int {
	count := map[rune]int{}
	for _, r := range text {
		count[r]++
	}
	return count
}

func main() {
	fmt.Printf("day02: %d\n", part01("input.txt"))
}
