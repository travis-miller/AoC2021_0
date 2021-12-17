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
	var incCount int
	var prevDepth int = -1
	for scanner.Scan() {
		curDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Unexpected input. Cannot convert %s to an integer: %v", scanner.Text(), err)
		}
		if curDepth > prevDepth && prevDepth != -1 {
			incCount++
		}
		prevDepth = curDepth
	}

	fmt.Printf("There are %d increasing measurements.\n", incCount)
}
