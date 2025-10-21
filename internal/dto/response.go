package dto

// CreateGameResponse represents the response after creating a game
type CreateGameResponse struct {
	GameID    string    `json:"game_id"`
	BoardSize int       `json:"board_size"`
	Player    int       `json:"player"`
	AIMove    *MoveInfo `json:"ai_move,omitempty"` // AI move (for highlighting)
	Board     [][]int   `json:"board"`
	Message   string    `json:"message"`
}

// MakeMoveResponse represents the response after making a move
type MakeMoveResponse struct {
	GameID      string     `json:"game_id"`
	Row         int        `json:"row"`
	Col         int        `json:"col"`
	AIMove      *MoveInfo  `json:"ai_move,omitempty"` // AI move (for highlighting)
	Board       [][]int    `json:"board"`
	GameOver    bool       `json:"game_over"`
	Winner      *int       `json:"winner,omitempty"`
	WinningLine []MoveInfo `json:"winning_line,omitempty"`
	Message     string     `json:"message"`
}

// MoveInfo represents information about a move
type MoveInfo struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
