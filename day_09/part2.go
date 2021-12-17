package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	heightmap := [][]location{}
	for scanner.Scan() {
		line := scanner.Text()
		loc := []location{}
		for _, n := range line {
			height, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatalf("Failed to convert %c to an integer: %v", n, err)
			}
			loc = append(loc, location{height, false})
		}
		heightmap = append(heightmap, loc)
	}

	bs := findBasinSizes(heightmap)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] > bs[j]
	})
	if len(bs) < 3 {
		log.Fatalf("Less than 3 basins found.")
	}
	output := bs[0] * bs[1] * bs[2]
	fmt.Printf("Output: %d\n", output)
}

type location struct {
	height   int
	searched bool
}

func findBasinSizes(heightmap [][]location) []int {
	basinSizes := []int{}
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap); j++ {
			if heightmap[i][j].height != 9 && !heightmap[i][j].searched {
				b := searchBasin(heightmap, i, j)
				if len(b) != 0 {
					basinSizes = append(basinSizes, len(b))
				}
			}
		}
	}
	return basinSizes
}

func searchBasin(heightmap [][]location, x, y int) []int {
	basin := []int{}
	loc := heightmap[x][y]
	if loc.height != 9 && !loc.searched {
		heightmap[x][y] = location{loc.height, true}
		if x != 0 {
			basin = append(basin, searchBasin(heightmap, x-1, y)...)
		}
		if x < len(heightmap)-1 {
			basin = append(basin, searchBasin(heightmap, x+1, y)...)
		}
		if y != 0 {
			basin = append(basin, searchBasin(heightmap, x, y-1)...)
		}
		if y < len(heightmap[x])-1 {
			basin = append(basin, searchBasin(heightmap, x, y+1)...)
		}
		basin = append(basin, loc.height)
	}
	return basin
}
