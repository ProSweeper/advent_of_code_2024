package main

import (
	"slices"
	"strconv"
)

func findDistance(listOne, listTwo []int) int {
	slices.Sort(listOne)
	slices.Sort(listTwo)
	totalDistance := 0
	for i, num := range listOne {
		totalDistance += absoluteDiff(num, listTwo[i])
	}
	return totalDistance
}

func absoluteDiff(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func makeTwoLists(input [][]string) ([]int, []int) {
	var listOne []int
	var listTwo []int

	for _, list := range input {
		intOne, _ := strconv.Atoi(list[0])
		intTwo, _ := strconv.Atoi(list[1])
		listOne = append(listOne, intOne)
		listTwo = append(listTwo, intTwo)
	}
	return listOne, listTwo
}
