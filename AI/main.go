package main

import (
	t3 "AI/game_v3"
	"fmt"
)

func main() {
	// t.WIN_SIZE = 5
	// t.MAX_DEPTH = 3
	// t.MAX_DIST = 2
	// t.MAX_TIME = 10
	// game := t.NewTicTacToe(15)
	// game.MakeMove(t.Move{Row: 8, Col: 8}, t.PLAYER_X)
	// game.MakeMove(t.Move{Row: 8, Col: 9}, t.PLAYER_O)
	// game.MakeMove(t.Move{Row: 7, Col: 9}, t.PLAYER_X)
	// game.PlayPvAI(t.PLAYER_X)
	t3.MAX_DEPTH = 2
	t3.MAX_TIME = 60
	t3.MAX_DIST = 1
	t3.DIST = 1
	game3 := t3.NewTicTacToe(15)

	// game3.MakeMove(t3.Move{Row: 4, Col: 4})
	// game3.MakeMove(t3.Move{Row: 4, Col: 3})
	// game3.MakeMove(t3.Move{Row: 4, Col: 5})
	// game3.MakeMove(t3.Move{Row: 4, Col: 2})
	// game3.MakeMove(t3.Move{Row: 4, Col: 6})
	// game3.MakeMove(t3.Move{Row: 4, Col: 1})
	// game3.MakeMove(t3.Move{Row: 4, Col: 7})
	// game3.MakeMove(t3.Move{Row: 3, Col: 1})
	// game3.MakeMove(t3.Move{Row: 4, Col: 8})
	fmt.Printf("%b\n", game3.BoardRow)
	game3.PlayPvAI(t3.PLAYER_X)

}
