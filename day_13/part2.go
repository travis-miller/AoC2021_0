package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var foldRegex = regexp.MustCompile(`fold along ([xy])=(\d+)`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	dots := dotSheet{}
	folds := []fold{}
	scanner := bufio.NewScanner(f)
	parseFold := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseFold = true
			continue
		}
		if parseFold {
			fs := foldRegex.FindStringSubmatch(line)
			if len(fs) != 3 {
				log.Fatalf("Line %q has an unexpected fold format.", line)
			}
			loc, err := strconv.Atoi(fs[2])
			if err != nil {
				log.Fatalf("Error converting %s to int: %v", fs[2], err)
			}
			folds = append(folds, fold{fs[1], loc})
		} else {
			d := strings.Split(line, ",")
			if len(d) != 2 {
				log.Fatalf("Line %q has an unexpected dots format.", line)
			}
			x, err := strconv.Atoi(d[0])
			if err != nil {
				log.Fatalf("Error converting %s to int: %v", d[0], err)
			}
			y, err := strconv.Atoi(d[1])
			if err != nil {
				log.Fatalf("Error converting %s to int: %v", d[1], err)
			}
			dots[dot{x, y}] = true
		}
	}
	for _, f := range folds {
		dots.applyFold(f)
	}
	fmt.Printf("%s", dots.display())
}

type dot struct {
	x, y int
}

type fold struct {
	axis     string
	location int
}

type dotSheet map[dot]bool

func (ds dotSheet) applyFold(f fold) {
	foldedDots := []dot{}
	for d := range ds {
		if f.axis == "x" {
			if d.x > f.location {
				foldedDots = append(foldedDots, d)
			}
		}
		if f.axis == "y" {
			if d.y > f.location {
				foldedDots = append(foldedDots, d)
			}
		}
	}

	for _, d := range foldedDots {
		delete(ds, d)
		if f.axis == "x" {
			newDot := dot{f.location*2 - d.x, d.y}
			ds[newDot] = true
		}
		if f.axis == "y" {
			newDot := dot{d.x, f.location*2 - d.y}
			ds[newDot] = true
		}
	}
}

func (ds dotSheet) display() string {
	var maxX, maxY int
	for d := range ds {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}
	var sb strings.Builder
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if ds[dot{x, y}] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
