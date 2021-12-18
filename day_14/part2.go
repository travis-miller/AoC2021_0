package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxSteps = 40

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	parseRules := false
	template := ""
	rules := map[string]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseRules = true
			continue
		}
		if parseRules {
			rule := strings.Split(line, " -> ")
			if len(rule) != 2 {
				log.Fatalf("Line %q has unexpected rule format.", line)
			}
			for _, r := range rule[1] {
				rules[rule[0]] = r
			}
		} else {
			template = line
		}
	}

	mem := map[stepMemKey]map[rune]int{}
	elementCount := map[rune]int{}
	var prevRune rune
	for _, curRune := range template {
		if prevRune != 0 {
			pair := []rune{prevRune, curRune}
			for k, v := range applyInsertion(pair, rules, 0, mem) {
				elementCount[k] += v
			}
		} else {
			elementCount[curRune]++
		}
		prevRune = curRune
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

type stepMemKey struct {
	step int
	pair string
}

func applyInsertion(pair []rune, rules map[string]rune, step int, mem map[stepMemKey]map[rune]int) map[rune]int {
	smk := stepMemKey{step, string(pair)}
	if m, ok := mem[smk]; ok {
		return m
	}
	if step == maxSteps {
		return map[rune]int{pair[1]: 1}
	}
	elementCount := map[rune]int{}
	newElement := rules[string(pair)]
	for k, v := range applyInsertion([]rune{pair[0], newElement}, rules, step+1, mem) {
		elementCount[k] += v
	}
	for k, v := range applyInsertion([]rune{newElement, pair[1]}, rules, step+1, mem) {
		elementCount[k] += v
	}
	mem[smk] = elementCount
	return elementCount
}
