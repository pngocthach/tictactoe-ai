package game

import "testing"

func TestVectorPatternMatchCount(t *testing.T) {
	t.Run("Pattern exists", func(t *testing.T) {
		vector := []int{0, 0, -1, -1, 0, 0}

		pattern := []int{0, 1, 1, 0}
		player := -1
		expected := 1
		result := VectorPatternMatchCount(&vector, &pattern, player)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Pattern exists", func(t *testing.T) {
		vector := []int{0, -1, -2, -1, -1, -1, -1, 0}

		pattern := []int{-2, 1, 1, 1, 1, 0}
		player := -1
		expected := 1
		result := VectorPatternMatchCount(&vector, &pattern, player)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Pattern exists", func(t *testing.T) {
		vector := []int{0, -1, -2, -1, -1, -1, -1, -1, 0}

		pattern := []int{1, 1, 1, 1, 1}
		player := -1
		expected := 1
		result := VectorPatternMatchCount(&vector, &pattern, player)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

}
