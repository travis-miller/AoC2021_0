package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxSteps = 10

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	parseRules := false
	template := ""
	rules := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseRules = true
			continue
		}
		if parseRules {
			r := strings.Split(line, " -> ")
			if len(r) != 2 {
				log.Fatalf("Line %q has unexpected rule format.", line)
			}
			rules[r[0]] = r[1]
		} else {
			template = line
		}
	}
	for i := 0; i < maxSteps; i++ {
		template = applyInsertion(template, rules)
	}
	elementCount := map[rune]int{}
	for _, r := range template {
		elementCount[r]++
	}
	minCount := -1
	maxCount := -1
	for _, v := range elementCount {
		if minCount == -1 || v < minCount {
			minCount = v
		}
		if maxCount == -1 || v > maxCount {
			maxCount = v
		}
	}
	fmt.Printf("Max: %d Min: %d Diff: %d\n", minCount, maxCount, maxCount-minCount)
}

func applyInsertion(template string, rules map[string]string) string {
	var sb strings.Builder
	var prevRune rune
	for _, curRune := range template {
		if prevRune != 0 {
			pair := string([]rune{prevRune, curRune})
			sb.WriteString(rules[pair])
		}
		sb.WriteRune(curRune)
		prevRune = curRune
	}
	return sb.String()
}
