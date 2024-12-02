package main

import (
	"reflect"
	"testing"
)

func TestReadFileByLine(t *testing.T) {
	path := "./sample.txt"
	expected := [][]string{{"3", "4"}, {"4", "3"}, {"2", "5"}, {"1", "3"}, {"3", "9"}, {"3", "3"}}
	_, actual := readFileByLine(path)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, wanted %v", actual, expected)
	}
}

func TestFilterSlice(t *testing.T) {
	slice := []string{"1", " ", "o", " "}
	expected := []string{"1", "o"}
	actual := filterSlice(slice, isNotBlank)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, wanted %v", actual, expected)
	}
}
