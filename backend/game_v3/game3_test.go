package game3

import "testing"

func TestGetValueFromRow(t *testing.T) {
	tests := []struct {
		row      int
		col      int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{3, 0, 3},
		{3, 1, 0},
		{3, 2, 1},
		{3, 3, 2},
		{3, 4, 3},
	}

	// 11 00 10 01 11
	//  3  0  2  1  3

	for _, test := range tests {
		result := GetValueFromRow(test.row, test.col)
		if result != test.expected {
			t.Errorf("For row %d and col %d, expected %d, but got %d", test.row, test.col, test.expected, result)
		}
	}
}
