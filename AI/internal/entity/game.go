package entity

import (
	v3 "AI/internal/game/v3"
	"time"
)

// Game represents a game instance
type Game struct {
	ID        string
	Instance  *v3.TicTacToe
	BoardSize int
	Player    int // PLAYER_X or PLAYER_O - who is the human player
	CreatedAt time.Time
	UpdatedAt time.Time
}
