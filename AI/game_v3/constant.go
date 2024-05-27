package game3

import "time"

const (
	PLAYER_X = 0b01
	PLAYER_O = 0b10
	EMPTY    = 0b00
	WALL     = 0b11
	x        = PLAYER_X
	o        = PLAYER_O
	e        = EMPTY
	w        = WALL
)

var (
	AI_PLAYER                     = PLAYER_X
	EVAL_PARAM                    = 1.0
	DIST                          = 1
	WIN_SIZE                      = 5
	MAX_DEPTH                     = 3
	MAX_DIST                      = 1
	MAX_TIME        time.Duration = 60
	MOVE_UP                       = Move{-1, 0}
	MOVE_DOWN                     = Move{1, 0}
	MOVE_LEFT                     = Move{0, -1}
	MOVE_RIGHT                    = Move{0, 1}
	MOVE_UP_RIGHT                 = Move{-1, 1}
	MOVE_DOWN_LEFT                = Move{1, -1}
	MOVE_UP_LEFT                  = Move{-1, -1}
	MOVE_DOWN_RIGHT               = Move{1, 1}
	DIRECTION                     = [8]Move{MOVE_UP, MOVE_DOWN, MOVE_LEFT, MOVE_RIGHT, MOVE_UP_RIGHT, MOVE_DOWN_LEFT, MOVE_UP_LEFT, MOVE_DOWN_RIGHT}
	PATTERN                       = map[string][]Pattern{}
	RIGHT                         = [...]Move{MOVE_UP, MOVE_UP_RIGHT, MOVE_RIGHT, MOVE_DOWN_RIGHT}
	LEFT                          = [...]Move{MOVE_DOWN, MOVE_DOWN_LEFT, MOVE_LEFT, MOVE_UP_LEFT}
)

// type Pattern struct {
// 	Player  int
// 	Pattern string
// }
