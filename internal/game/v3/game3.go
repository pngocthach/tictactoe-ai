package v3

import (
	"errors"
	"math"
)

type TicTacToe struct {
	BoardRow  []int // bitmask of the board
	BoardCol  []int // bitmask of the board
	Dist      [][]int
	BoardSize int
	MoveCount int
}

type Move struct {
	Row int
	Col int
}

func NewTicTacToe(boardSize int) *TicTacToe {
	t := TicTacToe{}
	t.BoardSize = boardSize
	t.BoardRow = make([]int, boardSize+2)
	t.BoardCol = make([]int, boardSize+2)

	// init wall
	for i := 0; i < boardSize+2; i++ {
		t.SetValueToEmpty(0, i, WALL)
		t.SetValueToEmpty(boardSize+1, i, WALL)
		t.SetValueToEmpty(i, 0, WALL)
		t.SetValueToEmpty(i, boardSize+1, WALL)
	}

	// init dist intmax
	dist := make([][]int, boardSize+2)
	for i := range dist {
		dist[i] = make([]int, boardSize+2)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	InitPattern()

	t.Dist = dist
	return &t
}

func (t *TicTacToe) GetValue(row, col int) int {
	return (t.BoardRow[row] >> ((t.BoardSize + 2 - col) * 2)) & 0b11
}

func (t *TicTacToe) GetPlayer() int {
	if t.MoveCount%2 == 0 {
		return PLAYER_X
	}
	return PLAYER_O
}

func GetValueFromRow(row, col int) int {
	return (row >> (col * 2)) & 0b11
}

func (t *TicTacToe) SetMove(row, col int) error {
	if row < 1 || col < 1 || row > t.BoardSize || col > t.BoardSize {
		return errors.New("out of bound")
	}
	if t.GetValue(row, col) != EMPTY {
		return errors.New("invalid move (square not empty)")
	}

	player := t.GetPlayer()
	t.BoardRow[row] |= player << ((t.BoardSize + 2 - col) * 2)
	t.BoardCol[col] |= player << ((t.BoardSize + 2 - row) * 2)
	return nil
}

func (t *TicTacToe) SetValueToEmpty(row, col, value int) error {
	if row < 0 || col < 0 || row > t.BoardSize+1 || col > t.BoardSize+1 {
		return errors.New("out of bound")
	}

	t.BoardRow[row] |= value << ((t.BoardSize + 2 - col) * 2)
	t.BoardCol[col] |= value << ((t.BoardSize + 2 - row) * 2)
	return nil
}

func (t *TicTacToe) SetValue(row, col, value int) error {
	if row < 0 || col < 0 || row > t.BoardSize+1 || col > t.BoardSize+1 {
		return errors.New("out of bound")
	}

	// clear bit to 00
	t.BoardRow[row] &= ^(0b11 << ((t.BoardSize + 2 - col) * 2))
	t.BoardCol[col] &= ^(0b11 << ((t.BoardSize + 2 - row) * 2))
	// set bit to value
	t.BoardRow[row] |= value << ((t.BoardSize + 2 - col) * 2)
	t.BoardCol[col] |= value << ((t.BoardSize + 2 - row) * 2)
	return nil
}

func (t *TicTacToe) GetEmptyCells() []Move {
	var emptyCells []Move
	for i := 1; i < t.BoardSize+1; i++ {
		for j := 0; j < t.BoardSize+1; j++ {
			if t.GetValue(i, j) == EMPTY {
				emptyCells = append(emptyCells, Move{i, j})
			}
		}
	}
	return emptyCells
}

func (t *TicTacToe) UpdateDist(move Move) {
	// update dist
	t.Dist[move.Row][move.Col] = 0
	// update 8 direction
	for _, direction := range DIRECTION {
		for i := 1; i <= MAX_DIST; i++ {
			newRow := move.Row + direction.Row*i
			newCol := move.Col + direction.Col*i
			if newRow < 1 || newRow >= t.BoardSize+1 || newCol < 1 || newCol >= t.BoardSize+1 || t.GetValue(newRow, newCol) != EMPTY {
				continue
			}
			t.Dist[newRow][newCol] = min(t.Dist[newRow][newCol], i)
		}
	}
}

func (t *TicTacToe) CheckWin() bool {
	for i := range t.BoardRow {
		if VectorPatternMatchCount(t.BoardRow[i], 0b0101010101, 10) > 0 {
			return true
		}
		if VectorPatternMatchCount(t.BoardRow[i], 0b1010101010, 10) > 0 {
			return true
		}
	}

	for i := range t.BoardCol {
		if VectorPatternMatchCount(t.BoardCol[i], 0b0101010101, 10) > 0 {
			return true
		}
		if VectorPatternMatchCount(t.BoardCol[i], 0b1010101010, 10) > 0 {
			return true
		}
	}

	diags := t.getAllDiag()
	for i := range diags {
		if VectorPatternMatchCount(diags[i], 0b0101010101, 10) > 0 {
			return true
		}
		if VectorPatternMatchCount(diags[i], 0b1010101010, 10) > 0 {
			return true
		}
	}

	antiDiags := t.getAllAntiDiag()
	for i := range antiDiags {
		if VectorPatternMatchCount(antiDiags[i], 0b0101010101, 10) > 0 {
			return true
		}
		if VectorPatternMatchCount(antiDiags[i], 0b1010101010, 10) > 0 {
			return true
		}
	}

	return false
}

// GetWinningLine returns the positions of the 5 winning pieces
func (t *TicTacToe) GetWinningLine() []Move {
	// Check rows
	for row := 1; row <= t.BoardSize; row++ {
		for col := 1; col <= t.BoardSize-WIN_SIZE+1; col++ {
			player := t.GetValue(row, col)
			if player == EMPTY || player == WALL {
				continue
			}

			// Check horizontal
			match := true
			for k := 1; k < WIN_SIZE; k++ {
				if t.GetValue(row, col+k) != player {
					match = false
					break
				}
			}
			if match {
				line := make([]Move, WIN_SIZE)
				for k := 0; k < WIN_SIZE; k++ {
					line[k] = Move{Row: row, Col: col + k}
				}
				return line
			}
		}
	}

	// Check columns
	for row := 1; row <= t.BoardSize-WIN_SIZE+1; row++ {
		for col := 1; col <= t.BoardSize; col++ {
			player := t.GetValue(row, col)
			if player == EMPTY || player == WALL {
				continue
			}

			// Check vertical
			match := true
			for k := 1; k < WIN_SIZE; k++ {
				if t.GetValue(row+k, col) != player {
					match = false
					break
				}
			}
			if match {
				line := make([]Move, WIN_SIZE)
				for k := 0; k < WIN_SIZE; k++ {
					line[k] = Move{Row: row + k, Col: col}
				}
				return line
			}
		}
	}

	// Check diagonals (top-left to bottom-right)
	for row := 1; row <= t.BoardSize-WIN_SIZE+1; row++ {
		for col := 1; col <= t.BoardSize-WIN_SIZE+1; col++ {
			player := t.GetValue(row, col)
			if player == EMPTY || player == WALL {
				continue
			}

			match := true
			for k := 1; k < WIN_SIZE; k++ {
				if t.GetValue(row+k, col+k) != player {
					match = false
					break
				}
			}
			if match {
				line := make([]Move, WIN_SIZE)
				for k := 0; k < WIN_SIZE; k++ {
					line[k] = Move{Row: row + k, Col: col + k}
				}
				return line
			}
		}
	}

	// Check anti-diagonals (top-right to bottom-left)
	for row := 1; row <= t.BoardSize-WIN_SIZE+1; row++ {
		for col := WIN_SIZE; col <= t.BoardSize; col++ {
			player := t.GetValue(row, col)
			if player == EMPTY || player == WALL {
				continue
			}

			match := true
			for k := 1; k < WIN_SIZE; k++ {
				if t.GetValue(row+k, col-k) != player {
					match = false
					break
				}
			}
			if match {
				line := make([]Move, WIN_SIZE)
				for k := 0; k < WIN_SIZE; k++ {
					line[k] = Move{Row: row + k, Col: col - k}
				}
				return line
			}
		}
	}

	return nil
}

func (t *TicTacToe) MakeMove(move Move) error {
	err := t.SetMove(move.Row, move.Col)
	if err != nil {
		return errors.New("cannot make this move, error: " + err.Error())
	}
	t.UpdateDist(move)

	t.MoveCount++
	return nil
}

func (t *TicTacToe) UndoMove(move Move) {
	t.SetValue(move.Row, move.Col, EMPTY)
	t.MoveCount--
}
