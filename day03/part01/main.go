package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part01(inputFile string) int {
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
		// log.Printf("%d, %d - %dx%d", x, y, w, h)
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
	// fabric.print()
	return fabric.overlappingClaims()
}

func main() {
	fmt.Printf("day03: %d\n", part01("input.txt"))
}

type square struct {
	claimId string // claim id of '#' if claimed more than once
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
	p map[pos]square
}

func (f *fabric) addClaim(c claim) {
	if f.p == nil {
		f.p = make(map[pos]square)
	}
	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
			pos := pos{x: c.pos.x + i, y: c.pos.y + j}
			newSquare := square{claimId: "#", count: 1}
			if s, ok := f.p[pos]; ok {
				newSquare.count = s.count + 1
			} else {
				newSquare.claimId = strconv.Itoa(c.id)
			}
			f.p[pos] = newSquare
		}
	}
}

func (f *fabric) overlappingClaims() int {
	sum := 0
	for _, s := range f.p {
		if s.count > 1 {
			sum++
		}
	}
	return sum
}

func (f *fabric) _() { // print
	var xMax, yMax int
	for p := range f.p {
		if p.x > xMax {
			xMax = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
	}
	// fmt.Printf("xMax: %d, yMax: %d", xMax, yMax)
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			squareInch := f.p[pos{x, y}].claimId
			if squareInch == "" {
				squareInch = "."
			}
			fmt.Print(squareInch)
		}
		fmt.Println()
	}
}
