package day02

import (
	"testing"
)

var allowedDeltas = []int{1, 2, 3}

func TestSolvePartOne(t *testing.T) {
	path := "./sample.txt"
	got := SolvePartOne(path)
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
func TestIsSafe(t *testing.T) {
	allowedDeltas := []int{1, 2, 3}

	testCases := []struct {
		input    []int
		expected bool
		name     string
	}{
		{
			input:    []int{7, 6, 4, 2, 1},
			expected: true,
			name:     "decreasing by 1 or 2",
		},
		{
			input:    []int{1, 2, 7, 8, 9},
			expected: false,
			name:     "increasing by 5",
		},
		{
			input:    []int{9, 7, 6, 2, 1},
			expected: false,
			name:     "decreasing by 4",
		},
		{
			input:    []int{1, 3, 2, 4, 5},
			expected: false,
			name:     "changing direction",
		},
		{
			input:    []int{8, 6, 4, 4, 1},
			expected: false,
			name:     "no change between 4 4",
		},
		{
			input:    []int{1, 3, 6, 7, 9},
			expected: true,
			name:     "increasing by 1, 2, or 3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsSafe(allowedDeltas, tc.input)
			if actual != tc.expected {
				t.Errorf("got %t, want %t for input %v", actual, tc.expected, tc.input)
			}
		})
	}
}

// func TestIsSafeIncrement(t *testing.T) {
// 	isSafeTests := []struct {
// 		num1         int
// 		num2         int
// 		isIncreasing bool
// 		want         bool
// 		testCase     string
// 	}{
// 		{1, 2, (false), false, "one"},
// 		{1, 2, false, true, "two"},
// 		{2, 1, false, true, "three"},
// 		{2, 1, (false), true, "four"},
// 		{1, 2, (true), true, "five"},
// 		{2, 1, (true), false, "six"},
// 		{2, 2, (true), false, "seven"},
// 		{2, 2, (false), false, "eight"},
// 		{2, 9, (false), false, "nine"},
// 		{2, 9, (true), false, "ten"},
// 		{9, 2, (true), false, "eleven"},
// 		{9, 2, (false), false, "twelve"},
// 	}
// 	for _, test := range isSafeTests {
// 		got := IsSafeIncrement(allowedDeltas, &status{true, test.isIncreasing}, test.num1, test.num2)
//
// 		if got != test.want {
// 			t.Errorf("got %t, want %t  going from %d -> %d, increasing == %v case: %s", got, test.want, test.num1, test.num2, test.isIncreasing, test.testCase)
// 		}
// 	}
// }
