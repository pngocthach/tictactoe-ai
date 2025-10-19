package dto

// CreateGameResponse represents the response after creating a game
type CreateGameResponse struct {
	GameID    string    `json:"game_id"`
	BoardSize int       `json:"board_size"`
	Player    int       `json:"player"`
	AIMove    *MoveInfo `json:"ai_move,omitempty"` // First AI move if player goes second
	Board     [][]int   `json:"board"`
	Message   string    `json:"message"`
}

// MakeMoveResponse represents the response after making a move
type MakeMoveResponse struct {
	GameID   string    `json:"game_id"`
	Row      int       `json:"row"`
	Col      int       `json:"col"`
	AIMove   *MoveInfo `json:"ai_move,omitempty"`
	Board    [][]int   `json:"board"`
	GameOver bool      `json:"game_over"`
	Winner   *int      `json:"winner,omitempty"`
	Message  string    `json:"message"`
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
