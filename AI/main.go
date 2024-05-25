package main

import (
	t3 "AI/game_v3"
)

func main() {
	t3.MAX_DEPTH = 3
	t3.MAX_TIME = 60
	t3.MAX_DIST = 1
	t3.DIST = 1
	game3 := t3.NewTicTacToe(15)
	// game3.MakeMove(t3.Move{Row: 7, Col: 7})
	// game3.MakeMove(t3.Move{Row: 7, Col: 8})
	// game3.MakeMove(t3.Move{Row: 6, Col: 8})
	// game3.MakeMove(t3.Move{Row: 8, Col: 6})
	game3.PlayPvAI(t3.PLAYER_X)
}
