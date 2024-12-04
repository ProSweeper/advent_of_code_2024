package day02

import (
	"log"
	"slices"

	utils "github.com/prosweeper/advent_of_code_2024/utils"
)

func IsSafeProcedure(allowedDeltas, powerLevels []int) bool {
	rate := 0
	for i, num := range powerLevels {
		if i == 1 {
			diff := num - powerLevels[0]
			if diff > 0 {
				rate = 1
			} else if diff < 0 {
				rate = -1
			}
		}
		if i > 0 {
			if !IsSafeIncrement(allowedDeltas, rate, powerLevels[i-1], num) {
				return false
			}
		}
	}
	return true
}

func IsSafeIncrement(allowedDeltas []int, rate, num1, num2 int) bool {
	diff := num2 - num1
	if diff == 0 {
		return false
	}
	return slices.Contains(allowedDeltas, rate*diff)
}

func SolvePartOne(file string) int {
	var allowedDeltas = []int{1, 2, 3}
	err, reports := utils.ReadFileByLine(file)
	if err != nil {
		log.Fatal(err)
	}
	intReports := [][]int{}
	for _, report := range reports {
		intReports = append(intReports, utils.StringSliceToIntSlice(report))
	}
	safeReports := 0
	for _, report := range intReports {
		if IsSafeProcedure(allowedDeltas, report) {
			safeReports++
		} else {
		}
	}

	return safeReports
}
