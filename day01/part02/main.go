package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInputFromFile(inputFile string) []int {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open file %s: %v", inputFile, err)
	}
	defer f.Close()

	var input []int

	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err = fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}
	if err = s.Err(); err != nil {
		log.Fatalf("could read file %s: %v", inputFile, err)
	}
	return input
}

func part02(inputFile string) int {
	sum := 0
	seen := map[int]bool{0: true}

	for {
		for _, n := range getInputFromFile(inputFile) {
			sum += n
			if seen[sum] {
				return sum
			}
			seen[sum] = true
		}
	}
}

func main() {
	fmt.Printf("day01: %d\n", part02("input.txt"))
}
