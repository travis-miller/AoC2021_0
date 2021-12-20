package main

import (
	"bufio"
	"container/heap"
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

	riskGrid := [][]*PathNode{}
	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		riskRow := []*PathNode{}
		for x, r := range scanner.Text() {
			num, err := strconv.Atoi(string([]rune{r}))
			if err != nil {
				log.Fatalf("Failed to convert %c to int: %v", r, err)
			}
			riskRow = append(riskRow, &PathNode{Risk: num, Point: Point{x, y}})
		}
		riskGrid = append(riskGrid, riskRow)
		y++
	}
	for y, row := range riskGrid {
		for x, n := range row {
			if y != 0 {
				n.Neighbors = append(n.Neighbors, riskGrid[y-1][x])
			}
			if x != 0 {
				n.Neighbors = append(n.Neighbors, riskGrid[y][x-1])
			}
			if x < len(row)-1 {
				n.Neighbors = append(n.Neighbors, riskGrid[y][x+1])
			}
			if y < len(riskGrid)-1 {
				n.Neighbors = append(n.Neighbors, riskGrid[y+1][x])
			}
		}
	}
	start := riskGrid[0][0]
	maxX, maxY := len(riskGrid[0])-1, len(riskGrid)-1
	end := riskGrid[maxY][maxX]
	path, err := findLowestRisk(start, end)
	if err != nil {
		log.Fatalf("Failed to find lowest risk: %v", err)
	}
	fmt.Printf("The lowest risk is %d\n", pathRisk(path, riskGrid))
}

func findLowestRisk(start, end *PathNode) ([]Point, error) {
	pq := PathQueue{}
	heap.Init(&pq)
	heap.Push(&pq, start)
	cameFrom := map[Point]Point{}
	gScore := map[Point]int{start.Point: 0}
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*PathNode)
		if cur.Point == end.Point {
			return constuctPath(end, cameFrom), nil
		}
		for _, n := range cur.Neighbors {
			gs := gScore[cur.Point] + n.Risk
			if gScore[n.Point] == 0 || gs < gScore[n.Point] {
				cameFrom[n.Point] = cur.Point
				gScore[n.Point] = gs
				newNode := &PathNode{
					Point:     n.Point,
					Risk:      n.Risk,
					PathScore: gs + (end.x - n.x) + (end.y - n.y),
					Neighbors: n.Neighbors,
				}
				heap.Push(&pq, newNode)
			}
		}
	}
	return nil, fmt.Errorf("No path found.")
}

func constuctPath(end *PathNode, cameFrom map[Point]Point) []Point {
	path := []Point{end.Point}
	nextNode := cameFrom[end.Point]
	start := Point{0, 0}
	for nextNode != start {
		path = append(path, nextNode)
		nextNode = cameFrom[nextNode]
	}
	return path
}

func pathRisk(path []Point, riskGrid [][]*PathNode) int {
	totalRisk := 0
	for _, p := range path {
		totalRisk += riskGrid[p.y][p.x].Risk
	}
	return totalRisk
}

type PathNode struct {
	Point
	Risk      int
	PathScore int
	Neighbors []*PathNode
}

type Point struct {
	x, y int
}

type PathQueue []*PathNode

func (pq PathQueue) Len() int {
	return len(pq)
}

func (pq PathQueue) Less(i, j int) bool {
	return pq[i].PathScore < pq[j].PathScore
}

func (pq PathQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PathQueue) Push(x interface{}) {
	node := x.(*PathNode)
	*pq = append(*pq, node)
}

func (pq *PathQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return node
}
