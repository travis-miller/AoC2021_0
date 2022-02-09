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
	var sNums []*snailNumber
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sn := parseSnailNumber(scanner.Text())
		sNums = append(sNums, sn)
	}
	var total *snailNumber
	for _, n := range sNums {
		total = total.sum(n)
		for total.explode() || total.split() {
			// Reduce until finished
		}
	}
	fmt.Println(total)
	fmt.Printf("Magnitude: %d\n", total.magnitude())
}

func parseSnailNumber(str string) *snailNumber {
	var prevNum *snailNumber
	sNum := &snailNumber{}
	curNum := sNum
	var rightNum bool
	firstNum := true
	for _, c := range str {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num, _ := strconv.Atoi(string(c))
			if rightNum {
				curNum.rightValue = num
			} else {
				curNum.leftValue = num
			}
		case '[':
			if !firstNum {
				prevNum = curNum
				curNum = &snailNumber{parent: prevNum}
				if rightNum {
					prevNum.childRight = curNum
					rightNum = false
				} else {
					prevNum.childLeft = curNum
				}
			}
			firstNum = false
		case ']':
			curNum = curNum.parent
		case ',':
			rightNum = true
		default:
			panic(fmt.Sprintf("unknown snailfish number rune: %c", c))
		}
	}
	return sNum
}

type snailNumber struct {
	leftValue  int
	childLeft  *snailNumber
	rightValue int
	childRight *snailNumber
	parent     *snailNumber
}

func (n *snailNumber) sum(other *snailNumber) *snailNumber {
	if n == nil {
		return other
	}
	newNum := &snailNumber{childLeft: n, childRight: other}
	n.parent = newNum
	other.parent = newNum
	return newNum
}

func (n *snailNumber) depth() int {
	depth := 1
	if n.parent != nil {
		depth = 1 + n.parent.depth()
	}
	return depth
}

func (n *snailNumber) explode() bool {
	if n.childLeft != nil && n.childLeft.depth() > 4 {
		n.explodeLeft()
		return true
	}
	if n.childLeft != nil && n.childLeft.explode() {
		return true
	}
	if n.childRight != nil && n.childRight.depth() > 4 {
		n.explodeRight()
		return true
	}
	if n.childRight != nil && n.childRight.explode() {
		return true
	}
	return false
}

func (n *snailNumber) split() bool {
	if n.leftValue > 9 {
		n.splitLeft()
		return true
	}
	if n.childLeft != nil && n.childLeft.split() {
		return true
	}
	if n.rightValue > 9 {
		n.splitRight()
		return true
	}
	if n.childRight != nil && n.childRight.split() {
		return true
	}
	return false
}

func (n *snailNumber) explodeLeft() {
	if n.childRight == nil {
		n.rightValue = n.rightValue + n.childLeft.rightValue
	} else {
		n.addToNextRight(n.childLeft, n.childLeft.rightValue)
	}

	if n.parent != nil {
		n.addToNextLeft(n.childLeft, n.childLeft.leftValue)
	}

	n.leftValue = 0
	n.childLeft = nil
}

func (n *snailNumber) explodeRight() {
	if n.childLeft == nil {
		n.leftValue = n.leftValue + n.childRight.leftValue
	} else {
		n.addToNextLeft(n.childRight, n.childRight.leftValue)
	}

	if n.parent != nil {
		n.addToNextRight(n.childRight, n.childRight.rightValue)
	}

	n.rightValue = 0
	n.childRight = nil
}

func (n *snailNumber) addToNextLeft(prev *snailNumber, num int) {
	if n.childLeft == prev {
		if n.parent != nil {
			n.parent.addToNextLeft(n, num)
		}
		return
	}
	if n.childLeft == nil {
		n.leftValue += num
		return
	}
	n.childLeft.addToRight(num)
}

func (n *snailNumber) addToNextRight(prev *snailNumber, num int) {
	if n.childRight == prev {
		if n.parent != nil {
			n.parent.addToNextRight(n, num)
		}
		return
	}
	if n.childRight == nil {
		n.rightValue += num
		return
	}
	n.childRight.addToLeft(num)
}

func (n *snailNumber) addToRight(num int) {
	if n.childRight != nil {
		n.childRight.addToRight(num)
		return
	}
	n.rightValue += num
}

func (n *snailNumber) addToLeft(num int) {
	if n.childLeft != nil {
		n.childLeft.addToLeft(num)
		return
	}
	n.leftValue += num
}

func (n *snailNumber) splitLeft() {
	n.childLeft = &snailNumber{
		leftValue:  n.leftValue / 2,
		rightValue: (n.leftValue + 1) / 2,
		parent:     n,
	}
	n.leftValue = 0
}

func (n *snailNumber) splitRight() {
	n.childRight = &snailNumber{
		leftValue:  n.rightValue / 2,
		rightValue: (n.rightValue + 1) / 2,
		parent:     n,
	}
	n.rightValue = 0
}

func (n *snailNumber) magnitude() int {
	total := 0
	if n.childLeft != nil {
		total += n.childLeft.magnitude() * 3
	} else {
		total += n.leftValue * 3
	}

	if n.childRight != nil {
		total += n.childRight.magnitude() * 2
	} else {
		total += n.rightValue * 2
	}
	return total
}

func (n *snailNumber) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	if n.childLeft == nil {
		sb.WriteString(strconv.Itoa(n.leftValue))
	} else {
		sb.WriteString(n.childLeft.String())
	}
	sb.WriteRune(',')
	if n.childRight == nil {
		sb.WriteString(strconv.Itoa(n.rightValue))
	} else {
		sb.WriteString(n.childRight.String())
	}
	sb.WriteRune(']')
	return sb.String()
}
