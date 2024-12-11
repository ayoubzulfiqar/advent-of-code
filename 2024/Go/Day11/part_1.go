package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// readLinesToList reads integers from a file into a slice of slices of integers.
func readLinesToList() [][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strValues := strings.Fields(line)
		intValues := make([]int, len(strValues))
		for i, val := range strValues {
			intValues[i], err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
		lines = append(lines, intValues)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

// blink implements the blink logic for a given stone.
func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		part1, _ := strconv.Atoi(s[:mid])
		part2, _ := strconv.Atoi(s[mid:])
		return []int{part1, part2}
	}

	return []int{stone * 2024}
}
