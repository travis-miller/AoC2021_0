package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	octos := octopusGrid{}
	for scanner.Scan() {
		oct := []*octopus{}
		for _, n := range scanner.Text() {
			num, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatalf("Failed to convert %v to an integer: %v", n, err)
			}
			oct = append(oct, &octopus{num, false})
		}
		octos = append(octos, oct)
	}

	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += octos.stepAndCountFlashes()
	}
	fmt.Printf("Total flashes %d\n", totalFlashes)
}

type octopus struct {
	level   int
	flashed bool
}

type point struct {
	x, y int
}

type octopusGrid [][]*octopus

func (og octopusGrid) stepAndCountFlashes() int {
	flashedOct := []point{}
	for i := 0; i < len(og); i++ {
		for j := 0; j < len(og[i]); j++ {
			o := og[i][j]
			o.level++
			if o.level > 9 {
				o.flashed = true
				flashedOct = append(flashedOct, point{i, j})
			}
		}
	}
	for len(flashedOct) != 0 {
		flashedOct = og.spreadFlash(flashedOct)
	}
	return og.resetAndCountFlashes()
}

func (og octopusGrid) spreadFlash(flashedOcts []point) []point {
	newFlashed := []point{}
	for _, fo := range flashedOcts {
		if fo.x != 0 && fo.y != 0 {
			newFlashed = og.incrementLevel(point{fo.x - 1, fo.y - 1}, newFlashed)
		}
		if fo.x != 0 {
			newFlashed = og.incrementLevel(point{fo.x - 1, fo.y}, newFlashed)
		}
		if fo.x != 0 && fo.y != len(og[fo.x])-1 {
			newFlashed = og.incrementLevel(point{fo.x - 1, fo.y + 1}, newFlashed)
		}
		if fo.y != 0 {
			newFlashed = og.incrementLevel(point{fo.x, fo.y - 1}, newFlashed)
		}
		if fo.y != len(og[fo.x])-1 {
			newFlashed = og.incrementLevel(point{fo.x, fo.y + 1}, newFlashed)
		}
		if fo.x != len(og)-1 && fo.y != 0 {
			newFlashed = og.incrementLevel(point{fo.x + 1, fo.y - 1}, newFlashed)
		}
		if fo.x != len(og)-1 {
			newFlashed = og.incrementLevel(point{fo.x + 1, fo.y}, newFlashed)
		}
		if fo.x != len(og)-1 && fo.y != len(og[fo.x])-1 {
			newFlashed = og.incrementLevel(point{fo.x + 1, fo.y + 1}, newFlashed)
		}
	}
	return newFlashed
}

func (og octopusGrid) incrementLevel(p point, flashedOcts []point) []point {
	o := og[p.x][p.y]
	o.level++
	if o.level > 9 && !o.flashed {
		o.flashed = true
		flashedOcts = append(flashedOcts, p)
	}
	return flashedOcts
}

func (og octopusGrid) resetAndCountFlashes() int {
	flashCount := 0
	for _, row := range og {
		for _, p := range row {
			if p.level > 9 {
				p.level = 0
				p.flashed = false
				flashCount++
			}
		}
	}
	return flashCount
}
