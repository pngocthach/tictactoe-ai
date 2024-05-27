package game3

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
	"time"

	"github.com/huandu/go-clone"
)

func (t *TicTacToe) AlphaBeta(depth int, alpha float64, beta float64) float64 {
	// return t.Evaluate() after 10 seconds
	if time.Since(AlphaBetaStart) > MAX_TIME*time.Second && depth > 1 {
		return t.Evaluate()
	}

	if depth == MAX_DEPTH || t.CheckWin() || t.MoveCount == t.BoardSize*t.BoardSize {
		return t.Evaluate()
	}

	isMaximize := t.GetPlayer() == PLAYER_X
	var bestScore float64
	moves := t.GetNeighbor(DIST)

	isPrune := false
	if !isPrune {
		patternKey := []string{"X open four", "O open four"}
		for _, key := range patternKey {
			if t.CheckPatterns(PATTERN[key]) {
				moves = moves[:5]
				break
			}
		}
	}

	if !isPrune {
		patternKey := []string{"X broken four", "O broken four", "X close four", "O close four"}

		for _, key := range patternKey {
			if t.CheckPatterns(PATTERN[key]) {
				moves = moves[:1]
				break
			}
		}
	}

	if isMaximize {
		bestScore = math.Inf(-1)
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i])
			score := t.AlphaBeta(depth+1, alpha, beta)
			t.UndoMove(moves[i])
			t.Dist = oldDist
			bestScore = max(bestScore, score)
			alpha = max(alpha, bestScore)
			if beta <= alpha {
				break
			}
		}
		return bestScore
	} else {
		bestScore = math.Inf(1)
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i])
			score := t.AlphaBeta(depth+1, alpha, beta)
			// println("player", t.GetPlayer(), "move to", move.Row, move.Col, "eval: ", int(t.Evaluate()*10), " - ", score, "depth:", depth)
			t.UndoMove(moves[i])
			t.Dist = oldDist
			bestScore = min(bestScore, score)
			beta = min(beta, bestScore)
			if beta <= alpha {
				break
			}
		}
		return bestScore
	}
}

func (t *TicTacToe) AlphaBetaParallel(depth int, alpha float64, beta float64) float64 {
	if (time.Since(AlphaBetaStart) > MAX_TIME*time.Second && depth > 1) ||
		depth == MAX_DEPTH ||
		t.CheckWin() ||
		t.MoveCount == t.BoardSize*t.BoardSize {

		eval := t.Evaluate()
		return eval
	}

	isMaximize := t.GetPlayer() == PLAYER_X
	var bestScore float64
	moves := t.GetNeighbor(DIST)

	isPrune := false
	if len(moves) == 0 {
		moves = []Move{{t.BoardSize / 2, t.BoardSize / 2}}
		isPrune = true
	}

	if !isPrune {
		patternKey := []string{"X open four", "O open four"}
		for _, key := range patternKey {
			if t.CheckPatterns(PATTERN[key]) {
				moves = moves[:5]
				break
			}
		}
	}

	if !isPrune {
		patternKey := []string{"X broken four", "O broken four", "X close four", "O close four"}

		for _, key := range patternKey {
			if t.CheckPatterns(PATTERN[key]) {
				moves = moves[:1]
				break
			}
		}
	}

	if isMaximize {
		bestScore = math.Inf(-1)
		resultChan := make(chan float64)
		branchCount := 0

		for i := 0; i < len(moves); i++ {
			tCopy := t.deepClone()
			go AlphaBetaBranch(&tCopy, depth+1, alpha, beta, resultChan)
			branchCount++
		}

		for i := 0; i < branchCount; i++ {
			res := <-resultChan
			if res > bestScore {
				bestScore = res
			}
		}

		return bestScore
	} else {
		bestScore = math.Inf(1)
		resultChan := make(chan float64)
		branchCount := 0

		for i := 0; i < len(moves); i++ {
			tCopy := t.deepClone()
			go AlphaBetaBranch(&tCopy, depth+1, alpha, beta, resultChan)
			branchCount++
		}

		for i := 0; i < branchCount; i++ {
			res := <-resultChan
			if res < bestScore {
				bestScore = res
			}
		}

		return bestScore
	}
}

