package solver

import (
	"errors"
	"testing"
)

func TestSolveTask1(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		expected    int
		expectedErr error
	}{
		{"decreasing sequence 1", []int{5, 4, 3, 2, 1}, 0, nil},
		{"decreasing sequence 2", []int{10, 9, 8, 7, 6}, 0, nil},

		{"increasing sequence 1", []int{1, 2, 3, 4, 5}, 4, nil},
		{"increasing sequence 2", []int{0, 1, 2, 3, 4, 5}, 5, nil},
		{"increasing sequence 3", []int{-5, -4, -3, -2, -1, 0}, 5, nil},

		{"arbitrary sequences", []int{3, 5, 2, 6, 1, 4}, 5, nil},
		{"arbitrary sequences", []int{10, -5, 7, 2, 3, 15, 12}, 20, nil},

		{"length less than 2, one value", []int{1}, 0, errors.New("nums cannot have a length less than 2")},
		{"length less than 2, empty", []int{}, 0, errors.New("nums cannot have a length less than 2")},
	}

	solver := TaskSolver{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := solver.SolveTask1(test.nums)

			switch {
			case err != nil && test.expectedErr == nil:
				t.Errorf("got: %v, expected: %v", err, test.expectedErr)
			case err == nil && test.expectedErr != nil:
				t.Errorf("got: %v, expected: %v", err, test.expectedErr)
			case err != nil && test.expectedErr != nil:
				if err.Error() != test.expectedErr.Error() {
					t.Errorf("got: %v, expected: %v", err, test.expectedErr)
				}
			}

			if result != test.expected {
				t.Errorf("got %d, expected %d", result, test.expected)
			}

		})

	}

}

func TestSolveTask2(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		{"equal strings 1", "abc", "abc", true},
		{"equal strings 2", "", "", true},
		{"equal strings 3", "12345", "12345", true},

		{"different strings 1", "abc", "def", false},
		{"different strings 2", "", "xyz", false},

		{"strings with same characters but different order 1", "abc", "bca", true},
		{"strings with same characters but different order 2", "12345", "54312", true},

		{"strings with different lengths 1", "abc", "abcd", false},
		{"strings with different lengths 2", "abcd", "abc", false},
		{"strings with different lengths 3", "", "a", false},
		{"strings with different lengths 4", "a", "", false},
	}

	solver := TaskSolver{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := solver.SolveTask2(test.str1, test.str2)

			if result != test.expected {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})

	}
}

func TestSolveTask3(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		numRows     int
		expected    string
		expectedErr error
	}{
		{"valid test case 1", "incomprehensibilities", 4, "iriinpesbtecmhniisoel", nil},
		{"valid test case 2", "incomprehensibilities", 3, "imhiisnopeesbltecrnii", nil},
		{"valid test case 3", "ABCDE", 1, "ABCDE", nil},
		{"valid test case 4", "ABCDEF", 2, "ACEBDF", nil},

		{"with errors 1", "", 3, "", errors.New("string cannot have zero length")},
		{"with errors 2", "ABC", 0, "", errors.New("num rows cannot be less than or equal to zero")},
	}

	solver := TaskSolver{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := solver.SolveTask3(test.s, test.numRows)

			switch {
			case err != nil && test.expectedErr == nil:
				t.Errorf("got: %v, expected: %v", err, test.expectedErr)
			case err == nil && test.expectedErr != nil:
				t.Errorf("got: %v, expected: %v", err, test.expectedErr)
			case err != nil && test.expectedErr != nil:
				if err.Error() != test.expectedErr.Error() {
					t.Errorf("got: %v, expected: %v", err, test.expectedErr)
				}
			}

			if result != test.expected {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})

	}
}
