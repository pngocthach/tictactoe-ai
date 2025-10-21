package dto

// CreateGameRequest represents the request to create a new game
type CreateGameRequest struct {
	BoardSize  int    `json:"board_size"`
	Player     int    `json:"player"`     // PLAYER_X (1) or PLAYER_O (2) - who is the human player
	Difficulty string `json:"difficulty"` // "easy" or "hard"
}

// MakeMoveRequest represents the request to make a move
type MakeMoveRequest struct {
	GameID string `json:"game_id"`
	Row    int    `json:"row"`
	Col    int    `json:"col"`
}
