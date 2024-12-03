package utils

import (
	"reflect"
	"testing"
)

func TestReadFileByLine(t *testing.T) {
	path := "./sample.txt"
	expected := [][]string{{"33", "44"}, {"44", "33"}, {"22", "55"}, {"11", "33"}, {"33", "99"}, {"33", "33"}}
	_, actual := ReadFileByLine(path)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, wanted %v", actual, expected)
	}
}

func TestFilterSlice(t *testing.T) {
	slice := []string{"1", " ", "o", " ", "     ", "  ", "3343"}
	actual := FilterSlice(slice, IsNotBlank)
	expected := []string{"1", "o", "3343"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, wanted %v", actual, expected)
	}

	sliceTwo := []string{" ", "", "      "}
	actualTwo := FilterSlice(sliceTwo, IsNotBlank)
	expectedTwo := []string{}

	if !reflect.DeepEqual(actualTwo, expectedTwo) {
		t.Errorf("got %v, wanted %v", actualTwo, expectedTwo)
	}
}
