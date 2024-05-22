package game2

import "slices"

var (
	RIGHT = [...]Move{MOVE_UP, MOVE_UP_RIGHT, MOVE_RIGHT, MOVE_DOWN_RIGHT}
	LEFT  = [...]Move{MOVE_DOWN, MOVE_DOWN_LEFT, MOVE_LEFT, MOVE_UP_LEFT}
)

func (t *TicTacToe) GetOpenThree(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{0, 0, 1, 1, 1, 0, 0}, []int{1, 5, 0, 6}, move)...)
	return threats
}

func (t *TicTacToe) GetBlockOpenThree(move Move) []Threat {
	threats := []Threat{}

	threats = append(threats, t.GetThreatMatching([]int{-1, 0, 1, 1, 1, 0, 0}, []int{1, 5, 6}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{0, 0, 1, 1, 1, 0, -1}, []int{5, 1, 0}, move)...)

	return threats
}

func (t *TicTacToe) Get2BlockOpenThree(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{-1, 0, 1, 1, 1, 0, -1}, []int{1, 5}, move)...)
	return threats
}

func (t *TicTacToe) GetOpenTwo(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{0, 0, 1, 1, 0, 0}, []int{1, 4, 0, 5}, move)...)

	return threats
}

func (t *TicTacToe) GetBlockOpenTwo(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{-1, 0, 1, 1, 0, 0}, []int{1, 4, 5}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{0, 0, 1, 1, 0, -1}, []int{4, 0, 1}, move)...)

	return threats
}

func (t *TicTacToe) GetOpenFour(move Move) []Threat {
	player := t.GetPlayer()
	threats := []Threat{}

	// --xxx-- -> -xxxx--
	for i := range 4 {
		r := move
		l := move
		cnt := 0
		for t.GetValue(r) == player {
			cnt++
			r1, err := r.AddMove(RIGHT[i])
			if err != nil {
				break
			}
			r = r1
		}
		for t.GetValue(l) == player {
			cnt++
			l1, err := l.AddMove(LEFT[i])
			if err != nil {
				break
			}
			l = l1
		}
		if t.GetValue(r) == EMPTY && t.GetValue(l) == EMPTY && cnt == 5 {
			threats = append(threats, Threat{[]Move{r, l}})
		}
	}

	return threats
}

func (t *TicTacToe) GetCloseThree(move Move) []Threat {
	// player := t.GetPlayer()
	// // opponent := GetOpponent(player)
	threats := []Threat{}

	threats = append(threats, t.GetThreatMatching([]int{-1, 1, 1, 1, 0, 0}, []int{4, 5}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{0, 0, 1, 1, 1, -1}, []int{0, 1}, move)...)

	return threats
}

func (t *TicTacToe) GetCloseFour(move Move) []Threat {
	player := t.GetPlayer()
	// opponent := GetOpponent(player)
	threats := []Threat{}

	// -oxxx-- -> -oxxxx-
	for i := range 4 {
		r := move
		l := move
		cnt := 0
		for t.GetValue(r) == player {
			cnt++
			r1, err := r.AddMove(RIGHT[i])
			if err != nil {
				break
			}
			r = r1
		}
		for t.GetValue(l) == player {
			cnt++
			l1, err := l.AddMove(LEFT[i])
			if err != nil {
				break
			}
			l = l1
		}

		// oxxxx-  ||  -xxxxo
		if t.GetValue(r) == EMPTY && t.GetValue(l) != EMPTY && cnt == 5 {
			threats = append(threats, Threat{[]Move{r}})
		}
		if t.GetValue(r) != EMPTY && t.GetValue(l) == EMPTY && cnt == 5 {
			threats = append(threats, Threat{[]Move{l}})
		}
	}

	return threats
}

func (t *TicTacToe) GetBrokenThree(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{0, 1, 1, 0, 1, 0}, []int{0, 3, 5}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{0, 1, 0, 1, 1, 0}, []int{0, 2, 5}, move)...)

	return threats
}

