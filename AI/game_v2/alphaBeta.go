package game2

import (
	"fmt"
	"math"

	"github.com/huandu/go-clone"
)

func Minimax(t *TicTacToe, isMaximize bool, depth int, alpha float64, beta float64) float64 {
	if depth == MAX_DEPTH || t.CheckWin() || t.MoveCount == t.boardSize*t.boardSize {
		return t.Evaluate()
	}

	if isMaximize {
		bestScore := math.Inf(-1)
		for _, move := range t.GetEmptyCells() {
			oldOstate := t.OState
			oldXstate := t.XState
			t.MakeMove(move)
			t.OState = oldOstate
			t.XState = oldXstate
			t.MoveCount--
			t.Board[move.Row][move.Col].Value = EMPTY
			score := Minimax(t, false, depth+1, alpha, beta)
			bestScore = max(bestScore, score)
			alpha = max(alpha, bestScore)
			if beta <= alpha {
				break
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for _, move := range t.GetEmptyCells() {
			oldOstate := t.OState
			oldXstate := t.XState
			t.MakeMove(move)
			t.OState = oldOstate
			t.XState = oldXstate
			t.Board[move.Row][move.Col].Value = EMPTY
			t.MoveCount--
			score := Minimax(t, true, depth+1, alpha, beta)
			bestScore = min(bestScore, score)
			beta = min(beta, bestScore)
			if beta <= alpha {
				break
			}
		}
		return bestScore
	}
}

func (t *TicTacToe) Evaluate() float64 {
	return float64(len(t.XState.Five)*1000+len(t.XState.OpenFour)*500+len(t.XState.BrokenCloseFour)*100+len(t.XState.CloseFour)*50+len(t.XState.CloseThree)*10+len(t.XState.OpenTwo)) - 1.2*float64(len(t.OState.Five)*1000-len(t.OState.OpenFour)*500-len(t.OState.BrokenCloseFour)*100-len(t.OState.CloseFour)*50-len(t.OState.CloseThree)*10-len(t.OState.OpenTwo))
}

func (t *TicTacToe) GetBestMove() Move {
	bestScore := math.Inf(1)
	var bestMove Move
	for _, move := range t.GetEmptyCells() {
		oldT := clone.Clone(t).(*TicTacToe)
		t.MakeMove(move)
		score := Minimax(t, false, 0, math.MinInt32, math.MaxInt32)
		t = oldT
		t.MoveCount--
		t.Board[move.Row][move.Col].Value = EMPTY
		// score := 1
		// t.Move(move, EMPTY)
		if score < bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func (t *TicTacToe) PlayPvAI(AIplayer int) {
	player := PLAYER_X
	for !t.CheckWin() {
		t.PrintBoard()
		fmt.Println("Player", player)
		if player == AIplayer {
			move := t.GreedyMove()
			t.MakeMove(move)
		} else {
			move := getUserInput()
			err := t.MakeMove(move)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		player *= -1
	}

	t.PrintBoard()
	if player == PLAYER_X {
		fmt.Println("Player O wins")
	} else {
		fmt.Println("Player X wins")
	}
}

func (t *TicTacToe) GreedyMove() Move {
	bestScore := math.Inf(1)
	var bestMove Move
	for _, move := range t.GetEmptyCells() {
		oldT := t.CloneTicTacToe()
		t.MakeMove(move)
		score := t.Evaluate()
		t = oldT
		if score < bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}

func (t *TicTacToe) CloneTicTacToe() *TicTacToe {
	newBoard := make([][]Node, t.boardSize)
	for i := 0; i < t.boardSize; i++ {
		newBoard[i] = make([]Node, t.boardSize)
		for j := 0; j < t.boardSize; j++ {
			newBoard[i][j] = t.Board[i][j]
		}
	}

	newXstate := clone.Clone(t.XState).(State)
	newOstate := clone.Clone(t.OState).(State)

	return &TicTacToe{
		Board:     newBoard,
		boardSize: t.boardSize,
		MoveCount: t.MoveCount,
		winSize:   t.winSize,
		XState:    newXstate,
		OState:    newOstate,
	}
}
