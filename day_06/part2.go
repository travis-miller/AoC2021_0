package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const numDays = 256

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	fishTimerStr := strings.Split(string(input), ",")
	fishTimer := make([]int, 9)
	for _, s := range fishTimerStr {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert %s to an integer: %v", s, err)
		}
		if num > 6 {
			log.Fatalf("Unexpeceted input number %d. Should not be greater than 6", num)
		}
		fishTimer[num]++
	}

	for i := 0; i < numDays; i++ {
		newFishTimer := make([]int, 9)
		for j := 0; j < 8; j++ {
			newFishTimer[j] = fishTimer[j+1]
		}
		newFishTimer[6] += fishTimer[0]
		newFishTimer[8] += fishTimer[0]
		fishTimer = newFishTimer
	}

	totalFish := 0
	for _, f := range fishTimer {
		totalFish += f
	}
	fmt.Printf("There are %d fish.\n", totalFish)
}
