package main

import "fmt"

func main() {
	start := gameState{pos1: 9, pos2: 10}
	mem := map[gameState]gameResults{}
	result := gameResults{}
	result = result.combine(runRound(1, start, mem))
	result = result.combine(runRound(2, start, mem))
	result = result.combine(runRound(3, start, mem))
	fmt.Printf("Player 1 wins: %d, Player 2 wins: %d\n", result.player1Wins, result.player2Wins)
}

type gameState struct {
	pos1    int
	pos2    int
	score1  int
	score2  int
	rollPos int
}

type gameResults struct {
	player1Wins int
	player2Wins int
}

func (r gameResults) combine(other gameResults) gameResults {
	return gameResults{r.player1Wins + other.player1Wins, r.player2Wins + other.player2Wins}
}

func runRound(roll int, state gameState, mem map[gameState]gameResults) gameResults {
	if state.rollPos < 3 {
		state.pos1 = (state.pos1+roll-1)%10 + 1
	} else {
		state.pos2 = (state.pos2+roll-1)%10 + 1
	}

	if state.rollPos == 2 {
		state.score1 = state.score1 + state.pos1
	}
	if state.rollPos == 5 {
		state.score2 = state.score2 + state.pos2
	}

	if state.score1 >= 21 {
		return gameResults{1, 0}
	}
	if state.score2 >= 21 {
		return gameResults{0, 1}
	}

	state.rollPos = (state.rollPos + 1) % 6

	if r, ok := mem[state]; ok {
		return r
	}

	result := gameResults{}
	result = result.combine(runRound(1, state, mem))
	result = result.combine(runRound(2, state, mem))
	result = result.combine(runRound(3, state, mem))
	mem[state] = result
	return result
}
