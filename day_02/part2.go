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
	var horPos, verPos, aim int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if len(cmd) < 2 {
			log.Fatalf("Command %v is in an unexpected format", cmd)
		}
		amount, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalf("Unable to convert command amount %s to an integer: %v", cmd[1], err)
		}
		switch cmd[0] {
		case "forward":
			horPos += amount
			verPos += amount * aim
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	fmt.Printf("Horizontal Position: %d Vertical Position: %d Multiplied: %d\n", horPos, verPos, horPos*verPos)
}