func AlphaBetaBranch(t *TicTacToe, depth int, alpha float64, beta float64, result chan float64) {
	if (time.Since(AlphaBetaStart) > MAX_TIME*time.Second && depth > 1) ||
		depth == MAX_DEPTH ||
		t.CheckWin() ||
		t.MoveCount == t.BoardSize*t.BoardSize {

		eval := t.Evaluate()
		result <- eval
	}

	isMaximize := t.GetPlayer() == PLAYER_X
	var bestScore float64
	moves := t.GetNeighbor(DIST)
	if len(moves) == 0 {
		moves = []Move{{t.BoardSize / 2, t.BoardSize / 2}}
	}

	if isMaximize {
		bestScore = math.Inf(-1)
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i])
			score := t.AlphaBeta(depth+1, alpha, beta)
			t.UndoMove(moves[i])
			t.Dist = oldDist
			bestScore = max(bestScore, score)
			alpha = max(alpha, bestScore)
			if beta <= alpha {
				break
			}
		}
		result <- bestScore
	} else {
		bestScore = math.Inf(1)
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i])
			score := t.AlphaBeta(depth+1, alpha, beta)
			// println("player", t.GetPlayer(), "move to", move.Row, move.Col, "eval: ", int(t.Evaluate()*10), " - ", score, "depth:", depth)
			t.UndoMove(moves[i])
			t.Dist = oldDist
			bestScore = min(bestScore, score)
			beta = min(beta, bestScore)
			if beta <= alpha {
				break
			}
		}
		result <- bestScore
	}

}

func (t *TicTacToe) deepClone() TicTacToe {
	return TicTacToe{
		BoardSize: t.BoardSize,
		BoardRow:  clone.Clone(t.BoardRow).([]int),
		BoardCol:  clone.Clone(t.BoardCol).([]int),
		Dist:      clone.Clone(t.Dist).([][]int),
		MoveCount: t.MoveCount,
	}
}

var AlphaBetaStart = time.Now()

func (t *TicTacToe) GetBestMove() Move {
	AlphaBetaStart = time.Now()
	player := t.GetPlayer()
	if player == PLAYER_X {
		bestScore := math.Inf(-1)
		moves := t.GetNeighbor(DIST)
		var bestMove Move
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i])
			// score := t.AlphaBeta(0, math.Inf(-1), math.Inf(1))
			score := t.AlphaBetaParallel(0, math.Inf(-1), math.Inf(1))
			t.UndoMove(moves[i])
			t.Dist = oldDist
			if score > bestScore {
				bestScore = score
				bestMove = moves[i]
			}
		}
		fmt.Println("time: ", time.Since(AlphaBetaStart).Seconds())
		return bestMove
	} else {
		bestScore := math.Inf(1)
		moves := t.GetNeighbor(DIST)
		bestMove := moves[rand.IntN(len(moves))]
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i])
			// score := t.AlphaBeta(0, math.Inf(-1), math.Inf(1))
			score := t.AlphaBetaParallel(0, math.Inf(-1), math.Inf(1))
			t.UndoMove(moves[i])
			t.Dist = oldDist
			if score < bestScore {
				bestScore = score
				bestMove = moves[i]
			}
		}
		fmt.Println("time: ", time.Since(AlphaBetaStart).Seconds())
		return bestMove
	}
}

