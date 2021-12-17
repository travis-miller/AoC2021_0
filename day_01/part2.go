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
	prevWindow := make([]int, 3, 3)
	curWindow := make([]int, 3, 3)

	for i := 0; i < 3; i++ {
		if ok := scanner.Scan(); !ok {
			break
		}
		curDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Unexpected input. Cannot convert %s to an integer: %v", scanner.Text(), err)
		}
		curWindow[i] = curDepth
	}
	for scanner.Scan() {
		copyWindow(curWindow, prevWindow)
		curDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Unexpected input. Cannot convert %s to an integer: %v", scanner.Text(), err)
		}
		shiftWindow(curDepth, curWindow)
		if sumWindow(curWindow) > sumWindow(prevWindow) {
			incCount++
		}
	}

	fmt.Printf("There are %d increasing measurements.\n", incCount)
}

func shiftWindow(d int, w []int) {
	w[0] = w[1]
	w[1] = w[2]
	w[2] = d
}

func sumWindow(w []int) int {
	sum := 0
	for _, d := range w {
		sum += d
	}
	return sum
}

func copyWindow(sw, dw []int) {
	for i := range sw {
		dw[i] = sw[i]
	}
}
