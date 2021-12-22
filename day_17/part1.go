package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	var minX, maxX, minY, maxY int
	_, err = fmt.Sscanf(string(input), "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)
	if err != nil {
		log.Fatalf("Failed to scan input: %v", err)
	}
	highestY := 0
	target := targetArea{minX, maxX, minY, maxY}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= int(math.Abs(float64(minY))); y++ {
			tempY, success := reachesTarget(velocity{x, y}, target)
			if success && tempY > highestY {
				highestY = tempY
			}
		}
	}
	fmt.Printf("Highest Y: %d\n", highestY)
}

type velocity struct {
	x, y int
}

type targetArea struct {
	minX, maxX, minY, maxY int
}

func (t targetArea) containsPoint(p point) bool {
	return t.minX <= p.x && t.maxX >= p.x && t.minY <= p.y && t.maxY >= p.y
}

type point struct {
	x, y int
}

func reachesTarget(v velocity, t targetArea) (int, bool) {
	maxY := 0
	loc := point{}
	for tryAgain(loc, t) {
		loc, v = step(loc, v)
		if loc.y > maxY {
			maxY = loc.y
		}
		if t.containsPoint(loc) {
			return maxY, true
		}
	}
	return 0, false
}

func tryAgain(loc point, t targetArea) bool {
	return loc.x <= t.maxX && loc.y >= t.minY
}

func step(loc point, v velocity) (point, velocity) {
	newLoc := point{loc.x + v.x, loc.y + v.y}
	newVel := velocity{}
	switch {
	case v.x > 0:
		newVel.x = v.x - 1
	case v.x < 0:
		newVel.x = v.x + 1
	}
	newVel.y = v.y - 1
	return newLoc, newVel
}