func (t *TicTacToe) GetNeighbor(dist int) []Move {
	var neighbors []Move
	for k := 1; k <= dist; k++ {
		for i := 1; i < t.BoardSize+1; i++ {
			for j := 1; j < t.BoardSize+1; j++ {
				if t.Dist[i][j] == k && t.GetValue(i, j) == EMPTY {
					neighbors = append(neighbors, Move{i, j})
				}
			}
		}
	}

	// sort moves by score
	sort.SliceStable(neighbors, func(i, j int) bool {
		a := t.EvaluateMove(&neighbors[i])
		b := t.EvaluateMove(&neighbors[j])
		// fmt.Println("Eval: ", neighbors[i], a, " - ", neighbors[j], b)
		return a > b
	})
	return neighbors
}

func (t *TicTacToe) EvaluateMove(move *Move) int {
	player := t.GetPlayer()
	opp := t.GetOpponent(player)
	// minMove := 1
	// maxMove := t.boardSize

	// 1. check xem co 5 con lien tiep khong
	// if t.CheckWin(player) {
	// 	return 5
	// }

	// 2. check xem co 4 con lien tiep khong
	for i := range LEFT {
		// 2.1. Tim o xa nhat ben trai va ben phai trong khoang 4 o
		left := *move
		leftCount := 0
		for j := 1; j <= 4; j++ {
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.BoardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) != opp {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) == opp {
				left.Row += LEFT[i].Row
				left.Col += LEFT[i].Col
				leftCount++
			}
		}
		if leftCount == 4 {
			return 4
		}
		right := *move
		rightCount := 0
		for j := 1; j <= 4; j++ {
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.BoardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) != opp {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) == opp {
				right.Row += RIGHT[i].Row
				right.Col += RIGHT[i].Col
				rightCount++
			}
		}
		if rightCount+leftCount == 4 {
			return 4
		}
	}

	// 2.3. Neu la 4 con lien tiep cua minh thi tra ve 3
	for i := range LEFT {
		left := *move
		leftCount := 0
		for j := 1; j <= 3; j++ {
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.BoardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) != player {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) == player {
				left.Row += LEFT[i].Row
				left.Col += LEFT[i].Col
				leftCount++
			}
		}
		if leftCount == 3 {
			return 3
		}
		right := *move
		rightCount := 0
		for j := 1; j <= 3; j++ {
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.BoardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) != player {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) == player {
				right.Row += RIGHT[i].Row
				right.Col += RIGHT[i].Col
				rightCount++
			}
		}
		if rightCount+leftCount == 3 {
			return 3
		}
	}
	// 3.1 Neu la 3 con lien tiep cua doi thu thi tra ve 2
	for i := range LEFT {
		left := *move
		leftCount := 0
		for j := 1; j <= 3; j++ {
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.BoardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) != opp {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) == opp {
				left.Row += LEFT[i].Row
				left.Col += LEFT[i].Col
				leftCount++
			}
		}
		if leftCount == 3 {
			return 2
		}
		right := *move
		rightCount := 0
		for j := 1; j <= 3; j++ {
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.BoardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) != opp {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) == opp {
				right.Row += RIGHT[i].Row
				right.Col += RIGHT[i].Col
				rightCount++
			}
		}
		if rightCount+leftCount == 3 {
			return 2
		}
	}
	// 3.2. Neu la 3 con lien tiep cua minh thi tra ve 1
	for i := range LEFT {
		left := *move
		leftCount := 0
		for j := 1; j <= 2; j++ {
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.BoardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) != player {
				break
			}
			if t.GetValue(left.Row+LEFT[i].Row, left.Col+LEFT[i].Col) == player {
				left.Row += LEFT[i].Row
				left.Col += LEFT[i].Col
				leftCount++
			}
		}
		if leftCount == 2 {
			return 1
		}
		right := *move
		rightCount := 0
		for j := 1; j <= 2; j++ {
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.BoardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.BoardSize {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) != player {
				break
			}
			if t.GetValue(right.Row+RIGHT[i].Row, right.Col+RIGHT[i].Col) == player {
				right.Row += RIGHT[i].Row
				right.Col += RIGHT[i].Col
				rightCount++
			}
		}
		if rightCount+leftCount == 2 {
			return 1
		}
	}
	// 2.6. Con lai tra ve 0
	return 0
}
