package handler

import (
	"AI/internal/dto"
	"AI/internal/repository"
	"AI/internal/service"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// WebHandler handles HTTP requests for web interface
type WebHandler struct {
	service   *service.GameService
	repo      *repository.GameRepository
	templates *template.Template
}

// NewWebHandler creates a new web handler
func NewWebHandler(service *service.GameService, repo *repository.GameRepository) *WebHandler {
	// Get absolute path to templates
	wd, _ := os.Getwd()
	tmplPath := filepath.Join(wd, "web", "templates", "*.html")

	tmpl, err := template.ParseGlob(tmplPath)
	if err != nil {
		fmt.Printf("Error parsing templates from %s: %v\n", tmplPath, err)
		// Try alternative path
		tmplPath = "web/templates/*.html"
		tmpl, err = template.ParseGlob(tmplPath)
		if err != nil {
			fmt.Printf("Error parsing templates from %s: %v\n", tmplPath, err)
			panic(err)
		}
	}

	fmt.Printf("Templates loaded successfully from: %s\n", tmplPath)

	return &WebHandler{
		service:   service,
		repo:      repo,
		templates: tmpl,
	}
}

// Index serves the main page
func (h *WebHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Only handle root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := h.templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		return
	}
}

// CreateGame handles game creation from web
func (h *WebHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	boardSize, err := strconv.Atoi(r.FormValue("board_size"))
	if err != nil {
		http.Error(w, "Invalid board size", http.StatusBadRequest)
		return
	}

	player, err := strconv.Atoi(r.FormValue("player"))
	if err != nil {
		http.Error(w, "Invalid player", http.StatusBadRequest)
		return
	}

	// Create game
	req := &dto.CreateGameRequest{
		BoardSize: boardSize,
		Player:    player,
	}

	response, err := h.service.CreateGame(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Render game board HTML
	h.renderGameBoard(w, response, "")
}

// MakeMove handles player moves from web
func (h *WebHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	gameID := r.FormValue("game_id")
	row, _ := strconv.Atoi(r.FormValue("row"))
	col, _ := strconv.Atoi(r.FormValue("col"))

	req := &dto.MakeMoveRequest{
		GameID: gameID,
		Row:    row,
		Col:    col,
	}

	response, err := h.service.MakeMove(req)
	if err != nil {
		// Return error message in HTML
		fmt.Fprintf(w, `<div class="game-info"><div class="status" style="color: red;">Error: %s</div></div>`, err.Error())
		return
	}

	// Prepare AI move message
	aiMoveMsg := ""
	if response.AIMove != nil {
		aiMoveMsg = fmt.Sprintf("AI played: Row %d, Col %d", response.AIMove.Row, response.AIMove.Col)
	}

	// Get player from repository
	game, err := h.repo.FindByID(response.GameID)
	player := 1
	if err == nil && game != nil {
		player = game.Player
	}

	// If game is over, include winner message
	if response.GameOver {
		winnerMsg := ""
		if response.Winner != nil {
			if *response.Winner == 1 {
				winnerMsg = "🎉 X Wins!"
			} else if *response.Winner == 2 {
				winnerMsg = "🎉 O Wins!"
			}
		} else {
			winnerMsg = "🤝 It's a Draw!"
		}

		h.renderGameBoardWithStatus(w, response.GameID, response.Board, player, winnerMsg, aiMoveMsg, true)
		return
	}

	// Return updated board
	h.renderBoardGrid(w, gameID, response.Board, player, aiMoveMsg)
}

// Helper functions to render HTML

func (h *WebHandler) renderGameBoard(w http.ResponseWriter, response *dto.CreateGameResponse, message string) {
	playerName := "X"
	if response.Player == 2 {
		playerName = "O"
	}

	aiMoveMsg := ""
	if response.AIMove != nil {
		aiMoveMsg = fmt.Sprintf("AI played: Row %d, Col %d", response.AIMove.Row, response.AIMove.Col)
	}

	// If AI moved, it's now player's turn. Otherwise, it's AI's turn (player 1 plays first)
	isPlayerTurn := true
	if response.Player == 2 && response.AIMove == nil {
		// Player is O (second), and AI hasn't moved yet, so it's AI's turn
		isPlayerTurn = false
	}

	fmt.Fprintf(w, `
	<div class="game-board active" 
		 x-data="{ gameOver: false, isMyTurn: %t, isProcessing: false }"
		 @htmx:after-swap.window="
			console.log('HTMX swap detected, updating state...');
			isMyTurn = true; 
			isProcessing = false;
			console.log('State after swap:', $data);
		 "
		 @game-over.window="
			console.log('Game over event received');
			gameOver = true;
			isMyTurn = false;
		 ">
		<div class="game-info" id="game-info">
			<h3>You are: %s</h3>
			<div class="status" x-show="!isMyTurn && !gameOver" style="color: #764ba2;">🤖 AI is thinking...</div>
			<div class="status" x-show="isMyTurn && !gameOver" style="color: #28a745;">✅ Your turn!</div>
			%s
		</div>
		<div id="board-container">
			%s
		</div>
		<div id="alpine-state-update"></div>
		<div id="winner-display"></div>
		<button class="new-game-btn" onclick="location.reload()">New Game</button>
	</div>
	`, isPlayerTurn, playerName,
		func() string {
			if aiMoveMsg != "" {
				return fmt.Sprintf(`<div class="ai-move">%s</div>`, aiMoveMsg)
			}
			return ""
		}(),
		h.getBoardHTML(response.GameID, response.Board, response.Player, false))
}

func (h *WebHandler) renderBoardGrid(w http.ResponseWriter, gameID string, board [][]int, player int, aiMoveMsg string) {
	playerName := "X"
	if player == 2 {
		playerName = "O"
	}

	// After AI move, it's player's turn again
	fmt.Fprintf(w, `
	<div id="board-container" hx-swap-oob="true">
		%s
	</div>
	<div class="game-info" id="game-info" hx-swap-oob="true">
		<h3>You are: %s</h3>
		<div class="status" style="color: #28a745;">✅ Your turn!</div>
		%s
	</div>
	<div id="alpine-state-update" hx-swap-oob="true">
		<!-- State update handled by @htmx:after-swap event on game-board -->
	</div>
	`, h.getBoardHTML(gameID, board, player, false),
		playerName,
		func() string {
			if aiMoveMsg != "" {
				return fmt.Sprintf(`<div class="ai-move">%s</div>`, aiMoveMsg)
			}
			return ""
		}())
}

func (h *WebHandler) renderGameBoardWithStatus(w http.ResponseWriter, gameID string, board [][]int, player int, statusMsg, aiMoveMsg string, gameOver bool) {
	playerName := "X"
	if player == 2 {
		playerName = "O"
	}

	// Return winner banner directly in HTML
	fmt.Fprintf(w, `
	<div id="board-container" hx-swap-oob="true">
		%s
	</div>
	<div class="game-info" id="game-info" hx-swap-oob="true">
		<h3>You are: %s</h3>
		%s
	</div>
	<div id="winner-display" hx-swap-oob="true">
		<div class="overlay"></div>
		<div class="winner-banner">
			<h2>%s</h2>
			<button class="new-game-btn" onclick="location.reload()">New Game</button>
		</div>
	</div>
	<div id="alpine-state-update" hx-swap-oob="true">
		<script>
			// Mark game as over - dispatch event to trigger Alpine handler
			document.dispatchEvent(new CustomEvent('game-over'));
		</script>
	</div>
	`, h.getBoardHTML(gameID, board, player, true),
		playerName,
		func() string {
			if aiMoveMsg != "" {
				return fmt.Sprintf(`<div class="ai-move">%s</div>`, aiMoveMsg)
			}
			return ""
		}(),
		statusMsg)
}

func (h *WebHandler) getBoardHTML(gameID string, board [][]int, player int, isGameOver bool) string {
	size := len(board)
	cellSize := 400 / size // Max 400px board width
	if cellSize > 60 {
		cellSize = 60
	}

	html := fmt.Sprintf(`<div class="board-grid" 
		x-init="console.log('Board grid init, isMyTurn:', isMyTurn, 'isProcessing:', isProcessing)"
		style="grid-template-columns: repeat(%d, %dpx); width: fit-content; margin: 0 auto;">`, size, cellSize)

	playerSymbol := "X"
	if player == 2 {
		playerSymbol = "O"
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			value := board[i][j]
			cellClass := ""
			cellContent := ""
			occupied := ""

			if value == 1 {
				cellClass = "x"
				cellContent = "X"
				occupied = "occupied"
			} else if value == 2 {
				cellClass = "o"
				cellContent = "O"
				occupied = "occupied"
			}

			if isGameOver {
				occupied = "occupied"
			}

			// Row and col are 1-indexed for the API
			// Simple solution: disable all cells when not player's turn
			baseStyle := fmt.Sprintf("font-size: %dpx;", cellSize/2)

			if occupied != "" {
				// Already occupied - always disabled
				html += fmt.Sprintf(`
				<div class="cell %s occupied" 
					 style="%s pointer-events: none;">
					%s
				</div>
			`, cellClass, baseStyle, cellContent)
			} else if !isGameOver {
				// Empty cell - use Alpine to control enable/disable via class
				// Determine which class to add (x or o) based on player symbol
				pendingClass := "x"
				if player == 2 {
					pendingClass = "o"
				}

				html += fmt.Sprintf(`
				<div class="cell %s" 
					 hx-post="/web/move" 
					 hx-vals='{"game_id": "%s", "row": %d, "col": %d}'
					 hx-target="#board-container"
					 hx-swap="outerHTML"
					 :class="{ 'cell-disabled': !isMyTurn || isProcessing }"
					 @click="if(isMyTurn && !isProcessing) { isMyTurn = false; isProcessing = true; $el.innerHTML = '%s'; $el.classList.add('pending', '%s'); }"
					 style="%s">
					%s
				</div>
			`, cellClass, gameID, i+1, j+1, playerSymbol, pendingClass, baseStyle, cellContent)
			} else {
				// Game over - disabled
				html += fmt.Sprintf(`
				<div class="cell %s occupied" 
					 style="%s pointer-events: none;">
					%s
				</div>
			`, cellClass, baseStyle, cellContent)
			}
		}
	}

	html += "</div>"
	return html
}

// GetGameState returns current game state as JSON (for debugging)
func (h *WebHandler) GetGameState(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")
	if gameID == "" {
		http.Error(w, "game_id required", http.StatusBadRequest)
		return
	}

	// This would need to be implemented in the service
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"game_id": gameID,
		"status":  "active",
	})
}
