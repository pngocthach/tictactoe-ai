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
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X five"][j].Bitmask, 10)) * 10000
		}
		for j := 0; j < len(PATTERN["X open four"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X open four"][j].Bitmask, 12)) * 5000
		}
		for j := 0; j < len(PATTERN["X close four"]); j++ {
			fourX += VectorPatternMatchCount(vectors[i], PATTERN["X close four"][j].Bitmask, 12)
		}
		for j := 0; j < len(PATTERN["X broken four"]); j++ {
			fourX += VectorPatternMatchCount(vectors[i], PATTERN["X broken four"][j].Bitmask, 10)
		}
		for j := 0; j < len(PATTERN["X open three"]); j++ {
			threeX += VectorPatternMatchCount(vectors[i], PATTERN["X open three"][j].Bitmask, 10)
		}
		for j := 0; j < len(PATTERN["X broken three"]); j++ {
			threeX += VectorPatternMatchCount(vectors[i], PATTERN["X broken three"][j].Bitmask, 12)
		}
		for _, pattern := range PATTERN["X close three"] {
			scoreX += float64(VectorPatternMatchCount(vectors[i], pattern.Bitmask, pattern.Length)) * 10
		}
		for j := 0; j < len(PATTERN["X open two"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X open two"][j].Bitmask, 8)) * 20
		}
		for j := 0; j < len(PATTERN["X broken two"]); j++ {
			scoreX += float64(VectorPatternMatchCount(vectors[i], PATTERN["X broken two"][j].Bitmask, 10)) * 20
		}
		for j := 0; j < len(PATTERN["O five"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O five"][j].Bitmask, 10)) * 10000
		}
		for j := 0; j < len(PATTERN["O open four"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O open four"][j].Bitmask, 12)) * 5000
		}
		for j := 0; j < len(PATTERN["O close four"]); j++ {
			fourO += VectorPatternMatchCount(vectors[i], PATTERN["O close four"][j].Bitmask, 12)
		}
		for j := 0; j < len(PATTERN["O broken four"]); j++ {
			fourO += VectorPatternMatchCount(vectors[i], PATTERN["O broken four"][j].Bitmask, 10)
		}
		for j := 0; j < len(PATTERN["O open three"]); j++ {
			threeO += VectorPatternMatchCount(vectors[i], PATTERN["O open three"][j].Bitmask, 10)
		}
		for _, pattern := range PATTERN["O close three"] {
			scoreO += float64(VectorPatternMatchCount(vectors[i], pattern.Bitmask, pattern.Length)) * 10
		}
		for j := 0; j < len(PATTERN["O broken three"]); j++ {
			threeO += VectorPatternMatchCount(vectors[i], PATTERN["O broken three"][j].Bitmask, 12)
		}
		for j := 0; j < len(PATTERN["O open two"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O open two"][j].Bitmask, 8)) * 20
		}
		for j := 0; j < len(PATTERN["O broken two"]); j++ {
			scoreO += float64(VectorPatternMatchCount(vectors[i], PATTERN["O broken two"][j].Bitmask, 10)) * 20
		}
	}

	// if fourX > 1 {
	// 	scoreX += 10000
	// }
	// if fourO > 1 {
	// 	scoreO += 10000
	// }
	// if threeX > 1 {
	// 	scoreX += 4000
	// }
	// if threeO > 1 {
	// 	scoreO += 4000
	// }
	// if fourX+threeX > 1 {
	// 	scoreX += 10000
	// }
	// if fourO+threeO > 1 {
	// 	scoreO += 10000
	// }
	scoreX += float64(threeX)*100 + float64(fourX)*1000
	scoreO += float64(threeO)*100 + float64(fourO)*1000
	return scoreX - scoreO*EVAL_PARAM
}

func (t *TicTacToe) GetAllVectors() []int {
	res := []int{}
	res = append(res, t.BoardRow...)
	res = append(res, t.BoardCol...)
	res = append(res, t.getAllDiag()...)
	res = append(res, t.getAllAntiDiag()...)
	return res
}

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

// type VectorPatternCacheKey struct {
// 	vector  int
// 	pattern int
// }

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