func (t *TicTacToe) GetBrokenTwo(move Move) []Threat {
	player := t.GetPlayer()
	threats := []Threat{}

	// -xx-- -> -xx-x-
	for i := range 4 {
		// -x-x- -> (left)x-x-
		left, err := move.AddMove(LEFT[i])
		if err != nil {
			// print("error: ", err, "\n")
			continue
		}
		if t.GetValue(left) != EMPTY {
			continue
		}
		// -x-x- -> (left)x(right)x-
		right, err := move.AddMove(RIGHT[i])
		if err != nil {
			// print("error: ", err, "\n")

			continue
		}
		if t.GetValue(right) != EMPTY {
			continue
		}
		// -x-x- -> (left)x(right)(right2)-
		right2, err := right.AddMove(RIGHT[i])
		if err != nil {
			// print("error: ", err, "\n")

			continue
		}
		if t.GetValue(right2) != player {
			continue
		}
		// -x-x- -> (left)x(right)(right2)(right3)
		right3, err := right2.AddMove(RIGHT[i])
		if err != nil {
			// print("error: ", err, "\n")

			continue
		}
		if t.GetValue(right3) != EMPTY {
			continue
		}
		threats = append(threats, Threat{[]Move{left, right, right3}})
	}

	for i := range 4 {
		left, err := move.AddMove(LEFT[i])
		if err != nil {
			// print("error: ", err, "\n")
			continue
		}
		if t.GetValue(left) != EMPTY {
			continue
		}
		// -x-x- -> (left)x(right)x-
		right, err := move.AddMove(RIGHT[i])
		if err != nil {
			// print("error: ", err, "\n")

			continue
		}
		if t.GetValue(right) != EMPTY {
			continue
		}
		left2, err := left.AddMove(LEFT[i])
		if err != nil {
			// print("error: ", err, "\n")

			continue
		}
		if t.GetValue(left2) != player {
			continue
		}
		left3, err := left2.AddMove(LEFT[i])
		if err != nil {
			// print("error: ", err, "\n")

			continue
		}
		if t.GetValue(left3) != EMPTY {
			continue
		}
		threats = append(threats, Threat{[]Move{left, left3, right}})
	}

	return threats
}

func (t *TicTacToe) GetBrokenCloseThree(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{-1, 1, 1, 0, 1, 0}, []int{3, 5}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{-1, 1, 0, 1, 1, 0}, []int{2, 5}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{0, 1, 0, 1, 1, -1}, []int{0, 2}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{0, 1, 1, 0, 1, -1}, []int{0, 3}, move)...)

	return threats
}

func (t *TicTacToe) GetBrokenCloseFour(move Move) []Threat {
	threats := []Threat{}
	threats = append(threats, t.GetThreatMatching([]int{-1, 1, 1, 1, 0, 1}, []int{4}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{-1, 1, 1, 0, 1, 1}, []int{3}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{-1, 1, 0, 1, 1, 1}, []int{2}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{1, 0, 1, 1, 1, -1}, []int{1}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{1, 1, 0, 1, 1, -1}, []int{2}, move)...)
	threats = append(threats, t.GetThreatMatching([]int{1, 1, 1, 0, 1, -1}, []int{3}, move)...)

	return threats
}

func (t *TicTacToe) GetFive(move Move) []Threat {
	player := t.GetPlayer()
	threats := []Threat{}

	// -oxxx-- -> -oxxxx-
	for i := range 4 {
		r := move
		l := move
		cnt := 0
		for t.GetValue(r) == player {
			cnt++
			r1, err := r.AddMove(RIGHT[i])
			if err != nil {
				break
			}
			r = r1
		}
		for t.GetValue(l) == player {
			cnt++
			l1, err := l.AddMove(LEFT[i])
			if err != nil {
				break
			}
			l = l1
		}

		if cnt >= 6 {
			threats = append(threats, Threat{[]Move{r}})
			break
		}
	}

	return threats
}

func (t *TicTacToe) GetThreatMatching(pattern, defendPattern []int, move Move) []Threat {
	threats := []Threat{}

	// 1. convert pattern to player
	player := t.GetPlayer()
	for i := range pattern {
		if pattern[i] == 1 {
			pattern[i] = player
		} else if pattern[i] == -1 {
			pattern[i] = GetOpponent(player)
		}
	}

	for i := range 4 {
		// 2. find all possible moves
		patternLength := len(pattern)
		moves := []Move{}
		leftMoves := []Move{}
		right := move
		left := move

		// 2.1. find all possible moves to the left
		for j := 0; j < patternLength-1; j++ {
			left1, err := left.AddMove(LEFT[i])
			if err != nil {
				break
			}
			left = left1
			leftMoves = append(leftMoves, left)
		}

		for j := len(leftMoves) - 1; j >= 0; j-- {
			moves = append(moves, leftMoves[j])
		}
		moves = append(moves, move)

		// 2.2. find all possible moves to the right
		for j := 0; j < patternLength-1; j++ {
			right1, err := right.AddMove(RIGHT[i])
			if err != nil {
				break
			}
			right = right1
			moves = append(moves, right)
		}

		// 3. check if the pattern matches
		for j := range len(moves) {
			it := j
			match := true
			for k := range pattern {
				if it >= len(moves) {
					break
				}
				value, err := t.GetValueError(moves[it])
				if err != nil {
					break
				}
				if pattern[k] != value {
					match = false
					break
				}
				it++
			}
			if match && it-j == patternLength {
				defendMoves := []Move{}
				for k := range defendPattern {
					defendMoves = append(defendMoves, Move{
						moves[j].Row + RIGHT[i].Row*defendPattern[k],
						moves[j].Col + RIGHT[i].Col*defendPattern[k]})
				}
				threats = append(threats, Threat{defendMoves})
			}
		}
	}

	return threats
}

