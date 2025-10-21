# Tic-Tac-Toe AI Server

A Tic-Tac-Toe game server with AI opponent using alpha-beta pruning algorithm.
[Live demo](https://tictactoe-ai-kac2.onrender.com)
## Features

- RESTful API for game management
- AI opponent using alpha-beta pruning with threat pattern evaluation
- In-memory game storage
- Support for variable board sizes (5-20)
- Pure Go implementation (no external dependencies)

## Project Structure

```
AI/
├── internal/
│   ├── game/
│   │   └── v3/          # Game logic implementation
│   ├── handler/         # HTTP handlers
│   ├── service/         # Business logic
│   ├── repository/      # Data storage
│   ├── entity/          # Domain models
│   └── dto/             # Data transfer objects
├── build/               # Compiled binaries
├── main.go              # Server entry point
└── go.mod
```

## API Endpoints

### Create Game

**POST** `/api/games`

Request:

```json
{
  "board_size": 12,
  "player": 1
}
```

- `board_size`: Size of the board (5-20)
- `player`: Your player (1 for X, 2 for O)

Response (player goes first):

```json
{
  "game_id": "12345678-1234-1234-1234-123456789abc",
  "board_size": 12,
  "player": 1,
  "board": [[0, 0, ...], ...],
  "message": "Game created successfully"
}
```

Response (player goes second - AI makes first move):

```json
{
  "game_id": "12345678-1234-1234-1234-123456789abc",
  "board_size": 12,
  "player": 2,
  "ai_move": {
    "row": 7,
    "col": 7
  },
  "board": [[0, 0, ...], [0, 0, 1, ...], ...],
  "message": "Game created successfully"
}
```

### Make Move

**POST** `/api/games/move`

Request:

```json
{
  "game_id": "12345678-1234-1234-1234-123456789abc",
  "row": 6,
  "col": 6
}
```

- `game_id`: Game UUID from create game response
- `row`: Row position (1-based index)
- `col`: Column position (1-based index)

Response:

```json
{
  "game_id": "12345678-1234-1234-1234-123456789abc",
  "row": 6,
  "col": 6,
  "ai_move": {
    "row": 7,
    "col": 6
  },
  "board": [[0, 0, ...], ...],
  "game_over": false,
  "winner": null,
  "message": "Move processed successfully"
}
```

- `board`: 2D array representing the board state (0: empty, 1: X, 2: O, 3: wall)
- `game_over`: Boolean indicating if the game is finished
- `winner`: Player who won (1 or 2), null if game is ongoing or draw

### Health Check

**GET** `/health`

Response: `OK`

## Run

### Using Pre-built Binary

```bash
./build/ai_linux_amd64    # Linux
./build/ai_windows_amd64.exe    # Windows
```

### Build and Run

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### 🎮 Play in Browser

Open your browser and go to: **http://localhost:8080**

The web interface uses HTMX for interactive gameplay - no JavaScript coding needed!

Features:

- Beautiful gradient UI
- Choose board size and who goes first
- Real-time AI responses
- Visual feedback for moves
- Win/draw detection with animations

See [WEB_INTERFACE.md](WEB_INTERFACE.md) for details.

## Build

Build for your platform:

```bash
go build -o build/ai_server
```

Cross-compile:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o build/ai_linux_amd64

# Windows
GOOS=windows GOARCH=amd64 go build -o build/ai_windows_amd64.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o build/ai_darwin_amd64
```

## Example Usage

Create a game:

```bash
curl -X POST http://localhost:8080/api/games \
  -H "Content-Type: application/json" \
  -d '{"board_size": 12, "player": 1}'
```

Make a move:

```bash
curl -X POST http://localhost:8080/api/games/move \
  -H "Content-Type: application/json" \
  -d '{"game_id": "YOUR_GAME_ID", "row": 6, "col": 6}'
```

## AI Configuration

The AI uses the following parameters (configured in `internal/service/game_service.go`):

- `MAX_DEPTH`: 2 - Maximum search depth
- `MAX_TIME`: 2 seconds - Maximum time for AI to think
- `MAX_DIST`: 2 - Maximum distance for move consideration
- `DIST`: 2 - Distance parameter for neighbor calculation

## Game Rules

- Connect 5 pieces in a row (horizontal, vertical, or diagonal) to win
- Players alternate turns
- AI automatically makes its move after the player's move
- Board coordinates are 1-based (1 to board_size)
