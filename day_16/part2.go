package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type operation int

const (
	sum operation = iota
	product
	min
	max
	value
	greaterThan
	lessThan
	equal
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
			log.Fatalf("Failed to convert %c into integer: %v", c, err)
		}
		sb.WriteString(fmt.Sprintf("%04b", num))
	}
	encoded := sb.String()
	ptr := 0
	p := parsePacket(strings.Split(encoded, ""), &ptr)
	fmt.Printf("Packet value: %d\n", p.value())
}

func parsePacket(encoded []string, ptr *int) packet {
	p := packet{}
	version := parseVersion(encoded, ptr)
	typeId := parseTypeId(encoded, ptr)
	switch typeId {
	case value:
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

func parseTypeId(encoded []string, ptr *int) operation {
	typeId, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+3], ""), 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse type ID: %v", err)
	}
	*ptr += 3
	return operation(int(typeId))
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
	return packet{literalValue: int(value)}
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
			log.Fatalf("Failed to parse length: %v", err)
		}
		*ptr += 15
		return 0, int(length)
	} else {
		*ptr++
		length, err := strconv.ParseInt(strings.Join(encoded[*ptr:*ptr+11], ""), 2, 64)
		if err != nil {
			log.Fatalf("Failed to parse length: %v", err)
		}
		*ptr += 11
		return 1, int(length)
	}
}

type packet struct {
	version      int
	typeId       operation
	literalValue int
	subPackets   []packet
}

func (p packet) value() int {
	switch p.typeId {
	case sum:
		v := 0
		for _, sp := range p.subPackets {
			v += sp.value()
		}
		return v
	case product:
		v := 1
		for _, sp := range p.subPackets {
			v *= sp.value()
		}
		return v
	case min:
		v := -1
		for _, sp := range p.subPackets {
			subV := sp.value()
			if v == -1 || subV < v {
				v = subV
			}
		}
		return v
	case max:
		v := -1
		for _, sp := range p.subPackets {
			subV := sp.value()
			if v == -1 || subV > v {
				v = subV
			}
		}
		return v
	case value:
		return p.literalValue
	case greaterThan:
		if p.subPackets[0].value() > p.subPackets[1].value() {
			return 1
		}
		return 0
	case lessThan:
		if p.subPackets[0].value() < p.subPackets[1].value() {
			return 1
		}
		return 0
	case equal:
		if p.subPackets[0].value() == p.subPackets[1].value() {
			return 1
		}
		return 0
	default:
		log.Fatalf("Packet type ID is %d is not recognized.", p.typeId)
		return 0
	}
}
