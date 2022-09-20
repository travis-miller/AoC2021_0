package main

import "fmt"

func main() {
	numRolls := 0
	position1 := 9
	position2 := 10
	score1 := 0
	score2 := 0
	isPlayerOneTurn := true

	for score1 < 1000 && score2 < 1000 {
		for i := 0; i < 3; i++ {
			numRolls++
			if isPlayerOneTurn {
				position1 += numRolls % 100
				position1 = (position1-1)%10 + 1
			} else {
				position2 += numRolls % 100
				position2 = (position2-1)%10 + 1
			}
		}
		if isPlayerOneTurn {
			score1 += position1
		} else {
			score2 += position2
		}
		isPlayerOneTurn = !isPlayerOneTurn
	}

	output := numRolls * score1
	if score2 < score1 {
		output = numRolls * score2
	}

	fmt.Printf("The output is %d\n", output)
}
