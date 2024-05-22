package game

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
	"time"

	"github.com/huandu/go-clone"
)

type Move struct {
	Row int
	Col int
}

// tao struct tic-tac-toe nxn
type TicTacToe struct {
	Board     [][]int
	boardSize int
	winSize   int
	MoveCount int
	Dist      [][]int
}

// tao ham new tic-tac-toe
func NewTicTacToe(boardSize int) *TicTacToe {
	board := make([][]int, boardSize+2)
	for i := range board {
		board[i] = make([]int, boardSize+2)
	}
	// set wall for board
	for i := 0; i < boardSize+2; i++ {
		board[0][i] = W
		board[boardSize+1][i] = W
		board[i][0] = W
		board[i][boardSize+1] = W
	}

	// init dist intmax
	dist := make([][]int, boardSize+2)
	for i := range dist {
		dist[i] = make([]int, boardSize+2)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	return &TicTacToe{board, boardSize, WIN_SIZE, 0, dist}
}

// tao ham check win
// func (t *TicTacToe) CheckWin(player int) bool {
// 	// check hang
// 	for i := 1; i < t.boardSize+1; i++ {
// 		for j := 1; j <= t.boardSize-t.winSize+1; j++ {
// 			count := 0
// 			for k := 0; k < t.winSize+1; k++ {
// 				if t.Board[i][j+k] == player {
// 					count++
// 				}
// 			}
// 			if count == t.winSize {
// 				return true
// 			}
// 		}
// 	}
// 	// check cot
// 	for i := 0; i < t.boardSize; i++ {
// 		for j := 0; j <= t.boardSize-t.winSize; j++ {
// 			count := 0
// 			for k := 0; k < t.winSize; k++ {
// 				if t.Board[j+k][i] == player {
// 					count++
// 				}
// 			}
// 			if count == t.winSize {
// 				return true
// 			}
// 		}
// 	}
// 	// check duong cheo chinh
// 	for i := 0; i <= t.boardSize+2-t.winSize; i++ {
// 		for j := 0; j <= t.boardSize+2-t.winSize; j++ {
// 			count := 0
// 			for k := 0; k < t.winSize; k++ {
// 				if t.Board[i+k][j+k] == player {
// 					count++
// 				}
// 			}
// 			if count == t.winSize {
// 				return true
// 			}
// 		}
// 	}
// 	// check duong cheo phu
// 	for i := 0; i <= t.boardSize+2-t.winSize; i++ {
// 		for j := t.winSize - 1; j < t.boardSize+2; j++ {
// 			count := 0
// 			for k := 0; k < t.winSize; k++ {
// 				if t.Board[i+k][j-k] == player {
// 					count++
// 				}
// 			}
// 			if count == t.winSize {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func (t *TicTacToe) CheckWin(player int) bool {
	for i := 1; i < t.boardSize+1; i++ {
		for j := 1; j < t.boardSize+1; j++ {
			if t.Board[i][j] == player {
				for _, direction := range DIRECTIONS {
					count := 1
					for k := 1; k < t.winSize; k++ {
						newRow := i + direction[0]*k
						newCol := j + direction[1]*k
						if newRow < 1 || newRow >= t.boardSize+1 || newCol < 1 || newCol >= t.boardSize+1 || t.Board[newRow][newCol] != player {
							break
						}
						count++
					}
					if count == t.winSize {
						return true
					}
				}
			}
		}
	}
	return false
}

// tao ham move
func (t *TicTacToe) MakeMove(move Move, player int) error {
	if move.Row < 1 || move.Row >= t.boardSize+1 || move.Col < 1 || move.Col >= t.boardSize+1 {
		return fmt.Errorf("invalid move: (%d, %d)", move.Row, move.Col)
	}
	if t.Board[move.Row][move.Col] != EMPTY {
		return fmt.Errorf("invalid move: cell (%d, %d) is already occupied", move.Row, move.Col)
	}
	// println("player", player, "move to", move.Row, move.Col, "eval: ", int(t.Evaluate()*10), " - ", t.Evaluate())
	t.Board[move.Row][move.Col] = player
	t.MoveCount++

	// update dist
	t.Dist[move.Row][move.Col] = 0
	// update 8 direction
	for _, direction := range DIRECTION {
		for i := 1; i <= MAX_DIST; i++ {
			newRow := move.Row + direction.Row*i
			newCol := move.Col + direction.Col*i
			if newRow < 1 || newRow >= t.boardSize+1 || newCol < 1 || newCol >= t.boardSize+1 || t.Board[newRow][newCol] != EMPTY {
				continue
			}
			t.Dist[newRow][newCol] = min(t.Dist[newRow][newCol], i)
		}
	}

	return nil
}

