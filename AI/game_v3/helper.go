package game3

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/bits"
	"net/http"
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
			println("eval_param: ", EVAL_PARAM, " move_count: ", t.MoveCount)
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
	PATTERN = map[string][]Pattern{
		"X five":      {{GetBinaryPattern([]int{x, x, x, x, x}), 10}},
		"X open four": {{GetBinaryPattern([]int{e, x, x, x, x, e}), 12}},
		"X close four": {
			{GetBinaryPattern([]int{o, x, x, x, x, e}), 12},
			{GetBinaryPattern([]int{w, x, x, x, x, e}), 12},
			{GetBinaryPattern([]int{e, x, x, x, x, o}), 12},
			{GetBinaryPattern([]int{e, x, x, x, x, w}), 12},
		},
		"X broken four": {
			{GetBinaryPattern([]int{x, e, x, x, x}), 10},
			{GetBinaryPattern([]int{x, x, e, x, x}), 10},
			{GetBinaryPattern([]int{x, x, x, e, x}), 10},
		},
		"X open three": {
			{GetBinaryPattern([]int{e, x, x, x, e}), 10},
		},
		"X broken three": {
			{GetBinaryPattern([]int{e, x, e, x, x, e}), 12},
			{GetBinaryPattern([]int{e, x, x, e, x, e}), 12},
		},
		"X close three": {
			{GetBinaryPattern([]int{o, x, x, x, e}), 10},
			{GetBinaryPattern([]int{e, x, x, x, o}), 10},
			{GetBinaryPattern([]int{w, x, x, x, e}), 10},
			{GetBinaryPattern([]int{e, x, x, x, w}), 10},
			{GetBinaryPattern([]int{o, x, e, x, x, e}), 12},
			{GetBinaryPattern([]int{w, x, x, e, x, e}), 12},
			{GetBinaryPattern([]int{e, x, e, x, x, o}), 12},
			{GetBinaryPattern([]int{e, x, x, e, x, w}), 12},
		},
		"X open two":   {{GetBinaryPattern([]int{e, x, x, e}), 8}},
		"X broken two": {{GetBinaryPattern([]int{e, x, e, x, e}), 10}},

		"O five":      {{GetBinaryPattern([]int{o, o, o, o, o}), 10}},
		"O open four": {{GetBinaryPattern([]int{e, o, o, o, o, e}), 12}},
		"O close four": {
			{GetBinaryPattern([]int{x, o, o, o, o, e}), 12},
			{GetBinaryPattern([]int{w, o, o, o, o, e}), 12},
			{GetBinaryPattern([]int{e, o, o, o, o, x}), 12},
			{GetBinaryPattern([]int{e, o, o, o, o, w}), 12},
		},
		"O broken four": {
			{GetBinaryPattern([]int{o, e, o, o, o}), 10},
			{GetBinaryPattern([]int{o, o, e, o, o}), 10},
			{GetBinaryPattern([]int{o, o, o, e, o}), 10},
		},
		"O open three": {
			{GetBinaryPattern([]int{e, o, o, o, e}), 10},
		},
		"O close three": {
			{GetBinaryPattern([]int{x, o, o, o, e}), 10},
			{GetBinaryPattern([]int{e, o, o, o, x}), 10},
			{GetBinaryPattern([]int{w, o, o, o, e}), 10},
			{GetBinaryPattern([]int{e, o, o, o, w}), 10},
			{GetBinaryPattern([]int{x, o, e, o, o, e}), 12},
			{GetBinaryPattern([]int{w, o, o, e, o, e}), 12},
			{GetBinaryPattern([]int{e, o, e, o, o, x}), 12},
			{GetBinaryPattern([]int{e, o, o, e, o, w}), 12},
		},
		"O broken three": {
			{GetBinaryPattern([]int{e, o, e, o, o, e}), 12},
			{GetBinaryPattern([]int{e, o, o, e, o, e}), 12},
		},
		"O open two":   {{GetBinaryPattern([]int{e, o, o, e}), 8}},
		"O broken two": {{GetBinaryPattern([]int{e, o, e, o, e}), 10}},
	}

}

func (t *TicTacToe) PrintNeighbors() {
	moves := t.GetNeighbor(DIST)
	fmt.Println("Neighbors:", len(moves), moves)
}

func HttpServer() {
	game := &TicTacToe{}

	type InitRequest struct {
		BoardSize int
		Player    string
	}

	type InitResponse struct {
		Success   bool
		FirstMove Move
	}

	http.HandleFunc("POST /init", func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		var req InitRequest
		err := d.Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(req)
		if req.Player == "X" {
			AI_PLAYER = PLAYER_X
		} else {
			AI_PLAYER = PLAYER_O
		}
		game = NewTicTacToe(req.BoardSize)
		if AI_PLAYER == PLAYER_X {
			firstMove := Move{Row: game.BoardSize / 2, Col: game.BoardSize / 2}
			game.MakeMove(firstMove)
			json.NewEncoder(w).Encode(InitResponse{Success: true, FirstMove: firstMove})
		} else {
			json.NewEncoder(w).Encode(InitResponse{Success: true, FirstMove: Move{-1, -1}})
		}
	})

	http.HandleFunc("POST /move", func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		var move Move
		err := d.Decode(&move)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = game.MakeMove(move)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		nextMove := game.GetBestMove()
		err = game.MakeMove(nextMove)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		game.PrintBoard()

		json.NewEncoder(w).Encode(struct {
			Move     Move
			GameOver bool
		}{nextMove, game.CheckWin()})
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
