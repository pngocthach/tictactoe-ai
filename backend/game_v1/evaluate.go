package game

type Score struct {
	X int
	O int
}

// ham danh gia
func (t *TicTacToe) Evaluate() float64 {
	scoreX := 0.0
	scoreO := 0.0

	vectors := t.GetAllTrimVectors()
	scoreX = CalculatePoint(vectors, scoreX, PLAYER_X)
	scoreO = CalculatePoint(vectors, scoreO, PLAYER_O)
	// return scoreX / scoreO
	return scoreX - scoreO*0.8
}

func CalculatePoint(vectors [][]int, score float64, player int) float64 {
	closeFour := 0
	openThree := 0
	// openFour := 0
	for i := 0; i < len(vectors); i++ {
		score += float64(VectorPatternMatchCount(&vectors[i], &[]int{1, 1, 1, 1, 1}, player)) * 9000000
		score += float64(VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 1, 1, 1, 0}, player)) * 10000
		score += float64(VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 1, 0}, player)) * 10
		score += float64(VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 0, 1, 0}, player)) * 10

		closeFour += VectorPatternMatchCount(&vectors[i], &[]int{-1, 1, 1, 1, 0, 1}, player)
		closeFour += VectorPatternMatchCount(&vectors[i], &[]int{1, 0, 1, 1, 1, -1}, player)
		closeFour += VectorPatternMatchCount(&vectors[i], &[]int{1, 1, 0, 1, 1}, player)
		closeFour += VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 1, 1, 1, -1}, player)
		closeFour += VectorPatternMatchCount(&vectors[i], &[]int{-1, 1, 1, 1, 1, 0}, player)

		openThree += VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 1, 1, 0}, player)
		openThree += VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 1, 0, 1, 0}, player)
		openThree += VectorPatternMatchCount(&vectors[i], &[]int{0, 1, 0, 1, 1, 0}, player)
		// scoreX += float64(VectorPatternMatchCount(&vectors[i], []int{-1, 1, 1, 1, 0}, player)) * 20
		// scoreX += float64(VectorPatternMatchCount(&vectors[i], []int{0, 1, 1, 1, -1}, player)) * 20

		// scoreX += float64(VectorPatternMatchCount(&vectors[i], []int{0, 1, 1, 0}, player)) * 5
		// scoreX += float64(VectorPatternMatchCount(&vectors[i], []int{0, 1, 1, -1}, player))
		// scoreX += float64(VectorPatternMatchCount(&vectors[i], []int{-1, 1, 1, 0}, player))
	}

	if openThree >= 2 {
		score += 9000
	}
	if closeFour >= 2 {
		score += 20000
	}
	if closeFour >= 1 && openThree >= 1 {
		score += 15000
	}
	score += float64(closeFour)*1000 + float64(openThree)*900
	return score
}

func VectorPatternMatchCount(vector *[]int, pattern *[]int, player int) int {

	if len(*vector) < len(*pattern) {
		return 0
	}
	count := 0
	for i := range len(*vector) - len(*pattern) + 1 {
		var j int
		for j = 0; j < len(*pattern); j++ {
			if (*pattern)[j] == 1 && (*vector)[i+j] != player {
				break
			}
			if (*pattern)[j] == -1 && ((*vector)[i+j] != -player || (*vector)[i+j] != W) {
				break
			}
			if (*pattern)[j] == 0 && (*vector)[i+j] != 0 {
				break
			}
		}
		if j == len(*pattern) {
			count++
		}
	}
	return count
}

func (t *TicTacToe) GetAllTrimVectors() [][]int {
	vectors := make([][]int, (t.boardSize + 2))
	for i := 0; i < t.boardSize+2; i++ {
		vectors[i] = make([]int, (t.boardSize + 2))
	}

	rowVectors := make([][]int, (t.boardSize + 2))
	for i := 1; i < t.boardSize+1; i++ {
		rowVectors[i] = t.Board[i]
		// trim head and tail of vector if dist > MAX_DIST and isEmpty
		head := 0
		tail := len(rowVectors[i]) - 1
		for head < len(rowVectors[i]) && rowVectors[i][head] == 0 && t.Dist[i][head] > MAX_DIST {
			head++
		}
		if head == len(rowVectors[i]) {
			rowVectors[i] = []int{}
			continue
		}
		for tail >= 0 && rowVectors[i][tail] == 0 && t.Dist[i][tail] > MAX_DIST {
			tail--
		}
		if head < 3 {
			head = 0
		}
		if tail > len(rowVectors[i])-4 {
			tail = len(rowVectors[i]) - 1
		}
		rowVectors[i] = rowVectors[i][head : tail+1]
	}

	columnVectors := make([][]int, (t.boardSize + 2))
	for i := 1; i < t.boardSize+1; i++ {
		for j := 1; j < t.boardSize+1; j++ {
			columnVectors[j] = append(columnVectors[j], t.Board[i][j])
		}
	}

	vectors = append(vectors, rowVectors...)
	vectors = append(vectors, columnVectors...)
	vectors = append(vectors, getAllDiag(&t.Board)...)
	vectors = append(vectors, getAllAntiDiag(&t.Board)...)

	return vectors
}

func getDiag(a *[][]int, i, j int) []int {
	n := len(*a)
	row := i
	column := j
	res := []int{}
	for row < n && column < n {
		// print(" ", a[row][column])
		res = append(res, (*a)[row][column])
		row++
		column++
	}
	return res
}

func getAllDiag(a *[][]int) [][]int {
	n := len(*a)
	res := [][]int{}
	for i := 0; i < n; i++ {
		diag := getDiag(a, 0, i)
		if len(diag) >= 5 {
			res = append(res, diag)
		}
	}
	for i := 1; i < n; i++ {
		diag := getDiag(a, i, 0)
		if len(diag) >= 5 {
			res = append(res, diag)
		}
	}
	return res
}

func getAntiDiag(a *[][]int, i, j int) []int {
	n := len(*a)
	row := i
	column := j
	res := []int{}
	for row < n && column >= 0 {
		// print(" ", a[row][column])
		res = append(res, (*a)[row][column])
		row++
		column--
	}
	return res
}

func getAllAntiDiag(a *[][]int) [][]int {
	n := len(*a)
	res := [][]int{}
	for i := 0; i < n; i++ {
		diag := getAntiDiag(a, 0, i)
		if len(diag) >= 5 {
			res = append(res, diag)
		}
	}
	for i := 1; i < n; i++ {
		diag := getAntiDiag(a, i, n-1)
		if len(diag) >= 5 {
			res = append(res, diag)
		}
	}
	return res
}
