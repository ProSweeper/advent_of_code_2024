package day01

import (
	"log"
	"slices"
	"strconv"

	utils "github.com/prosweeper/advent_of_code_2024/utils"
)

func FindDistance(listOne, listTwo []int) int {
	slices.Sort(listOne)
	slices.Sort(listTwo)
	totalDistance := 0
	for i, num := range listOne {
		dist := AbsoluteDiff(num, listTwo[i])
		totalDistance += dist
	}
	return totalDistance
}

func AbsoluteDiff(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func MakeTwoLists(input [][]string) ([]int, []int) {
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

func Solve(path string) int {
	err, input := utils.ReadFileByLine(path)
	if err != nil {
		log.Fatal(err)
	}
	listOne, listTwo := MakeTwoLists(input)
	return FindDistance(listOne, listTwo)
}
