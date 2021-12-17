package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var bitCounter []int
	var lineCount int
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if bitCounter == nil {
			bitCounter = make([]int, len(line))
		}
		for pos, bit := range line {
			switch bit {
			case '0':
				// Do nothing
			case '1':
				bitCounter[pos]++
			default:
				log.Fatalf("Unknown bit character %v", bit)
			}
		}
	}
	var gamma, epsilon int
	for i, b := range bitCounter {
		if b > lineCount/2 {
			gamma += int(math.Pow(2, float64(len(bitCounter)-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(len(bitCounter)-i-1)))
		}
	}
	fmt.Printf("Gamma: %d Epsilon: %d Power: %d\n", gamma, epsilon, gamma*epsilon)
}
