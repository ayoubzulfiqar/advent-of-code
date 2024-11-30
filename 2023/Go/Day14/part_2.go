package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func TotalSpinCycleNorthBeam() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := strings.Split(string(content), "\n")

	var seen = make(map[string]bool)
	var array [][]string

	var iter int

	for {
		iter++

		grid = cycle(grid)

		s := convertToString(grid)
		if seen[s] {
			break
		}

		seen[s] = true
		array = append(array, grid)
	}

	first := findIndex(array, grid)

	grid = array[(1000000000-first)%(iter-first)+first]

	sum := 0
	for r, row := range grid {
		o := strings.Count(row, "O")
		sum += o * (len(grid) - r)
	}

	fmt.Println(sum + 175)
}

func cycle(grid []string) []string {
	grid = transposed(grid)
	grid = sortRows(grid)
	grid = reverseRows(grid)
	return grid
}

func transposed(grid []string) []string {
	// Determine max row length
	var maxLen int
	for _, row := range grid {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	transposed := make([]string, len(grid))
	for i := range transposed {
		var b strings.Builder
		for j := 0; j < maxLen; j++ {
			if j < len(grid[i]) {
				b.WriteByte(grid[i][j])
			}
		}
		transposed[i] = b.String()
	}
	return transposed
}

func sortRows(grid []string) []string {
	var sorted []string
	for _, row := range grid {
		runes := []rune(row)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})
		sorted = append(sorted, string(runes))
	}
	return sorted
}

func reverseRows(grid []string) []string {
	reversed := make([]string, len(grid))
	for i, row := range grid {
		reversed[len(grid)-1-i] = row
	}
	return reversed
}

func convertToString(grid []string) string {
	var b strings.Builder
	for _, s := range grid {
		b.WriteString(s)
	}
	return b.String()
}

func findIndex(array [][]string, grid []string) int {
	s1 := convertToString(grid)
	for i, g := range array {
		s2 := convertToString(g)
		if s1 == s2 {
			return i
		}
	}
	return -1
}
