package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scores := []int{}
	for scanner.Scan() {
		s := scoreLine(scanner.Text())
		if s != -1 {
			scores = append(scores, s)
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	mid := len(scores) / 2
	fmt.Printf("Mid score is %d\n", scores[mid])
}

var closingCharOpener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var charPoints = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func scoreLine(line string) int {
	chunks := []rune{}
	for _, r := range line {
		if opener, ok := closingCharOpener[r]; ok {
			if chunks[len(chunks)-1] == opener {
				chunks = chunks[:len(chunks)-1]
				continue
			}
			return -1
		}
		chunks = append(chunks, r)
	}
	score := 0
	for i := len(chunks) - 1; i >= 0; i-- {
		score = score*5 + charPoints[chunks[i]]
	}
	return score
}
