## Run

- Run the file in the build directory or use the command `go run main.go`.

## Build

- Linux: `GOOS=linux GOARCH=amd64 go build -o build/ai_linux_amd64`
- Windows: `GOOS=windows GOARCH=amd64 go build -o build/ai_windows_amd64.exe`

## Http Server

### POST /init

- Init game
- Request:

```json
{
  "BoardSize": 15,
  "Player": "X" // AI Player X/O
}
```

- Response:

```json
{
  "Success": true,
  // first move of AI (if init Player = X)
  "FirstMove": {
    "Row": -1, // -1 if AI is not first
    "Col": -1
  }
}
```

### POST /move

- Make a move
- Request:

```json
{
  "Row": 6, // 1-based indexing [1 -> BoardSize]
  "Col": 5
}
```

- Response:

```json
{
  "Move": {
    "Row": 6,
    "Col": 4
  },
  "GameOver": true
}
```
