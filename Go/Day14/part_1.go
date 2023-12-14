package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func TotalLoadOnNorthSupport() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := strings.Split(string(content), "\n")
	grid = transpose(grid)
	grid = processGrid(grid)
	grid = transpose(grid)

	count := 0
	for r, row := range grid {
		count += strings.Count(row, "O") * (len(grid) - r)
	}

	fmt.Println(count)
}

func transpose(grid []string) []string {
	if len(grid) == 0 {
		return grid
	}

	transposed := make([]string, len(grid[0]))
	for i := range transposed {
		var sb strings.Builder
		for _, row := range grid {
			if i < len(row) {
				sb.WriteByte(row[i])
			}
		}
		transposed[i] = sb.String()
	}
	return transposed
}

func processGrid(grid []string) []string {
	for i, row := range grid {
		groups := strings.Split(row, "#")
		for j, group := range groups {
			characters := strings.Split(group, "")
			sort.Sort(sort.Reverse(sort.StringSlice(characters)))
			groups[j] = strings.Join(characters, "")
		}
		grid[i] = strings.Join(groups, "#")
	}
	return grid
}