// TODO: decrease game state when place a move
// TODO: find threats in boundary
func (t *TicTacToe) GetBlockThreats(move Move) {
	// --ooo-- -> -xooo- = openThree -> closeThree
	player := t.GetPlayer()
	state := &t.OState
	if player != PLAYER_X {
		state = &t.XState
	}

	// 1. Find all threats that contain the move
	// threats = append(threats, t.GetThreatMatching([]int{0, 0, 1, 1, 1, 0, 0}, []int{1, 5, 0, 6}, move)...)
	for i := 0; i < len(state.OpenThree); i++ {
		// 1.1. If the move is in the threat
		for j := 0; i < len(state.OpenThree) && j < len(state.OpenThree[i].DefendMove); j++ {
			defMove := state.OpenThree[i].DefendMove[j]
			if defMove.Row == move.Row && defMove.Col == move.Col && j == 0 {
				// 1.1.1. Remove the threat from the state
				newMove := []Move{state.OpenThree[i].DefendMove[1], state.OpenThree[i].DefendMove[3]}
				// newMove := []Move{{Row: 0, Col: 0}, {Row: 1, Col: 1}}
				state.OpenThree = removeFromSlice(state.OpenThree, i)
				// 1.1.2. Add new threat to the state
				state.CloseThree = append(state.CloseThree, Threat{newMove})
				continue
			}

			if defMove.Row == move.Row && defMove.Col == move.Col && j == 1 {
				// 1.1.1. Remove the threat from the state
				newMove := []Move{state.OpenThree[i].DefendMove[2], state.OpenThree[i].DefendMove[0]}
				// newMove := []Move{{Row: 0, Col: 0}, {Row: 1, Col: 1}}

				state.OpenThree = removeFromSlice(state.OpenThree, i)
				// 1.1.2. Add new threat to the state
				state.CloseThree = append(state.CloseThree, Threat{newMove})
				continue
			}

			if defMove.Row == move.Row && defMove.Col == move.Col && slices.Contains([]int{2, 3}, j) {
				// 1.1.1. Remove the threat from the state
				newMove := removeFromSlice(state.OpenThree[i].DefendMove, j)
				state.OpenThree = removeFromSlice(state.OpenThree, i)
				// 1.1.2. Add new threat to the state
				state.BlockOpenThree = append(state.BlockOpenThree, Threat{newMove})
			}
		}
	}

	for i := 0; i < len(state.BlockOpenThree); i++ {
		for j := 0; i < len(state.BlockOpenThree) && j < len(state.BlockOpenThree[i].DefendMove); j++ {
			defMove := state.BlockOpenThree[i].DefendMove[j]
			if defMove.Row == move.Row && defMove.Col == move.Col && j == 0 {
				newMove := removeFromSlice(state.BlockOpenThree[i].DefendMove, j)
				state.BlockOpenThree = removeFromSlice(state.BlockOpenThree, i)
				state.CloseThree = append(state.CloseThree, Threat{newMove})
			}
			if defMove.Row == move.Row && defMove.Col == move.Col && j == 1 {
				state.BlockOpenThree = removeFromSlice(state.BlockOpenThree, i)
			}
			if defMove.Row == move.Row && defMove.Col == move.Col && j == 2 {
				newMove := removeFromSlice(state.BlockOpenThree[i].DefendMove, j)
				state.BlockOpenThree = removeFromSlice(state.BlockOpenThree, i)
				state.TwoBlockOpenThree = append(state.TwoBlockOpenThree, Threat{newMove})
			}
		}
	}

	for i := 0; i < len(state.OpenFour); i++ {
		// 1.1. If the move is in the threat
		for j := 0; i < len(state.OpenFour) && j < len(state.OpenFour[i].DefendMove); j++ {
			defMove := state.OpenFour[i].DefendMove[j]
			if defMove.Row == move.Row && defMove.Col == move.Col {
				// 1.1.1. Remove the threat from the state
				newMove := removeFromSlice(state.OpenFour[i].DefendMove, j)
				state.OpenFour = removeFromSlice(state.OpenFour, i)
				// 1.1.2. Add new threat to the state
				state.CloseFour = append(state.CloseFour, Threat{newMove})
			}
		}
	}

	// []int{0, 0, 1, 1, 0, 0}, []int{1, 4, 0, 5}
	for i := 0; i < len(state.BrokenTwo); i++ {
		// 1.1. If the move is in the threat
		for j := 0; i < len(state.BrokenTwo) && j < len(state.OpenTwo[i].DefendMove); j++ {
			defMove := state.OpenTwo[i].DefendMove[j]
			if defMove.Row == move.Row && defMove.Col == move.Col && slices.Contains([]int{0, 1}, j) {
				// 1.1.1. Remove the threat from the state
				state.OpenTwo = removeFromSlice(state.OpenTwo, i)
				continue
			}
			if defMove.Row == move.Row && defMove.Col == move.Col && slices.Contains([]int{2, 3}, j) {
				// 1.1.1. Remove the threat from the state
				newMove := []Move{}
				if len(state.OpenTwo[i].DefendMove) > 0 {
					newMove = removeFromSlice(state.OpenTwo[i].DefendMove, j)
				}
				state.OpenTwo = removeFromSlice(state.OpenTwo, i)
				// 1.1.2. Add new threat to the state
				state.BlockOpenTwo = append(state.BlockOpenTwo, Threat{newMove})
			}
		}
	}

	// threats = append(threats, t.GetThreatMatching([]int{0, 1, 1, 0, 1, 0}, []int{0, 3, 5}, move)...)
	// threats = append(threats, t.GetThreatMatching([]int{0, 1, 0, 1, 1, 0}, []int{0, 2, 5}, move)...)
	for i := 0; i < len(state.BrokenThree); i++ {
		for j := 0; i < len(state.BrokenThree) && j < len(state.BrokenThree[i].DefendMove); j++ {
			defMove := state.BrokenThree[i].DefendMove[j]
			if defMove.Row == move.Row && defMove.Col == move.Col && slices.Contains([]int{0, 2}, j) {
				newMove := removeFromSlice(state.BrokenThree[i].DefendMove, j)
				state.BrokenThree = removeFromSlice(state.BrokenThree, i)
				state.BrokenCloseThree = append(state.BrokenCloseThree, Threat{newMove})
			}
			if defMove.Row == move.Row && defMove.Col == move.Col && j == 1 {
				state.BrokenThree = removeFromSlice(state.BrokenThree, i)
			}
		}
	}

	// BrokenTwo

	removeThreats := []*[]Threat{&state.BrokenCloseFour, &state.CloseFour, &state.CloseThree,
		&state.BrokenCloseThree, &state.BrokenTwo, &state.BlockOpenTwo, &state.TwoBlockOpenThree}

	for k := 0; k < len(removeThreats); k++ {
		for i := 0; i < len(*removeThreats[k]); i++ {
			// 1.1. If the move is in the threat
			for j := 0; i < len(*removeThreats[k]) && j < len((*removeThreats[k])[i].DefendMove); j++ {
				defMove := (*removeThreats[k])[i].DefendMove[j]
				if defMove.Row == move.Row && defMove.Col == move.Col {
					// 1.2. Remove the threat from the state
					*removeThreats[k] = removeFromSlice((*removeThreats[k]), i)
				}
			}
		}
	}
}

// TODO: decrease game state when place a move
func (t *TicTacToe) DecreaseThreats(move Move) {
	player := t.GetPlayer()
	state := &t.OState
	if player == PLAYER_X {
		state = &t.XState
	}

	threats := []*[]Threat{
		&state.OpenTwo,
		&state.BlockOpenTwo,
		&state.BrokenTwo,
		&state.OpenThree,
		&state.BlockOpenThree,
		&state.TwoBlockOpenThree,
		&state.BrokenThree,
		&state.BrokenCloseThree,
		&state.CloseThree,
		&state.OpenFour,
		&state.BrokenCloseFour,
		&state.CloseFour,
		&state.Five,
	}

	for k := range threats {
		for i := 0; i < len(*threats[k]); i++ {
			for _, defMove := range (*threats[k])[i].DefendMove {
				if defMove.Row == move.Row && defMove.Col == move.Col {
					*threats[k] = removeFromSlice((*threats[k]), i)
					break
				}
			}
		}
	}
}

// helper
func removeFromSlice[T Threat | Move](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}
