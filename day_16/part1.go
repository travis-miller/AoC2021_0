package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	var sb strings.Builder
	for _, c := range string(input) {
		num, err := strconv.ParseInt(string([]rune{c}), 16, 64)
		if err != nil {
			log.Fatalf("Failed to convert %c into integer. Not a hex character.", c)
		}
		sb.WriteString(fmt.Sprintf("%04b", num))
	}
	encoded := sb.String()
	ptr := 0
	p := parsePacket(strings.Split(encoded, ""), &ptr)
	fmt.Printf("Sum of the versions: %d\n", p.sumVersions())
}

func parsePacket(encoded []string, ptr *int) packet {
	p := packet{}
	version := parseVersion(encoded, ptr)
	typeId := parseTypeId(encoded, ptr)
	switch typeId {
	case 4:
		p = parseLiteralValuePacket(encoded, ptr)
	default:
		p = parseOperatorPacket(encoded, ptr)
	}
	p.version = version
	p.typeId = typeId
	return p
}

func parseVersion(encoded []string, ptr *int) int {
	version, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+3], ""), 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse version: %v", err)
	}
	*ptr += 3
	return int(version)
}

func parseTypeId(encoded []string, ptr *int) int {
	typeId, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+3], ""), 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse type ID: %v", err)
	}
	*ptr += 3
	return int(typeId)
}

func parseLiteralValuePacket(encoded []string, ptr *int) packet {
	end := false
	var sb strings.Builder
	for !end {
		if encoded[*ptr] == "0" {
			end = true
		}
		sb.WriteString(strings.Join(encoded[*ptr+1:*ptr+5], ""))
		*ptr += 5
	}
	value, err := strconv.ParseInt(sb.String(), 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse value string %s: %v", sb.String(), err)
	}
	return packet{value: int(value)}
}

func parseOperatorPacket(encoded []string, ptr *int) packet {
	lenId, lenSize := parseSubLength(encoded, ptr)
	sp := []packet{}
	if lenId == 0 {
		subEnd := *ptr + lenSize
		for *ptr < subEnd {
			sp = append(sp, parsePacket(encoded, ptr))
		}
	} else {
		for i := 0; i < lenSize; i++ {
			sp = append(sp, parsePacket(encoded, ptr))
		}
	}
	return packet{subPackets: sp}
}

func parseSubLength(encoded []string, ptr *int) (int, int) {
	if encoded[*ptr] == "0" {
		*ptr++
		length, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+15], ""), 2, 64)
		if err != nil {
			log.Fatalf("Failed to parse type ID: %v", err)
		}
		*ptr += 15
		return 0, int(length)
	} else {
		*ptr++
		length, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+11], ""), 2, 64)
		if err != nil {
			log.Fatalf("Failed to parse type ID: %v", err)
		}
		*ptr += 11
		return 1, int(length)
	}
}

type packet struct {
	version    int
	typeId     int
	value      int
	subPackets []packet
}

func (p packet) sumVersions() int {
	total := p.version
	for _, sp := range p.subPackets {
		total += sp.sumVersions()
	}
	return total
}
