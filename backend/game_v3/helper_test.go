package game3

import "testing"

func TestGetBinaryPattern(t *testing.T) {
	t.Run("Simple pattern", func(t *testing.T) {
		pattern := []int{0, 1, 0, 1}
		expected := 0b00010001
		result := GetBinaryPattern(pattern)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Five X", func(t *testing.T) {
		pattern := []int{x, x, x, x, x}
		expected := 0b0101010101
		result := GetBinaryPattern(pattern)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Five O", func(t *testing.T) {
		pattern := []int{o, o, o, o, o}
		expected := 0b1010101010
		result := GetBinaryPattern(pattern)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Open four X", func(t *testing.T) {
		pattern := []int{e, x, x, x, x, e}
		expected := 0b000101010100
		result := GetBinaryPattern(pattern)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}
