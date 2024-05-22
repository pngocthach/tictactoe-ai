package game

import "fmt"

// type Moves []Move

// func (m Moves) Len() int {
// 	return len(m)
// }

// func (m Moves) Swap(i, j int) {
// 	m[i], m[j] = m[j], m[i]
// }

// func (m Moves) Less(i, j int) bool {
// 	return EvaluateMove(&m[i]) > EvaluateMove(&m[j])
// }

// func EvaluateMove(move *Move) int {
// 	// dem so luong quan lien tiep xung quanh nuoc di
// 	count := 0
// 	for _, dir := range DIRECTION {
// 		if move.X+dir.X >= 0 && move.X+dir.X < BOARD_SIZE && move.Y+dir.Y >= 0 && move.Y+dir.Y < BOARD_SIZE {
// 			if Board[move.X+dir.X][move.Y+dir.Y] != EMPTY {
// 				count++
// 			}
// 		}
// 	}

// 	return 0
// }

func (t *TicTacToe) EvaluateMove(move *Move) int {
	player := t.GetPlayer()
	opp := t.GetOpponent()
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
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.boardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.boardSize {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] != opp {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] == opp {
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
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.boardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.boardSize {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] != opp {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] == opp {
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
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.boardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.boardSize {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] != player {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] == player {
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
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.boardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.boardSize {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] != player {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] == player {
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
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.boardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.boardSize {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] != opp {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] == opp {
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
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.boardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.boardSize {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] != opp {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] == opp {
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
			if left.Row+LEFT[i].Row < 1 || left.Row+LEFT[i].Row > t.boardSize || left.Col+LEFT[i].Col < 1 || left.Col+LEFT[i].Col > t.boardSize {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] != player {
				break
			}
			if t.Board[left.Row+LEFT[i].Row][left.Col+LEFT[i].Col] == player {
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
			if right.Row+RIGHT[i].Row < 1 || right.Row+RIGHT[i].Row > t.boardSize || right.Col+RIGHT[i].Col < 1 || right.Col+RIGHT[i].Col > t.boardSize {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] != player {
				break
			}
			if t.Board[right.Row+RIGHT[i].Row][right.Col+RIGHT[i].Col] == player {
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

func (t *TicTacToe) PrintNeighbors() {
	moves := t.GetNeighbor(1)
	fmt.Println("Neighbors:", len(moves), moves)
}
