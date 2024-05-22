package game3

import (
	"math/bits"
)

func (t *TicTacToe) Evaluate() float64 {
	scoreX := 0.0
	scoreO := 0.0
	fourX := 0
	fourO := 0
	threeX := 0
	threeO := 0

	vectors := t.GetAllVectors()

	for i := 0; i < len(vectors); i++ {
		for j := 0; j < len(PATTERN["X five"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X five"][j], 10)) * 1000000
		}
		for j := 0; j < len(PATTERN["X open four"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X open four"][j], 12)) * 100000
		}
		for j := 0; j < len(PATTERN["X close four"]); j++ {
			fourX += VectorPatternMatchCount(vectors[i], PATTERN["X close four"][j], 12)
		}
		for j := 0; j < len(PATTERN["X broken four"]); j++ {
			fourX += VectorPatternMatchCount(vectors[i], PATTERN["X broken four"][j], 10)
		}
		for j := 0; j < len(PATTERN["X open three"]); j++ {
			threeX += VectorPatternMatchCount(vectors[i], PATTERN["X open three"][j], 10)
		}
		for j := 0; j < len(PATTERN["X broken three"]); j++ {
			threeX += VectorPatternMatchCount(vectors[i], PATTERN["X broken three"][j], 12)
		}
		for j := 0; j < len(PATTERN["X open two"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X open two"][j], 8)) * 10
		}
		for j := 0; j < len(PATTERN["X broken two"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X broken two"][j], 10)) * 10
		}

		for j := 0; j < len(PATTERN["O five"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O five"][j], 10)) * 1000000
		}
		for j := 0; j < len(PATTERN["O open four"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O open four"][j], 12)) * 100000
		}
		for j := 0; j < len(PATTERN["O close four"]); j++ {
			fourO += VectorPatternMatchCount(vectors[i], PATTERN["O close four"][j], 12)
		}
		for j := 0; j < len(PATTERN["O broken four"]); j++ {
			fourO += VectorPatternMatchCount(vectors[i], PATTERN["O broken four"][j], 10)
		}
		for j := 0; j < len(PATTERN["O open three"]); j++ {
			threeO += VectorPatternMatchCount(vectors[i], PATTERN["O open three"][j], 10)
		}
		for j := 0; j < len(PATTERN["O broken three"]); j++ {
			threeO += VectorPatternMatchCount(vectors[i], PATTERN["O broken three"][j], 12)
		}
		for j := 0; j < len(PATTERN["O open two"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O open two"][j], 8)) * 10
		}
		for j := 0; j < len(PATTERN["O broken two"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O broken two"][j], 10)) * 10
		}
	}

	if fourX > 1 {
		scoreX += 100000
	}
	if fourO > 1 {
		scoreO += 100000
	}
	if threeX > 1 {
		scoreX += 40000
	}
	if threeO > 1 {
		scoreO += 80000
	}
	if fourX+threeX > 1 {
		scoreX += 90000
	}
	if fourO+threeO > 1 {
		scoreO += 90000
	}
	scoreX += float64(threeX)*100 + float64(fourX)*1000
	scoreO += float64(threeO)*100 + float64(fourO)*1000
	return scoreX - scoreO*0.8
}

func (t *TicTacToe) GetAllVectors() []int {
	res := []int{}
	res = append(res, t.BoardRow...)
	res = append(res, t.BoardCol...)
	res = append(res, t.getAllDiag()...)
	res = append(res, t.getAllAntiDiag()...)
	return res
}

// func VectorPatternMatchCount(vector int, pattern *[]int, player int) int {
// 	// vector is bit of a row / col / diagonal
// 	count := 0
// 	for i := range len(*vector) - len(*pattern) + 1 {
// 		var j int
// 		for j = 0; j < len(*pattern); j++ {
// 			if (*pattern)[j] == 1 && (*vector)[i+j] != player {
// 				break
// 			}
// 			if (*pattern)[j] == -1 && ((*vector)[i+j] != -player || (*vector)[i+j] != WALL) {
// 				break
// 			}
// 			if (*pattern)[j] == 0 && (*vector)[i+j] != 0 {
// 				break
// 			}
// 		}
// 		if j == len(*pattern) {
// 			count++
// 		}
// 	}
// 	return count
// }

// return bitmask of diagonal, length of bitmask
func (t *TicTacToe) getDiag(i, j int) (int, int) {
	n := len(t.BoardRow)
	row := i
	column := j
	res := 0
	for row < n && column < n {
		res = (res << 2) | t.GetValue(row, column)
		row++
		column++
	}
	return res, row - i
}

func (t *TicTacToe) getAllDiag() []int {
	n := len(t.BoardRow)
	res := []int{}
	for i := 0; i < n; i++ {
		diag, leng := t.getDiag(0, i)
		if leng >= 5 {
			res = append(res, diag)
		}
	}
	for i := 1; i < n; i++ {
		diag, leng := t.getDiag(i, 0)
		if leng >= 5 {
			res = append(res, diag)
		}
	}
	return res
}

func (t *TicTacToe) getAntiDiag(i, j int) (int, int) {
	n := len(t.BoardRow)
	row := i
	column := j
	res := 0
	for row < n && column >= 0 {
		res = (res << 2) | t.GetValue(row, column)
		row++
		column--
	}
	return res, row - i
}

func (t *TicTacToe) getAllAntiDiag() []int {
	n := len(t.BoardRow)
	res := []int{}
	for i := 0; i < n; i++ {
		diag, leng := t.getAntiDiag(0, i)
		if leng >= 5 {
			res = append(res, diag)
		}
	}
	for i := 1; i < n; i++ {
		diag, leng := t.getAntiDiag(i, n-1)
		if leng >= 5 {
			res = append(res, diag)
		}
	}
	return res
}

type VectorPatternCacheKey struct {
	vector  int
	pattern int
}

// var vectorPatternCache = map[VectorPatternCacheKey]int{}

// var vectorPatternCache = map[VectorPatternCacheKey]int{}

var cacheHit = 0

// var vectorPatternCache = sync.Map{}

func VectorPatternMatchCount(vector int, pattern int, patternLength int) int {
	// check if in cache
	// val, ok := vectorPatternCache[VectorPatternCacheKey{vector: vector, pattern: pattern}]
	// if ok {
	// 	cacheHit++
	// 	return val
	// }

	// val, ok := vectorPatternCache.Load(VectorPatternCacheKey{vector: vector, pattern: pattern})
	// if ok {
	// 	cacheHit++
	// 	return val.(int)
	// }

	// 1. get vector length
	vectorLength := bits.Len64(uint64(vector))

	// 2. shift vector and check if pattern match
	count := 0
	i := 0
	for i <= vectorLength-patternLength {
		vector >>= 2
		// fmt.Printf("vector: %b len: %v\n", vector, bits.Len64(uint64(vector)))
		compareVector := vector & ((1 << patternLength) - 1)
		// fmt.Printf("compareVector: %b\n", compareVector)
		if compareVector == pattern {
			count++
			break
		}
		i += 2
	}

	// vectorPatternCache[VectorPatternCacheKey{vector: vector, pattern: pattern}] = count
	// go func() {
	// 	vectorPatternCache.Store(VectorPatternCacheKey{vector: vector, pattern: pattern}, count)
	// }()
	return count
}

// func addToBitmask(bitmask int, value int, pos int) int {
// 	return bitmask | (value << pos)
// }
