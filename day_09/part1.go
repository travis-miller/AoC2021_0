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
	heightmap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		newLevel := []int{}
		for _, n := range line {
			num, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatalf("Failed to convert %v to an integer: %v", n, err)
			}
			newLevel = append(newLevel, num)
		}
		heightmap = append(heightmap, newLevel)
	}

	riskLevel := 0
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			isRisk := true
			curLoc := heightmap[i][j]
			if i != 0 {
				if curLoc >= heightmap[i-1][j] {
					isRisk = false
				}
			}
			if i < len(heightmap)-1 {
				if curLoc >= heightmap[i+1][j] {
					isRisk = false
				}
			}
			if j != 0 {
				if curLoc >= heightmap[i][j-1] {
					isRisk = false
				}
			}
			if j < len(heightmap[i])-1 {
				if curLoc >= heightmap[i][j+1] {
					isRisk = false
				}
			}
			if isRisk {
				riskLevel += curLoc + 1
			}
		}
	}
	fmt.Printf("Risk level is %d\n", riskLevel)
}
