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

func FindOccurences(x int, list []int) int {
	acc := 0
	for _, num := range list {
		if num == x {
			acc++
		}
	}
	return acc
}

func SolvePartTwo(path string) int {
	err, input := utils.ReadFileByLine(path)
	if err != nil {
		log.Fatal(err)
	}
	listOne, listTwo := MakeTwoLists(input)
	acc := 0
	for _, num := range listOne {
		occur := FindOccurences(num, listTwo)
		acc += num * occur
	}
	return acc
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

func SolvePartOne(path string) int {
	err, input := utils.ReadFileByLine(path)
	if err != nil {
		log.Fatal(err)
	}
	listOne, listTwo := MakeTwoLists(input)
	return FindDistance(listOne, listTwo)
}
