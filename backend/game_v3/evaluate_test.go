package game3

import "testing"

func TestVectorPatternMatchCount(t *testing.T) {
	t.Run("Pattern exists", func(t *testing.T) {
		vector := 0b11010111
		pattern := 0b0101
		patternLength := 4
		expected := 1
		result := VectorPatternMatchCount(vector, pattern, patternLength)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Pattern does not exist", func(t *testing.T) {
		vector := 0b0000000000000000000000000000000000000000000000000000000000000000
		pattern := 0b111111111111111111111111111111111111111111111111111111111111111
		patternLength := 64
		expected := 0
		result := VectorPatternMatchCount(vector, pattern, patternLength)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Pattern exists", func(t *testing.T) {
		vector := 0b1110101001010101011100
		pattern := 0b0101010101
		patternLength := 10
		expected := 1
		result := VectorPatternMatchCount(vector, pattern, patternLength)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Pattern exists", func(t *testing.T) {
		vector := 0b110010101100000
		pattern := 0b0010101100
		patternLength := 10
		expected := 1
		result := VectorPatternMatchCount(vector, pattern, patternLength)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	// t.Run("Pattern does not exist", func(t *testing.T) {
	// 	vector := 0b1010101010101010101010101010101010101010101010101010101010101010
	// 	pattern := 0b0101010101010101010101010101010101010101010101010101010101010101
	// 	patternLength := 64
	// 	expected := 0
	// 	result := VectorPatternMatchCount(vector, pattern, patternLength)
	// 	if result != expected {
	// 		t.Errorf("Expected %d, but got %d", expected, result)
	// 	}
	// })
}
