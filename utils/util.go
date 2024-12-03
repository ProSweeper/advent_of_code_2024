package utils

import (
	"bufio"
	"log"
	"os"
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
		lines = append(lines, newLine)
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

func IsNotBlank(str string) bool {
	if strings.Trim(str, " ") == "" {
		return false
	}
	return true
}
