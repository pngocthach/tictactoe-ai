package game

import "testing"

func TestCheckWin(t *testing.T) {
	t.Run("Player X wins horizontally", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{W, W, W, W, W, W, W},
			{W, X, X, X, X, X, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, W, W, W, W, W, W},
		}
		if !game.CheckWin(PLAYER_X) {
			t.Error("Expected Player X to win horizontally")
		}
	})

	t.Run("Player O wins vertically", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{W, W, W, W, W, W, W},
			{W, O, O, O, O, O, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, W, W, W, W, W, W},
		}
		if !game.CheckWin(PLAYER_O) {
			t.Error("Expected Player O to win vertically")
		}
	})

	t.Run("Player X wins diagonally (main diagonal)", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{W, W, W, W, W, W, W},
			{W, X, O, O, O, O, W},
			{W, 0, X, 0, 0, 0, W},
			{W, 0, 0, X, 0, 0, W},
			{W, 0, 0, 0, X, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, W, W, W, W, W, W},
		}
		if !game.CheckWin(PLAYER_X) {
			t.Error("Expected Player X to win diagonally (main diagonal)")
		}
	})

	t.Run("Player O wins diagonally (anti-diagonal)", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{W, W, W, W, W, W, W},
			{W, X, O, O, O, O, W},
			{W, 0, X, 0, O, 0, W},
			{W, 0, 0, O, 0, 0, W},
			{W, 0, O, 0, O, 0, W},
			{W, O, 0, 0, 0, X, W},
			{W, W, W, W, W, W, W},
		}
		if !game.CheckWin(PLAYER_O) {
			t.Error("Expected Player O to win diagonally (anti-diagonal)")
		}
	})

	t.Run("No winner", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{W, W, W, W, W, W, W},
			{W, X, O, O, O, O, W},
			{W, 0, 0, 0, O, 0, W},
			{W, 0, 0, 0, 0, 0, W},
			{W, 0, O, 0, O, 0, W},
			{W, O, 0, 0, 0, X, W},
			{W, W, W, W, W, W, W},
		}
		if game.CheckWin(PLAYER_X) || game.CheckWin(PLAYER_O) {
			t.Error("Expected no winner")
		}
	})
}
func TestCheckWin2(t *testing.T) {
	t.Run("Player X wins horizontally", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{X, X, X, X, X},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		if !game.CheckWin(PLAYER_X) {
			t.Error("Expected Player X to win horizontally")
		}
	})

	t.Run("Player O wins vertically", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{-1, 0, 0, 0, 0},
			{-1, 0, 0, 0, 0},
			{-1, 0, 0, 0, 0},
			{-1, 0, 0, 0, 0},
			{-1, 0, 0, 0, 0},
		}
		if !game.CheckWin(PLAYER_O) {
			t.Error("Expected Player O to win vertically")
		}
	})

	t.Run("Player X wins diagonally (main diagonal)", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		}
		if !game.CheckWin(PLAYER_X) {
			t.Error("Expected Player X to win diagonally (main diagonal)")
		}
	})

	t.Run("Player O wins diagonally (anti-diagonal)", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{0, 0, 0, 0, -1},
			{0, 0, 0, -1, 0},
			{0, 0, -1, 0, 0},
			{0, -1, 0, 0, 0},
			{-1, 0, 0, 0, 0},
		}
		if !game.CheckWin(PLAYER_O) {
			t.Error("Expected Player O to win diagonally (anti-diagonal)")
		}
	})

	t.Run("No winner", func(t *testing.T) {
		game := NewTicTacToe(5)
		game.Board = [][]int{
			{1, -1, 1, -1, 1},
			{-1, 1, -1, 1, -1},
			{1, -1, -1, -1, 1},
			{-1, 1, -1, -1, -1},
			{1, -1, 1, -1, 1},
		}
		if game.CheckWin(PLAYER_X) || game.CheckWin(PLAYER_O) {
			t.Error("Expected no winner")
		}
	})
}
