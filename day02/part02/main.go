package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part02(inputFile string) (result string, ok bool) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open file %s: %v", inputFile, err)
	}
	defer f.Close()

	seen := map[string]bool{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		id := s.Text() // one id per line
		for pos := range id {
			truncatedID := id[:pos] + "_" + id[pos+1:]
			if seen[truncatedID] {
				return strings.Replace(truncatedID, "_", "", 1), true
			}
			seen[truncatedID] = true
		}
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	return "", false
}

func main() {
	result, ok := part02("input.txt")
	if ok {
		fmt.Printf("day02: %s\n", result)
	}
	log.Fatal("could not calculate result")
}
