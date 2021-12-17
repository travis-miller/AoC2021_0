package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	carbPosStr := strings.Split(string(input), ",")
	crabPos := []int{}
	for _, s := range carbPosStr {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert %s to an integer: %v", s, err)
		}
		crabPos = append(crabPos, num)
	}
	minPos := crabPos[0]
	maxPos := crabPos[0]
	for _, p := range crabPos {
		if minPos > p {
			minPos = p
		}
		if maxPos < p {
			maxPos = p
		}
	}

	minFuel := -1
	for i := minPos; i < maxPos+1; i++ {
		tempMinFuel := 0
		for _, p := range crabPos {
			tempMinFuel += int(math.Abs(float64(p - i)))
		}
		if tempMinFuel < minFuel || minFuel == -1 {
			minFuel = tempMinFuel
		}
	}

	fmt.Printf("Minumum fuel required %d.\n", minFuel)
}