// get player turn
func (t *TicTacToe) GetPlayer() int {
	if t.MoveCount%2 == 0 {
		return PLAYER_X
	}
	return PLAYER_O
}

// get opponent
func (t *TicTacToe) GetOpponent() int {
	if t.GetPlayer() == PLAYER_X {
		return PLAYER_O
	}
	return PLAYER_X
}

// lay tat ca cac o con trong
func (t *TicTacToe) GetEmptyCells() []Move {
	var emptyCells []Move
	for i := 0; i < t.boardSize; i++ {
		for j := 0; j < t.boardSize; j++ {
			if t.Board[i][j] == EMPTY {
				emptyCells = append(emptyCells, Move{i, j})
			}
		}
	}
	return emptyCells
}

// tao ham play voi nguoi
func (t *TicTacToe) PlayPvP() {
	player := PLAYER_X
	for !t.CheckWin(PLAYER_X) && !t.CheckWin(PLAYER_O) {
		t.PrintBoard()
		fmt.Println("Player", player)
		// get move from user input
		move := getUserInput()
		err := t.MakeMove(move, player)
		if err != nil {
			fmt.Println(err)
			continue
		}
		player *= -1
	}

	if player == PLAYER_X {
		fmt.Println("Player O wins")
	} else {
		fmt.Println("Player X wins")
	}
}

func getUserInput() Move {
	var row, col int
	fmt.Print("Enter row: ")
	fmt.Scan(&row)
	fmt.Print("Enter col: ")
	fmt.Scan(&col)
	return Move{row, col}
}

// tao ham play voi AI
func (t *TicTacToe) PlayPvAI(AIplayer int) {
	if t.MoveCount == 0 {
		t.MakeMove(Move{t.boardSize / 2, t.boardSize / 2}, PLAYER_X)
	}
	// t.MakeMove(Move{Row: t.boardSize / 2, Col: t.boardSize / 2}, PLAYER_X)
	player := PLAYER_O
	for !t.CheckWin(PLAYER_X) && !t.CheckWin(PLAYER_O) {
		t.PrintDist()
		t.PrintNeighbors()
		t.PrintBoard()
		fmt.Println("Player", player)
		if player == AIplayer {
			move := t.GetBestMove()
			if t.MoveCount == 0 {
				move = Move{t.boardSize / 2, t.boardSize / 2}
			}
			t.MakeMove(move, player)
			println("AI move to", move.Row, move.Col)
		} else {
			move := getUserInput()
			err := t.MakeMove(move, player)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		player *= -1
	}
	t.PrintDist()
	t.PrintNeighbors()
	t.PrintBoard()
	if player == PLAYER_X {
		fmt.Println("Player O wins")
	} else {
		fmt.Println("Player X wins")
	}
}

func (t *TicTacToe) PrintDist() {
	// neu dist = max_int thi in ra -
	for i := 0; i < t.boardSize+2; i++ {
		for j := 0; j < t.boardSize+2; j++ {
			if t.Dist[i][j] == math.MaxInt32 {
				fmt.Print("- ")
			} else {
				fmt.Print(t.Dist[i][j], " ")
			}
		}
		fmt.Println()
	}
}

// tao ham print board
func (t *TicTacToe) PrintBoard() {

	print("  |")
	for i := 0; i < t.boardSize+2; i++ {

		print(i, "|")
	}
	println()

	for i := 0; i < t.boardSize+2; i++ {
		if i < 10 {
			print(" ")
		}
		print(i, "|")
		for j := 0; j < t.boardSize+2; j++ {
			switch t.Board[i][j] {
			case PLAYER_X:
				print("X")
			case PLAYER_O:
				print("O")
			case EMPTY:
				print(" ")
			case W:
				print("-")
			}
			print("|")
		}
		println()
	}
}

// tao ham get best move
func (t *TicTacToe) GetBestMove() Move {
	MiniMaxStart = time.Now()
	player := t.GetPlayer()
	if player == PLAYER_X {
		bestScore := math.Inf(-1)
		moves := t.GetNeighbor(DIST)
		var bestMove Move
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i], player)
			score := t.Minimax(0, math.Inf(-1), math.Inf(1))
			t.UndoMove(moves[i])
			t.Dist = oldDist
			if score > bestScore {
				bestScore = score
				bestMove = moves[i]
			}
		}
		fmt.Println("time: ", time.Since(MiniMaxStart).Seconds())
		return bestMove
	} else {
		bestScore := math.Inf(1)
		moves := t.GetNeighbor(DIST)
		bestMove := moves[rand.IntN(len(moves))]
		for i := 0; i < len(moves); i++ {
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i], player)
			score := t.Minimax(0, math.Inf(-1), math.Inf(1))
			t.UndoMove(moves[i])
			t.Dist = oldDist
			if score < bestScore {
				bestScore = score
				bestMove = moves[i]
			}
		}
		fmt.Println("time: ", time.Since(MiniMaxStart).Seconds())
		return bestMove
	}
}

