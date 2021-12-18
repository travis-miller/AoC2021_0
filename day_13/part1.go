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
	dots := map[dot]bool{}
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
	applyFold(dots, folds[0])
	fmt.Printf("There are %d dots visible\n", len(dots))
}

type dot struct {
	x, y int
}

type fold struct {
	axis     string
	location int
}

func applyFold(dots map[dot]bool, f fold) {
	foldedDots := []dot{}
	for d := range dots {
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
		delete(dots, d)
		if f.axis == "x" {
			newDot := dot{f.location*2 - d.x, d.y}
			dots[newDot] = true
		}
		if f.axis == "y" {
			newDot := dot{d.x, f.location*2 - d.y}
			dots[newDot] = true
		}
	}
}
