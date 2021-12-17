package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const numDays = 80

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	fishTimerStr := strings.Split(string(input), ",")
	fishTimer := []int{}
	for _, s := range fishTimerStr {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Failed to convert %s to an integer: %v", s, err)
		}
		fishTimer = append(fishTimer, num)
	}

	for i := 0; i < numDays; i++ {
		numNew := 0
		newFishTimer := []int{}
		for _, f := range fishTimer {
			if f == 0 {
				numNew++
				newFishTimer = append(newFishTimer, 6)
			} else {
				newFishTimer = append(newFishTimer, f-1)
			}
		}
		for i := 0; i < numNew; i++ {
			newFishTimer = append(newFishTimer, 8)
		}
		fishTimer = newFishTimer
	}

	fmt.Printf("There are %d fish.\n", len(fishTimer))
}
