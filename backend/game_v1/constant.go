package game

import "time"

type Pattern struct {
	pattern []int
	score   int
}

var (
	PLAYER_X                      = 1
	PLAYER_O                      = -1
	X                             = PLAYER_X
	O                             = PLAYER_O
	EMPTY                         = 0
	e                             = EMPTY
	WIN_SIZE                      = 5
	MAX_DEPTH                     = 7
	BOARD_SIZE                    = 10
	W                             = -2 // WALL
	MAX_DIST                      = 1
	DIST                          = 1
	MAX_TIME        time.Duration = 20
	MOVE_UP                       = Move{-1, 0}
	MOVE_DOWN                     = Move{1, 0}
	MOVE_LEFT                     = Move{0, -1}
	MOVE_RIGHT                    = Move{0, 1}
	MOVE_UP_RIGHT                 = Move{-1, 1}
	MOVE_DOWN_LEFT                = Move{1, -1}
	MOVE_UP_LEFT                  = Move{-1, -1}
	MOVE_DOWN_RIGHT               = Move{1, 1}
	DIRECTION                     = [8]Move{MOVE_UP, MOVE_DOWN, MOVE_LEFT, MOVE_RIGHT, MOVE_UP_RIGHT, MOVE_DOWN_LEFT, MOVE_UP_LEFT, MOVE_DOWN_RIGHT}
	RIGHT                         = [...]Move{MOVE_UP, MOVE_UP_RIGHT, MOVE_RIGHT, MOVE_DOWN_RIGHT}
	LEFT                          = [...]Move{MOVE_DOWN, MOVE_DOWN_LEFT, MOVE_LEFT, MOVE_UP_LEFT}
)
