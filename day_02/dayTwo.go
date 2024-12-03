package day02

import (
	utils "github.com/prosweeper/advent_of_code_2024/utils"
	"log"
)

func IsSafe(allowedDeltas, deltas []int) bool {
	return false
}

func IsSafeIncrement(allowedDeltas []int, num1, num2 int) bool {
	diff := num2 - num1
	if diff == 0 {
		return false
	}

	return false
}

func SolvePartOne(file string) int {
	err, reports := utils.ReadFileByLine(file)
	var allowedDeltas = []int{1, 2, 3}
	if err != nil {
		log.Fatal(err)
	}
	safeReports := 0

	return safeReports
}
