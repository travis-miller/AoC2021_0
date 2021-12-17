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
	report := []string{}
	for scanner.Scan() {
		report = append(report, scanner.Text())
	}
	oRating := findORating(0, report)
	co2Rating := findCO2Rating(0, report)
	fmt.Printf("Oxygen: %d CO2: %d Multiplied: %d\n", oRating, co2Rating, oRating*co2Rating)
}

func findORating(pos int, ratings []string) int {
	if len(ratings) == 1 {
		return lineToDecimal(ratings[0])
	}
	var zeroes, ones []string
	for _, r := range ratings {
		if r[pos] == '0' {
			zeroes = append(zeroes, r)
		}
		if r[pos] == '1' {
			ones = append(ones, r)
		}
	}
	if len(zeroes) > len(ones) {
		return findORating(pos+1, zeroes)
	} else {
		return findORating(pos+1, ones)
	}
}

func findCO2Rating(pos int, ratings []string) int {
	if len(ratings) == 1 {
		return lineToDecimal(ratings[0])
	}
	var zeroes, ones []string
	for _, r := range ratings {
		if r[pos] == '0' {
			zeroes = append(zeroes, r)
		}
		if r[pos] == '1' {
			ones = append(ones, r)
		}
	}
	if len(ones) < len(zeroes) {
		return findCO2Rating(pos+1, ones)
	} else {
		return findCO2Rating(pos+1, zeroes)
	}
}

func lineToDecimal(line string) int {
	num, err := strconv.ParseInt(line, 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse int: %v", err)
	}
	return int(num)
}
