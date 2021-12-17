package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()
	ds := []display{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		if len(line) != 2 {
			log.Fatalf("Unexpeceted input for line %s", scanner.Text())
		}
		signal := map[string]string{}
		for _, sig := range strings.Split(line[0], " ") {
			segs := strings.Split(sig, "")
			sort.Slice(segs, func(i, j int) bool {
				return segs[i] < segs[j]
			})
			signal[strings.Join(segs, "")] = "unknown"
		}
		output := []string{}
		for _, out := range strings.Split(line[1], " ") {
			segs := strings.Split(out, "")
			sort.Slice(segs, func(i, j int) bool {
				return segs[i] < segs[j]
			})
			output = append(output, strings.Join(segs, ""))
		}
		ds = append(ds, display{signal: signal, output: output})
	}

	sumOutputs := 0
	for _, d := range ds {
		sumOutputs += d.decodeOutput()
	}
	fmt.Printf("Sum of the outputs %d\n", sumOutputs)
}

type display struct {
	signal map[string]string
	output []string
}

func (d display) decodeOutput() int {
	d.findDigitByLength(2, "1")
	d.findDigitByLength(3, "7")
	d.findDigitByLength(4, "4")
	d.findDigitByLength(7, "8")
	d.findDigitByLengthAndMatch(6, "4", "9")
	d.findDigitByLengthAndMatch(5, "7", "3")
	d.findZero()
	d.markMissing(6, "6")
	d.findFive()
	d.markMissing(5, "2")

	var sb = strings.Builder{}
	for _, o := range d.output {
		sb.WriteString(d.signal[o])
	}
	total, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatalf("Failed to parse %s as integer.", sb.String())
	}
	return total
}

func (d display) findDigitByLength(segLength int, digit string) {
	for sig := range d.signal {
		if len(sig) == segLength {
			d.signal[sig] = digit
			return
		}
	}
}

func (d display) findDigitByLengthAndMatch(segLength int, match, digit string) {
	matchSig := ""
	for sig, d := range d.signal {
		if d == match {
			matchSig = sig
			break
		}
	}

	for sig, dig := range d.signal {
		if dig != "unknown" || len(sig) != segLength {
			continue
		}
		if !containsSegments(sig, matchSig) {
			continue
		}
		d.signal[sig] = digit
		break
	}
}

func (d display) findFive() {
	sixSig := d.findDigit("6")

	for sig, dig := range d.signal {
		if dig != "unknown" || len(sig) != 5 {
			continue
		}
		if !containsSegments(sixSig, sig) {
			continue
		}
		d.signal[sig] = "5"
		break
	}
}

func (d display) findZero() {
	eightSig := d.findDigit("8")
	sevenSig := d.findDigit("7")

	for sig, dig := range d.signal {
		if dig != "unknown" || len(sig) != 6 {
			continue
		}
		if !containsSegments(sig, eightSig) && !containsSegments(sig, sevenSig) {
			continue
		}
		d.signal[sig] = "0"
		break
	}
}

func (d display) findDigit(digit string) string {
	for sig, d := range d.signal {
		if d == digit {
			return sig
		}
	}
	return ""
}

func containsSegments(match, segs string) bool {
	for _, seg := range segs {
		if !strings.ContainsRune(match, seg) {
			return false
		}
	}
	return true
}

func (d display) markMissing(length int, digit string) {
	for sig, dig := range d.signal {
		if len(sig) == length && dig == "unknown" {
			d.signal[sig] = digit
		}
	}
}
