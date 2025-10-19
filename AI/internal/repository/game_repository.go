package repository

import (
	"AI/internal/entity"
	"errors"
	"sync"
)

// GameRepository manages game storage in memory
type GameRepository struct {
	games map[string]*entity.Game
	mu    sync.RWMutex
}

// NewGameRepository creates a new game repository
func NewGameRepository() *GameRepository {
	return &GameRepository{
		games: make(map[string]*entity.Game),
	}
}

// Save stores a game in memory
func (r *GameRepository) Save(game *entity.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if game.ID == "" {
		return errors.New("game ID cannot be empty")
	}

	r.games[game.ID] = game
	return nil
}

// FindByID retrieves a game by ID
func (r *GameRepository) FindByID(id string) (*entity.Game, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	game, exists := r.games[id]
	if !exists {
		return nil, errors.New("game not found")
	}

	return game, nil
}

// Update updates an existing game
func (r *GameRepository) Update(game *entity.Game) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.games[game.ID]; !exists {
		return errors.New("game not found")
	}

	r.games[game.ID] = game
	return nil
}

// Delete removes a game from storage
func (r *GameRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.games[id]; !exists {
		return errors.New("game not found")
	}

	delete(r.games, id)
	return nil
}

// Count returns the number of games in storage
func (r *GameRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.games)
}
