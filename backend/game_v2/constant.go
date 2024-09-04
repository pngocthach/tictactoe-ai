package game2

type Pattern struct {
	pattern []int
	score   int
}

var (
	PLAYER_X   = 1
	PLAYER_O   = -1
	X          = PLAYER_X
	O          = PLAYER_O
	EMPTY      = 0
	e          = EMPTY
	WIN_SIZE   = 5
	MAX_DEPTH  = 7
	BOARD_SIZE = 10

	X_6PATTERNS = [...][]int{
		{e, X, X, X, X, e},
		{e, X, X, X, e, e},
		{e, e, X, X, X, e},
		{e, X, X, e, X, e},
		{e, X, e, X, X, e},
		{e, e, X, X, e, e},
		{e, e, X, e, X, e},
		{e, X, e, X, e, e},
		{e, e, X, e, e, e},
		{e, e, e, X, e, e}}

	X_6SCORES = [...]int{50000, 5000, 5000, 500, 500, 100, 100, 100, 10, 10}

	X_5PATTERNS = []Pattern{
		{[]int{X, X, X, X, X}, 64},

		// 4 X and 1 empty
		{[]int{X, X, X, X, e}, 32},
		{[]int{e, X, X, X, X}, 32},
		{[]int{X, X, e, X, X}, 32},
		{[]int{X, e, X, X, X}, 32},
		{[]int{X, X, X, e, X}, 32},

		// 3 X and 2 empty
		{[]int{X, X, X, e, e}, 16},
		{[]int{e, e, X, X, X}, 16},
		{[]int{e, X, X, X, e}, 16},

		// 3 X, 1 empty, 1 X
		{[]int{X, X, X, e, X}, 8},
		{[]int{X, e, X, X, X}, 8},
		{[]int{X, X, X, X, e}, 8},

		{[]int{e, X, X, X, X}, 4},
		{[]int{X, X, X, X, e}, 4},
		{[]int{X, X, X, X, e}, 4},
		{[]int{e, X, X, X, X}, 4},
		{[]int{X, X, e, X, X}, 4},
		{[]int{X, e, X, X, X}, 4},
		{[]int{e, X, X, X, X}, 4},
		{[]int{e, X, X, X, X}, 4},

		// 2 X and 3 empty
		{[]int{X, X, e, e, e}, 2},
		{[]int{e, X, X, e, e}, 2},
		{[]int{e, e, X, X, e}, 2},
		{[]int{e, e, e, X, X}, 2},
		{[]int{X, e, X, e, e}, 2},
		{[]int{e, X, e, X, e}, 2},
		{[]int{e, e, X, e, X}, 2},

		// 2 X, 2 empty, 1 X
		{[]int{X, X, X, e, e}, 1},
		{[]int{e, X, X, X, e}, 1},
		{[]int{e, e, X, X, X}, 1},
		{[]int{e, e, X, X, X}, 1},
		{[]int{e, X, X, X, e}, 1},
		{[]int{X, X, X, e, e}, 1},
	}

	// X_5SCORES = [...]int{1000000, 5000, 5000, 5000, 5000, 5000}
	X_5SCORES = [...]int{64, 32, 32, 32, 32, 32}

	O_6PATTERNS = [...][]int{
		{e, O, O, O, O, e},
		{e, O, O, O, e, e},
		{e, e, O, O, O, e},
		{e, O, O, e, O, e},
		{e, O, e, O, O, e},
		{e, e, O, O, e, e},
		{e, e, O, e, O, e},
		{e, O, e, O, e, e},
		{e, e, O, e, e, e},
		{e, e, e, O, e, e}}

	O_6SCORES = [...]int{50000, 5000, 5000, 500, 500, 100, 100, 100, 10, 10}

	O_5PATTERNS = []Pattern{
		{[]int{O, O, O, O, O}, 64},

		// 4 O and 1 empty
		{[]int{O, O, O, O, e}, 32},
		{[]int{e, O, O, O, O}, 32},
		{[]int{O, O, e, O, O}, 32},
		{[]int{O, e, O, O, O}, 32},
		{[]int{O, O, O, e, O}, 32},

		// 3 O and 2 empty
		{[]int{O, O, O, e, e}, 16},
		{[]int{e, e, O, O, O}, 16},
		{[]int{e, O, O, O, e}, 16},

		// 3 O, 1 empty, 1 X
		{[]int{O, O, O, e, X}, 8},
		{[]int{X, e, O, O, O}, 8},
		{[]int{O, O, O, X, e}, 8},

		{[]int{e, X, O, O, O}, 4},
		{[]int{X, O, O, O, e}, 4},
		{[]int{O, O, O, X, e}, 4},
		{[]int{e, O, O, O, X}, 4},
		{[]int{O, O, e, O, X}, 4},
		{[]int{O, e, O, O, X}, 4},
		{[]int{e, O, X, O, O}, 4},
		{[]int{e, O, O, X, O}, 4},

		// 2 O and 3 empty
		{[]int{O, O, e, e, e}, 2},
		{[]int{e, O, O, e, e}, 2},
		{[]int{e, e, O, O, e}, 2},
		{[]int{e, e, e, O, O}, 2},
		{[]int{O, e, O, e, e}, 2},
		{[]int{e, O, e, O, e}, 2},
		{[]int{e, e, O, e, O}, 2},

		// 2 O, 2 empty, 1 X
		{[]int{O, O, X, e, e}, 1},
		{[]int{e, O, O, X, e}, 1},
		{[]int{e, e, O, O, X}, 1},
		{[]int{e, e, X, O, O}, 1},
		{[]int{e, X, O, O, e}, 1},
		{[]int{X, O, O, e, e}, 1},
	}

	// O_5SCORES = [...]int{1000000, 5000, 5000, 5000, 5000, 5000}
	O_5SCORES = [...]int{64, 32, 32, 32, 32, 32, 16, 16, 16, 16}

	MOVE_UP         = Move{-1, 0}
	MOVE_DOWN       = Move{1, 0}
	MOVE_LEFT       = Move{0, -1}
	MOVE_RIGHT      = Move{0, 1}
	MOVE_UP_RIGHT   = Move{-1, 1}
	MOVE_DOWN_LEFT  = Move{1, -1}
	MOVE_UP_LEFT    = Move{-1, -1}
	MOVE_DOWN_RIGHT = Move{1, 1}
)
