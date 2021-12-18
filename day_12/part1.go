package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	caves := map[string]*cave{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), "-")
		if len(edge) != 2 {
			log.Fatalf("Edge %s in unexpected format.", scanner.Text())
		}
		c, ok := caves[edge[0]]
		if !ok {
			isLarge := false
			if edge[0] == strings.ToUpper(edge[0]) {
				isLarge = true
			}
			c = &cave{edge[0], isLarge, []string{}}
			caves[c.name] = c
		}
		c.edges = append(c.edges, edge[1])

		c, ok = caves[edge[1]]
		if !ok {
			isLarge := false
			if edge[1] == strings.ToUpper(edge[1]) {
				isLarge = true
			}
			c = &cave{edge[1], isLarge, []string{}}
			caves[c.name] = c
		}
		c.edges = append(c.edges, edge[0])
	}
	paths := findPaths("start", caves, []string{})
	fmt.Printf("There are %d paths.\n", len(paths))
}

type cave struct {
	name    string
	isLarge bool
	edges   []string
}

func findPaths(curCave string, caves map[string]*cave, path []string) [][]string {
	if curCave == "end" {
		path = append(path, curCave)
		return [][]string{path}
	}
	if !caves[curCave].isLarge {
		for _, c := range path {
			if curCave == c {
				return [][]string{}
			}
		}
	}
	paths := [][]string{}
	path = append(path, curCave)
	for _, c := range caves[curCave].edges {
		// Copy slice to prevent corruption
		newPath := append([]string{}, path...)
		paths = append(paths, findPaths(c, caves, newPath)...)
	}
	return paths
}
