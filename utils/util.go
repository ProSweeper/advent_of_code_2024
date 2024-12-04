package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileByLine(path string) (error, [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		newLine := FilterSlice(line, IsNotBlank)
		if len(newLine) > 0 {
			lines = append(lines, newLine)
		}
	}
	return nil, lines
}

func FilterSlice(s []string, callback func(string) bool) []string {
	slice := []string{}
	for _, str := range s {
		if callback(str) {
			slice = append(slice, str)
		}
	}

	return slice
}

func StringSliceToIntSlice(strSlice []string) []int {
	intSlice := []int{}
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err == nil {
			intSlice = append(intSlice, num)
		}
	}
	return intSlice
}

func IsNotBlank(str string) bool {
	if strings.Trim(str, " ") == "" {
		return false
	}
	return true
}
