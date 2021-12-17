package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	matchCount := 0
	match := map[int]bool{2: true, 3: true, 4: true, 7: true}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		if len(line) != 2 {
			log.Fatalf("Unexpeceted input for line %s", scanner.Text())
		}
		for _, d := range strings.Split(line[1], " ") {
			if match[len(d)] {
				matchCount++
			}
		}
	}

	fmt.Printf("Found %d matches\n", matchCount)
}
