package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part02(inputFile string) int {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open file %s: %v", inputFile, err)
	}
	defer f.Close()

	var fabric fabric

	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, x, y, w, h int
		_, err = fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal("error scanning input line")
		}
		fabric.addClaim(claim{
			id:     id,
			pos:    pos{x: x, y: y},
			width:  w,
			height: h,
		})
	}
	if err = s.Err(); err != nil {
		log.Fatal("error reading file")
	}
	fmt.Println(fabric.idCount)
	return fabric.nonOverlappingClaims()
}

type square struct {
	claimId []int
	count   int
}

type claim struct {
	id            int
	pos           pos
	width, height int
}

type pos struct {
	x, y int
}

type fabric struct {
	idCount int
	p       map[pos]square
}

func (f *fabric) addClaim(c claim) {
	if f.p == nil {
		f.p = make(map[pos]square)
		f.idCount = 0
	}
	f.idCount++

	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
			pos := pos{x: c.pos.x + i, y: c.pos.y + j}
			newSquare := square{claimId: []int{}, count: 1}
			if s, ok := f.p[pos]; ok {
				newSquare.claimId = append(s.claimId, c.id)
				newSquare.count = s.count + 1
			} else {
				newSquare.claimId = []int{c.id}
			}
			f.p[pos] = newSquare
		}
	}
}

func (f *fabric) nonOverlappingClaims() int {
	overlaps := make(map[int]bool)
	for _, s := range f.p {
		if s.count > 1 {
			for _, id := range s.claimId {
				overlaps[id] = true
			}
		}
		if _, ok := overlaps[s.claimId[0]]; !ok {
			overlaps[s.claimId[0]] = false
		}
	}
	for id, overlap := range overlaps {
		if !overlap {
			return id
		}
	}
	log.Fatal("cloud not find non-overlapping claim id")
	return -1
}

func main() {
	fmt.Printf("day03: %d\n", part02("input.txt"))
}