// undo move
func (t *TicTacToe) UndoMove(move Move) {
	t.Board[move.Row][move.Col] = EMPTY
	t.MoveCount--
}

var MiniMaxStart = time.Now()

func (t *TicTacToe) Minimax(depth int, alpha float64, beta float64) float64 {
	// return t.Evaluate() after 10 seconds
	if time.Since(MiniMaxStart) > MAX_TIME*time.Second && depth > 1 {
		return t.Evaluate()
	}

	if depth == MAX_DEPTH || t.CheckWin(PLAYER_X) || t.CheckWin(PLAYER_O) || t.MoveCount == t.boardSize*t.boardSize {
		return t.Evaluate()
	}

	isMaximize := t.GetPlayer() == PLAYER_X
	var bestScore float64
	moves := t.GetNeighbor(DIST)
	if len(moves) == 0 {
		moves = []Move{{t.boardSize / 2, t.boardSize / 2}}
	}

	if isMaximize {
		bestScore = math.Inf(-1)
		randomIndex := 0
		if len(moves) > 0 {
			randomIndex = rand.IntN(len(moves))
		}
		for i := 0; i < randomIndex; i++ {
			// println("player", t.GetPlayer(), "move to", move.Row, move.Col, "eval: ", int(t.Evaluate()*10), " - ", t.Evaluate(), "depth: ", depth)
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i], PLAYER_X)
			score := t.Minimax(depth+1, alpha, beta)
			t.UndoMove(moves[i])
			t.Dist = oldDist
			bestScore = max(bestScore, score)
			alpha = max(alpha, bestScore)
			if beta <= alpha {
				break
			}
		}
		for i := randomIndex; i < len(moves); i++ {
			// println("player", t.GetPlayer(), "move to", move.Row, move.Col, "eval: ", int(t.Evaluate()*10), " - ", t.Evaluate(), "depth: ", depth)
			oldDist := clone.Clone(t.Dist).([][]int)
			t.MakeMove(moves[i], PLAYER_X)
			score := t.Minimax(depth+1, alpha, beta)
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
			t.MakeMove(moves[i], PLAYER_O)
			score := t.Minimax(depth+1, alpha, beta)
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

func (t *TicTacToe) GetNeighbor(dist int) []Move {
	var neighbors []Move
	for k := 1; k <= dist; k++ {
		for i := 0; i < t.boardSize; i++ {
			for j := 0; j < t.boardSize; j++ {
				if t.Dist[i][j] == k && t.Board[i][j] == EMPTY {
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

var DIRECTIONS = [][2]int{
	{1, 0},  // vertical
	{0, 1},  // horizontal
	{1, 1},  // diagonal
	{-1, 1}, // anti-diagonal
}

// tao ham get random move
func (t *TicTacToe) GetRandomMove() Move {
	emptyCells := t.GetEmptyCells()
	return emptyCells[0]
}
