package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inputPos int
	boardPos := -1
	var drawOrder []string
	bingoBoards := []*bingoBoard{}
	for scanner.Scan() {
		if inputPos == 0 {
			inputPos++
			drawOrder = strings.Split(scanner.Text(), ",")
			continue
		}
		if scanner.Text() == "" {
			boardPos++
			bingoBoards = append(bingoBoards, newBoard())
			continue
		}
		bingoBoards[boardPos].addRow(strings.Split(scanner.Text(), " "))
	}

	winDraw, winBoard := findWinner(drawOrder, bingoBoards)
	if winBoard == nil {
		log.Fatalf("Failed to find a winning board.")
	}
	unmarked := winBoard.unmarkedSquares()
	output, err := findOutput(winDraw, unmarked)
	if err != nil {
		log.Fatalf("Failed to find the output: %v", err)
	}
	fmt.Printf("Winning Draw: %s Unmarked: %v Output: %d\n", winDraw, unmarked, output)
}

func findWinner(draws []string, boards []*bingoBoard) (string, *bingoBoard) {
	for _, d := range draws {
		fmt.Printf("Marking %v\n", d)
		for _, b := range boards {
			if b.markNumber(d) {
				if b.isBingo() {
					return d, b
				}
			}
		}
	}
	return "", nil
}

func findOutput(draw string, row []string) (int, error) {
	drawNum, err := strconv.Atoi(draw)
	if err != nil {
		return 0, fmt.Errorf("Unable to convert draw number %s to an integer: %v", draw, err)
	}
	rowSum := 0
	for _, s := range row {
		rowNum, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("Unable to convert row number %s to an integer: %v", s, err)
		}
		rowSum += rowNum
	}
	return drawNum * rowSum, nil
}

type bingoBoard struct {
	numbers map[string]bool
	rows    []bingoRow
	columns []bingoRow
}

type bingoRow []bingoSquare

func (r bingoRow) isBingo() bool {
	for _, s := range r {
		if !s.marked {
			return false
		}
	}
	return true
}

func (r bingoRow) rowNums() []string {
	nums := []string{}
	for _, s := range r {
		nums = append(nums, s.number)
	}
	return nums
}

type bingoSquare struct {
	number string
	marked bool
}

func newBoard() *bingoBoard {
	return &bingoBoard{
		numbers: make(map[string]bool),
		rows:    make([]bingoRow, 0),
		columns: make([]bingoRow, 5),
	}
}

func (b *bingoBoard) addRow(row []string) {
	br := bingoRow{}
	colCount := 0
	for _, n := range row {
		if n == "" {
			continue
		}
		b.numbers[n] = true
		bs := bingoSquare{number: n}
		br = append(br, bs)
		b.columns[colCount] = append(b.columns[colCount], bs)
		colCount++
	}
	b.rows = append(b.rows, br)
}

func (b *bingoBoard) markNumber(num string) bool {
	if !b.numbers[num] {
		return false
	}
	for _, r := range b.rows {
		for i, s := range r {
			if s.number == num {
				r[i] = bingoSquare{marked: true, number: s.number}
			}
		}
	}
	for _, c := range b.columns {
		for i, s := range c {
			if s.number == num {
				c[i] = bingoSquare{marked: true, number: s.number}
			}
		}
	}
	return true
}

func (b *bingoBoard) isBingo() bool {
	for _, r := range b.rows {
		if r.isBingo() {
			return true
		}
	}
	for _, c := range b.columns {
		if c.isBingo() {
			return true
		}
	}
	return false
}

func (b *bingoBoard) unmarkedSquares() []string {
	unmarked := []string{}
	for _, r := range b.rows {
		for _, s := range r {
			if !s.marked {
				unmarked = append(unmarked, s.number)
			}
		}
	}
	return unmarked
}
