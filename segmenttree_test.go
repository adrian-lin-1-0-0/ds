package ds

import (
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		expected int
	}{
		{
			name:     "Single element",
			vals:     []int{5},
			expected: 5,
		},
		{
			name:     "Multiple elements",
			vals:     []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "Empty array",
			vals:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.vals) == 0 {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for empty array, but no panic occurred")
					}
				}()
				NewSegmentTree(tt.vals)
				return
			}

			root := NewSegmentTree(tt.vals)
			if root.sum != tt.expected {
				t.Errorf("Expected sum %d, got %d", tt.expected, root.sum)
			}
		})
	}
}
func TestSegmentNode_SumRange(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		start    int
		end      int
		expected int
	}{
		{
			name:     "Full range",
			vals:     []int{1, 2, 3, 4},
			start:    0,
			end:      3,
			expected: 10,
		},
		{
			name:     "Single element range",
			vals:     []int{1, 2, 3, 4},
			start:    2,
			end:      2,
			expected: 3,
		},
		{
			name:     "Partial range",
			vals:     []int{1, 2, 3, 4},
			start:    1,
			end:      2,
			expected: 5,
		},
		{
			name:     "Another partial range",
			vals:     []int{1, 2, 3, 4, 5},
			start:    1,
			end:      3,
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewSegmentTree(tt.vals)
			result := root.SumRange(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("Expected sum %d, got %d", tt.expected, result)
			}
		})
	}
}
func TestSegmentNode_Update(t *testing.T) {
	tests := []struct {
		name         string
		initialVals  []int
		updateIndex  int
		updateVal    int
		expectedVals []int
	}{
		{
			name:         "Update single element",
			initialVals:  []int{1, 2, 3, 4},
			updateIndex:  2,
			updateVal:    10,
			expectedVals: []int{1, 2, 10, 4},
		},
		{
			name:         "Update first element",
			initialVals:  []int{5, 6, 7, 8},
			updateIndex:  0,
			updateVal:    15,
			expectedVals: []int{15, 6, 7, 8},
		},
		{
			name:         "Update last element",
			initialVals:  []int{3, 1, 4, 1},
			updateIndex:  3,
			updateVal:    9,
			expectedVals: []int{3, 1, 4, 9},
		},
		{
			name:         "Update middle element",
			initialVals:  []int{10, 20, 30, 40, 50},
			updateIndex:  2,
			updateVal:    25,
			expectedVals: []int{10, 20, 25, 40, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewSegmentTree(tt.initialVals)

			// Perform the update
			root.Update(tt.updateIndex, tt.updateVal)

			// Verify the updated values using SumRange
			for i := 0; i < len(tt.expectedVals); i++ {
				result := root.SumRange(i, i)
				if result != tt.expectedVals[i] {
					t.Errorf("At index %d, expected %d, got %d", i, tt.expectedVals[i], result)
				}
			}

			// Verify the total sum
			totalSum := root.SumRange(0, len(tt.expectedVals)-1)
			expectedTotalSum := 0
			for _, val := range tt.expectedVals {
				expectedTotalSum += val
			}
			if totalSum != expectedTotalSum {
				t.Errorf("Expected total sum %d, got %d", expectedTotalSum, totalSum)
			}
		})
	}
}
