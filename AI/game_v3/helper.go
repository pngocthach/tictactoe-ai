package game3

import (
	"fmt"
	"math"
	"math/bits"
)

func (t *TicTacToe) PrintDist() {
	// neu dist = max_int thi in ra -
	for i := 0; i < t.BoardSize+2; i++ {
		for j := 0; j < t.BoardSize+2; j++ {
			if t.Dist[i][j] == math.MaxInt32 {
				fmt.Print("- ")
			} else {
				fmt.Print(t.Dist[i][j], " ")
			}
		}
		fmt.Println()
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

func (t *TicTacToe) PlayPvP() {
	player := PLAYER_X
	for !t.CheckWin() {
		t.PrintDist()
		t.PrintBoard()
		fmt.Println("Player", player)
		// get move from user input
		move := getUserInput()
		err := t.MakeMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		player = t.GetOpponent(player)
	}
	t.PrintBoard()

	if player == PLAYER_X {
		fmt.Println("Player O wins")
	} else {
		fmt.Println("Player X wins")
	}
}

// GetOpponent returns the opponent of the given player
func (t *TicTacToe) GetOpponent(player int) int {
	if player == PLAYER_X {
		return PLAYER_O
	}
	return PLAYER_X
}

func PrintLine(line int) {
	lineLen := bits.Len64(uint64(line)) / 2

	for i := lineLen - 1; i >= 0; i-- {
		value := (line >> (i * 2)) & 0b11
		if value == PLAYER_X {
			print("X")
		} else if value == PLAYER_O {
			print("O")
		} else if value == WALL {
			print("|")
		} else {
			print("-")
		}
	}
	println()
}

func (t *TicTacToe) PrintBoard() {

	print("  |")
	for i := 0; i < t.BoardSize+2; i++ {
		print(i, "|")
	}
	println()

	for i := 0; i < t.BoardSize+2; i++ {
		if i < 10 {
			print(" ")
		}
		print(i, "|")
		for j := 0; j < t.BoardSize+2; j++ {
			value := t.GetValue(i, j)
			switch value {
			case PLAYER_X:
				print("X")
			case PLAYER_O:
				print("O")
			case EMPTY:
				print(" ")
			case WALL:
				print("-")
			}
			print("|")
		}
		println()
	}
}

func (t *TicTacToe) PlayPvAI(AIplayer int) {
	// if t.MoveCount == 0 {
	// 	t.MakeMove(Move{t.BoardSize / 2, t.BoardSize / 2})
	// }
	// t.MakeMove(Move{Row: t.BoardSize / 2, Col: t.BoardSize / 2}, PLAYER_X)
	player := PLAYER_X
	for !t.CheckWin() {
		// t.PrintDist()
		t.PrintNeighbors()
		t.PrintBoard()
		fmt.Println("Player", player)
		fmt.Println("cache hit: ", cacheHit)
		if player == AIplayer {
			move := t.GetBestMove()
			if t.MoveCount == 0 {
				move = Move{t.BoardSize / 2, t.BoardSize / 2}
			}
			t.MakeMove(move)
			println("AI move to", move.Row, move.Col)
		} else {
			move := getUserInput()
			err := t.MakeMove(move)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		player = t.GetOpponent(player)
	}
	// t.PrintDist()
	t.PrintNeighbors()
	t.PrintBoard()
	if player == PLAYER_X {
		fmt.Println("Player O wins")
	} else {
		fmt.Println("Player X wins")
	}
}

func GetBinaryPattern(pattern []int) int {
	res := 0
	for i := 0; i < len(pattern); i++ {
		res = (res << 2) | pattern[i]
	}
	return res
}

func InitPattern() {
	PATTERN = map[string][]int{
		"X five":      {GetBinaryPattern([]int{x, x, x, x, x})},
		"X open four": {GetBinaryPattern([]int{e, x, x, x, x, e})},
		"X close four": {
			GetBinaryPattern([]int{o, x, x, x, x, e}),
			GetBinaryPattern([]int{w, x, x, x, x, e}),
			GetBinaryPattern([]int{e, x, x, x, x, o}),
			GetBinaryPattern([]int{e, x, x, x, x, w}),
		},
		"X broken four": {
			GetBinaryPattern([]int{x, e, x, x, x}),
			GetBinaryPattern([]int{x, x, e, x, x}),
			GetBinaryPattern([]int{x, x, x, e, x}),
		},
		"X open three": {
			GetBinaryPattern([]int{e, x, x, x, e}),
		},
		"X broken three": {
			GetBinaryPattern([]int{e, x, e, x, x, e}),
			GetBinaryPattern([]int{e, x, x, e, x, e}),
		},
		"X open two":   {GetBinaryPattern([]int{e, x, x, e})},
		"X broken two": {GetBinaryPattern([]int{e, x, e, x, e})},

		"O five":      {GetBinaryPattern([]int{o, o, o, o, o})},
		"O open four": {GetBinaryPattern([]int{e, o, o, o, o, e})},
		"O close four": {
			GetBinaryPattern([]int{x, o, o, o, o, e}),
			GetBinaryPattern([]int{w, o, o, o, o, e}),
			GetBinaryPattern([]int{e, o, o, o, o, x}),
			GetBinaryPattern([]int{e, o, o, o, o, w}),
		},
		"O broken four": {
			GetBinaryPattern([]int{o, e, o, o, o}),
			GetBinaryPattern([]int{o, o, e, o, o}),
			GetBinaryPattern([]int{o, o, o, e, o}),
		},
		"O open three": {
			GetBinaryPattern([]int{e, o, o, o, e}),
		},
		"O broken three": {
			GetBinaryPattern([]int{e, o, e, o, o, e}),
			GetBinaryPattern([]int{e, o, o, e, o, e}),
		},
		"O open two":   {GetBinaryPattern([]int{e, o, o, e})},
		"O broken two": {GetBinaryPattern([]int{e, o, e, o, e})},
	}
}

func (t *TicTacToe) PrintNeighbors() {
	moves := t.GetNeighbor(DIST)
	fmt.Println("Neighbors:", len(moves), moves)
}
