package service

import (
	"AI/internal/dto"
	"AI/internal/entity"
	v3 "AI/internal/game/v3"
	"AI/internal/repository"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
)

// GameService handles game logic
type GameService struct {
	repo *repository.GameRepository
}

// NewGameService creates a new game service
func NewGameService(repo *repository.GameRepository) *GameService {
	return &GameService{
		repo: repo,
	}
}

// generateUUID generates a simple UUID-like string
func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := hex.EncodeToString(b)
	return uuid[:8] + "-" + uuid[8:12] + "-" + uuid[12:16] + "-" + uuid[16:20] + "-" + uuid[20:], nil
}

// CreateGame creates a new game instance
func (s *GameService) CreateGame(req *dto.CreateGameRequest) (*dto.CreateGameResponse, error) {
	// Validate board size
	if req.BoardSize < 5 || req.BoardSize > 20 {
		return nil, errors.New("board size must be between 5 and 20")
	}

	// Validate player
	if req.Player != v3.PLAYER_X && req.Player != v3.PLAYER_O {
		return nil, errors.New("invalid player value")
	}

	// Validate and set difficulty
	difficulty := req.Difficulty
	if difficulty == "" {
		difficulty = "easy" // default
	}
	if difficulty != "easy" && difficulty != "hard" {
		return nil, errors.New("invalid difficulty, must be 'easy' or 'hard'")
	}

	// Generate UUID
	uuid, err := generateUUID()
	if err != nil {
		return nil, errors.New("failed to generate game ID")
	}

	// Initialize game configurations based on difficulty
	v3.SetDifficulty(difficulty)
	v3.MAX_TIME = 2
	v3.MAX_DIST = 1
	v3.DIST = 1

	// Create new game
	gameInstance := v3.NewTicTacToe(req.BoardSize)

	game := &entity.Game{
		ID:         uuid,
		Instance:   gameInstance,
		BoardSize:  req.BoardSize,
		Player:     req.Player,
		Difficulty: difficulty,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// If player goes second (PLAYER_O), AI makes first move
	var aiMove *dto.MoveInfo
	if req.Player == v3.PLAYER_O {
		// AI (PLAYER_X) makes the first move - usually center
		centerMove := v3.Move{Row: req.BoardSize/2 + 1, Col: req.BoardSize/2 + 1}
		if err := gameInstance.MakeMove(centerMove); err != nil {
			return nil, err
		}
		aiMove = &dto.MoveInfo{
			Row: centerMove.Row,
			Col: centerMove.Col,
		}
	}

	// Save to repository
	if err := s.repo.Save(game); err != nil {
		return nil, err
	}

	// Get board state
	board := s.getBoardState(gameInstance)

	return &dto.CreateGameResponse{
		GameID:    uuid,
		BoardSize: req.BoardSize,
		Player:    req.Player,
		AIMove:    aiMove,
		Board:     board,
		Message:   "Game created successfully",
	}, nil
}

// MakeMove processes a player move and optionally triggers AI move
func (s *GameService) MakeMove(req *dto.MakeMoveRequest) (*dto.MakeMoveResponse, error) {
	// Retrieve game
	game, err := s.repo.FindByID(req.GameID)
	if err != nil {
		return nil, err
	}

	// Set difficulty for this game
	v3.SetDifficulty(game.Difficulty)

	// Validate move coordinates
	if req.Row < 1 || req.Row > game.BoardSize || req.Col < 1 || req.Col > game.BoardSize {
		return nil, errors.New("move coordinates out of bounds")
	}

	// Check if it's the player's turn
	currentPlayer := game.Instance.GetPlayer()
	if currentPlayer != game.Player {
		return nil, errors.New("not your turn")
	}

	// Make the player's move
	move := v3.Move{Row: req.Row, Col: req.Col}
	if err := game.Instance.MakeMove(move); err != nil {
		return nil, err
	}

	// Check if game is over after player's move
	gameOver := game.Instance.CheckWin()
	var winner *int
	var aiMove *dto.MoveInfo
	var winningLine []dto.MoveInfo

	if gameOver {
		w := game.Player
		winner = &w
		// Get winning line
		line := game.Instance.GetWinningLine()
		if line != nil {
			winningLine = make([]dto.MoveInfo, len(line))
			for i, move := range line {
				winningLine[i] = dto.MoveInfo{Row: move.Row, Col: move.Col}
			}
		}
	} else if game.Instance.MoveCount == game.BoardSize*game.BoardSize {
		// Draw
		gameOver = true
	} else {
		// AI's turn
		aiMoveResult := game.Instance.GetBestMove()
		if err := game.Instance.MakeMove(aiMoveResult); err != nil {
			return nil, err
		}

		aiMove = &dto.MoveInfo{
			Row: aiMoveResult.Row,
			Col: aiMoveResult.Col,
		}

		// Check if game is over after AI's move
		if game.Instance.CheckWin() {
			gameOver = true
			aiPlayer := game.Instance.GetOpponent(game.Player)
			winner = &aiPlayer
			// Get winning line
			line := game.Instance.GetWinningLine()
			if line != nil {
				winningLine = make([]dto.MoveInfo, len(line))
				for i, move := range line {
					winningLine[i] = dto.MoveInfo{Row: move.Row, Col: move.Col}
				}
			}
		} else if game.Instance.MoveCount == game.BoardSize*game.BoardSize {
			// Draw
			gameOver = true
		}
	}

	// Update game
	game.UpdatedAt = time.Now()
	if err := s.repo.Update(game); err != nil {
		return nil, err
	}

	// Convert board to 2D array for response
	board := s.getBoardState(game.Instance)

	return &dto.MakeMoveResponse{
		GameID:      req.GameID,
		Row:         req.Row,
		Col:         req.Col,
		AIMove:      aiMove,
		Board:       board,
		GameOver:    gameOver,
		Winner:      winner,
		WinningLine: winningLine,
		Message:     "Move processed successfully",
	}, nil
}

// getBoardState converts the game board to a 2D array
func (s *GameService) getBoardState(game *v3.TicTacToe) [][]int {
	board := make([][]int, game.BoardSize)
	for i := 0; i < game.BoardSize; i++ {
		board[i] = make([]int, game.BoardSize)
		for j := 0; j < game.BoardSize; j++ {
			// Adjusting for 1-indexed game board
			board[i][j] = game.GetValue(i+1, j+1)
		}
	}
	return board
}
