package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		sum += errorScore(scanner.Text())
	}
	fmt.Printf("Total error score %d\n", sum)
}

var closingCharOpener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var closingCharPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func errorScore(line string) int {
	chunks := []rune{}
	for _, r := range line {
		if opener, ok := closingCharOpener[r]; ok {
			if chunks[len(chunks)-1] == opener {
				chunks = chunks[:len(chunks)-1]
				continue
			}
			return closingCharPoints[r]
		}
		chunks = append(chunks, r)
	}
	return 0
}
