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

	var sum int

	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err = fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}

	if err = s.Err(); err != nil {
		log.Fatalf("could read file %s: %v", inputFile, err)
	}
	return sum
}

func main() {
	fmt.Printf("day01: %d\n", part01("input.txt"))
}
