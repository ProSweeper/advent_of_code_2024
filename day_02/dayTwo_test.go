package day02

import (
	"testing"
)

func TestIsSafe(t *testing.T) {
	allowedDeltas := []int{1, 2, 3}
	inputOne := []int{7, 6, 4, 2, 1}
	expectedOne := true
	actualOne := IsSafe(allowedDeltas, inputOne)
	if expectedOne != actualOne {
		t.Errorf("got %t, want %t", actualOne, expectedOne)
	}
	inputTwo := []int{1, 2, 7, 8, 9}
	expectedTwo := false
	actualTwo := IsSafe(allowedDeltas, inputTwo)
	if expectedTwo != actualTwo {
		t.Errorf("got %t, want %t", actualTwo, expectedTwo)
	}
}
