package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readFileByLine(path string) (error, [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		newLine := filterSlice(line, isNotBlank)
		lines = append(lines, newLine)
	}
	return nil, lines
}

func filterSlice(s []string, callback func(string) bool) []string {
	slice := []string{}
	for _, str := range s {
		if callback(str) {
			slice = append(slice, str)
		}
	}

	return slice
}

func isNotBlank(str string) bool {
	if str == " " {
		return false
	}
	return true
}
