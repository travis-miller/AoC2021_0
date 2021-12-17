package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()

	lines := []ventLine{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		if len(line) != 2 {
			log.Fatalf("Unexpected line format %v", line)
		}
		startStr := strings.Split(line[0], ",")
		if len(startStr) != 2 {
			log.Fatalf("Unexpected start coordinate format %v", startStr)
		}
		endStr := strings.Split(line[1], ",")
		if len(endStr) != 2 {
			log.Fatalf("Unexpected end coordinate format %v", endStr)
		}
		startInt := make([]int, 2)
		startInt[0], err = strconv.Atoi(startStr[0])
		if err != nil {
			log.Fatalf("Failed to convert %s to integer", startStr[0])
		}
		startInt[1], err = strconv.Atoi(startStr[1])
		if err != nil {
			log.Fatalf("Failed to convert %s to integer", startStr[1])
		}
		endInt := make([]int, 2)
		endInt[0], err = strconv.Atoi(endStr[0])
		if err != nil {
			log.Fatalf("Failed to convert %s to integer", endStr[0])
		}
		endInt[1], err = strconv.Atoi(endStr[1])
		if err != nil {
			log.Fatalf("Failed to convert %s to integer", endStr[1])
		}
		lines = append(lines, ventLine{
			start: point{startInt[0], startInt[1]},
			end:   point{endInt[0], endInt[1]},
		})
	}

	ventMap := map[point]int{}
	for _, l := range lines {
		if l.start.x == l.end.x {
			x := l.start.x
			minY := l.start.y
			if minY > l.end.y {
				minY = l.end.y
			}
			maxY := l.start.y
			if maxY < l.end.y {
				maxY = l.end.y
			}
			for y := minY; y <= maxY; y++ {
				ventMap[point{x, y}]++
			}
		}
		if l.start.y == l.end.y {
			y := l.start.y
			minX := l.start.x
			if minX > l.end.x {
				minX = l.end.x
			}
			maxX := l.start.x
			if maxX < l.end.x {
				maxX = l.end.x
			}
			for x := minX; x <= maxX; x++ {
				ventMap[point{x, y}]++
			}
		}
	}

	multiPoint := 0
	for _, v := range ventMap {
		if v > 1 {
			multiPoint++
		}
	}

	fmt.Printf("Found %d overlapping points.\n", multiPoint)
}

type ventLine struct {
	start point
	end   point
}

type point struct {
	x int
	y int
}
