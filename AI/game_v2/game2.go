package game2

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Node struct {
	Value     int
	MoveCount int
	Dist      int // khoang cach den o gan nhat ma khong trong
}

type Threat struct {
	DefendMove []Move
}

type State struct {
	OpenTwo           []Threat // -xx-
	BlockOpenTwo      []Threat // o-xx--
	BrokenTwo         []Threat // -x-x-
	OpenThree         []Threat // -xxx-
	BlockOpenThree    []Threat // o-xxx--
	TwoBlockOpenThree []Threat // o-xxx-o
	BrokenThree       []Threat // -xx-x-
	BrokenCloseThree  []Threat // -oxx-x-
	CloseThree        []Threat // -oxxx-
	OpenFour          []Threat // -xxxx-
	BrokenCloseFour   []Threat // -oxxx-x-
	CloseFour         []Threat // -oxxxx-
	Five              []Threat // xxxxx
}

// tao struct tic-tac-toe nxn
type TicTacToe struct {
	Board     [][]Node
	boardSize int
	winSize   int
	MoveCount int
	XState    State
	OState    State
}

func NewTicTacToe() *TicTacToe {
	board := make([][]Node, BOARD_SIZE)
	for i := range board {
		board[i] = make([]Node, BOARD_SIZE)
	}
	for i := range BOARD_SIZE {
		for j := range BOARD_SIZE {
			board[i][j] = Node{EMPTY, math.MaxInt32, math.MaxInt32}
		}
	}
	return &TicTacToe{board, BOARD_SIZE, WIN_SIZE, 0, State{}, State{}}
}

func (t *TicTacToe) CheckWin() bool {
	return len(t.OState.Five) > 0 || len(t.XState.Five) > 0
}

type Move struct {
	Row int
	Col int
}

func (t *TicTacToe) MakeMove(move Move) error {

	if t.Board[move.Row][move.Col].Value != EMPTY {
		// print error
		println("cell", move.Row, move.Col, "is not empty")
		return errors.New("cell is not empty")
	}

	// 2. update MoveCount -> MoveCount++
	t.MoveCount++

	player := t.GetPlayer()
	println("player", player, "move to", move.Row, move.Col)
	// 1. update Board
	// 1.1. update Node Value
	t.Board[move.Row][move.Col].Value = player
	// 1.2. update Node MoveCount
	t.Board[move.Row][move.Col].MoveCount = t.MoveCount
	// 1.3. update dist (neighbor Node)

	// 3. update XState, OState
	// 3.1 if MoveCount % 2 == 0 -> playerX

	state := &t.XState
	if player == PLAYER_O {
		state = &t.OState
	}
	// 3.2 XState.OpenThreat.append(GetOpenThreats(move, size))
	t.DecreaseThreats(move)
	state.OpenTwo = append(state.OpenTwo, t.GetOpenTwo(move)...)
	state.BlockOpenTwo = append(state.BlockOpenTwo, t.GetBlockOpenTwo(move)...)
	state.BrokenTwo = append(state.BrokenTwo, t.GetBrokenTwo(move)...)
	state.OpenThree = append(state.OpenThree, t.GetOpenThree(move)...)
	state.BlockOpenThree = append(state.BlockOpenThree, t.GetBlockOpenThree(move)...)
	state.TwoBlockOpenThree = append(state.TwoBlockOpenThree, t.Get2BlockOpenThree(move)...)
	state.BrokenThree = append(state.BrokenThree, t.GetBrokenThree(move)...)
	state.CloseThree = append(state.CloseThree, t.GetCloseThree(move)...)
	state.BrokenCloseThree = append(state.BrokenCloseThree, t.GetBrokenCloseThree(move)...)
	state.OpenFour = append(state.OpenFour, t.GetOpenFour(move)...)
	state.BrokenCloseFour = append(state.BrokenCloseFour, t.GetBrokenCloseFour(move)...)
	state.CloseFour = append(state.CloseFour, t.GetCloseFour(move)...)
	state.Five = append(state.Five, t.GetFive(move)...)
	t.GetBlockThreats(move)
	return nil
}

func (m Move) AddMove(move Move) (Move, error) {
	if m.Row+move.Row < 0 || m.Row+move.Row >= BOARD_SIZE || m.Col+move.Col < 0 || m.Col+move.Col >= BOARD_SIZE {
		return Move{}, errors.New("move" + strconv.Itoa(m.Row+move.Row) + " " + strconv.Itoa(m.Col+move.Col) + "out of bound")
	}
	return Move{m.Row + move.Row, m.Col + move.Col}, nil
}

