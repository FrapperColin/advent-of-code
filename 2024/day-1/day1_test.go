package main

import "testing"

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	// Test cases
	tests := []struct {
		a, b     []int
		expected int
		expectedOccurence int
	}{
		{[]int{1, 4, 6}, []int{0, 5, 7}, 3, 0},
		{[]int{1, 4, 6}, []int{1, 4, 9}, 3, 5},
	}

	for _, test := range tests {
		t.Run("testing calculTotalDistance", func(t *testing.T) {
			result := calculTotalDistance(test.a, test.b)
			if result != test.expected {
				t.Errorf("calculTotalDistance(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
			}
		})

		t.Run("testing calculSimilarities", func(t *testing.T) {
			result := calculSimilarities(test.a, test.b)
			if result != test.expectedOccurence {
				t.Errorf("calculSo(%d, %d) = %d; expectedOccurence%d", test.a, test.b, result, test.expectedOccurence)
			}
		})
	}
}
