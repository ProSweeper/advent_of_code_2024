package day01

import (
	"reflect"
	"testing"
)

func TestFindDistance(t *testing.T) {
	listOne := []int{3, 4, 2, 1, 3, 3}
	listTwo := []int{4, 3, 5, 3, 9, 3}
	expectedDistance := 11
	actualDistance := FindDistance(listOne, listTwo)

	if expectedDistance != actualDistance {
		t.Errorf("got %d, expected %d", actualDistance, expectedDistance)
	}
}

func TestMakeTwoLists(t *testing.T) {
	input := [][]string{{"3", "4"}, {"4", "3"}, {"2", "5"}, {"1", "3"}, {"3", "9"}, {"3", "3"}}
	listOne := []int{3, 4, 2, 1, 3, 3}
	listTwo := []int{4, 3, 5, 3, 9, 3}

	actualOne, actualTwo := MakeTwoLists(input)

	if !reflect.DeepEqual(actualOne, listOne) || !reflect.DeepEqual(actualTwo, listTwo) {
		t.Errorf("got: %v and %v, wanted %v and %v", actualOne, actualTwo, listOne, listTwo)
	}
}

func TestSolveOne(t *testing.T) {
	file := "./sample.txt"
	expected := 11
	actual := SolvePartOne(file)

	if expected != actual {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}

func TestSolveTwo(t *testing.T) {
	file := "./sample.txt"
	expected := 31
	actual := SolvePartTwo(file)

	if expected != actual {
		t.Errorf("got %d, expected %d", actual, expected)
	}
}