func (t *TicTacToe) GetValue(move Move) int {
	return t.Board[move.Row][move.Col].Value
}

func (t *TicTacToe) GetValueError(move Move) (int, error) {
	if move.Row < 0 || move.Row >= t.boardSize || move.Col < 0 || move.Col >= t.boardSize {
		return -1, errors.New("move" + strconv.Itoa(move.Row+move.Row) + " " + strconv.Itoa(move.Col+move.Col) + "out of bound")
	}

	return t.Board[move.Row][move.Col].Value, nil
}

func GetOpponent(player int) int {
	if player == PLAYER_X {
		return PLAYER_O
	}
	return PLAYER_X
}

func (t *TicTacToe) GetPlayer() int {
	var player int
	if t.MoveCount%2 != 0 {
		player = PLAYER_X
	} else {
		player = PLAYER_O
	}
	return player
}

func (t *TicTacToe) GetEmptyCells() []Move {
	var emptyCells []Move
	for i := 0; i < t.boardSize; i++ {
		for j := 0; j < t.boardSize; j++ {
			if t.Board[i][j].Value == EMPTY {
				emptyCells = append(emptyCells, Move{i, j})
			}
		}
	}
	return emptyCells
}

func (t *TicTacToe) UndoMove(move Move) {
	// them previousState vao struct TicTacToe de giam thoi gian tinh toan

	// create slice of TicTacToe to save game state

	panic("not implemented")
}

func (t *TicTacToe) PrintBoard() {
	// print with column and row number
	// separate with |
	// print with player X and O
	// print with empty cell
	// print with move count

	print(" |")
	for i := 0; i < t.boardSize; i++ {
		print(i, "|")
	}
	println()

	for i := 0; i < t.boardSize; i++ {
		print(i, "|")
		for j := 0; j < t.boardSize; j++ {
			switch t.Board[i][j].Value {
			case PLAYER_X:
				print("X")
			case PLAYER_O:
				print("O")
			case EMPTY:
				print(" ")
			}
			print("|")
		}
		println()
	}
}

// player vs player
func (t *TicTacToe) PlayPvP() {
	player := PLAYER_X
	for !t.CheckWin() {
		println("\nO state")
		t.OState.PrintState()
		println("\nX state")
		t.XState.PrintState()
		println("eval: ", int(t.Evaluate()), " - ", t.Evaluate())
		println()
		t.PrintBoard()
		fmt.Println("Player", player)
		// get move from user input
		move := getUserInput()
		err := t.MakeMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		player = GetOpponent(player)
	}
	t.PrintBoard()

	if player == PLAYER_X {
		fmt.Println("Player O wins")
	} else {
		fmt.Println("Player X wins")
	}
}

func (s *State) PrintState() {
	fmt.Println("OpenTwo: ", len(s.OpenTwo), "-", s.OpenTwo)
	fmt.Println("BlockOpenTwo: ", len(s.BlockOpenTwo), "-", s.BlockOpenTwo)
	fmt.Println("BrokenTwo: ", len(s.BrokenTwo), "-", s.BrokenTwo)
	fmt.Println("OpenThree: ", len(s.OpenThree), "-", s.OpenThree)
	fmt.Println("BlockOpenThree: ", len(s.BlockOpenThree), "-", s.BlockOpenThree)
	fmt.Println("TwoBlockOpenThree: ", len(s.TwoBlockOpenThree), "-", s.TwoBlockOpenThree)
	fmt.Println("BrokenThree: ", len(s.BrokenThree), "-", s.BrokenThree)
	fmt.Println("CloseThree: ", len(s.CloseThree), "-", s.CloseThree)
	fmt.Println("BrokenCloseThree: ", len(s.BrokenCloseThree), "-", s.BrokenCloseThree)
	fmt.Println("OpenFour: ", len(s.OpenFour), "-", s.OpenFour)
	fmt.Println("BrokenCloseFour: ", len(s.BrokenCloseFour), "-", s.BrokenCloseFour)
	fmt.Println("CloseFour: ", len(s.CloseFour), "-", s.CloseFour)
	fmt.Println("Five: ", len(s.Five), "-", s.Five)
}

func getUserInput() Move {
	var row, col int
	fmt.Print("Enter row: ")
	fmt.Scan(&row)
	fmt.Print("Enter col: ")
	fmt.Scan(&col)
	return Move{row, col}
}
