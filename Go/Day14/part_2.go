// package main

// import (
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func Run() {
// 	content, err := os.ReadFile("input.txt") // Change "file.txt" to the actual file name
// 	if err != nil {
// 		panic(err)
// 	}

// 	grid := strings.Split(string(content), "\n")

// 	seen := make(map[string]struct{})
// 	array := [][]string{makeCopy(grid)}
// 	iter := 0

// 	for {
// 		iter++
// 		cycle(&grid)
// 		gridString := strings.Join(grid, "\n")

// 		if _, ok := seen[gridString]; ok {
// 			break
// 		}

// 		seen[gridString] = struct{}{}
// 		array = append(array, makeCopy(grid))
// 	}

// 	first := findFirstAppearance(array, grid)
// 	index := (1000000000-first)%(iter-first) + first
// 	grid = array[index]

// 	count := 0
// 	for r, row := range grid {
// 		count += strings.Count(row, "O") * (len(grid) - r)
// 	}

// 	fmt.Println(count)
// }

// func cycle(grid *[]string) {
// 	for i := 0; i < 4; i++ {
// 		*grid = transpose(*grid)
// 		*grid = processGrid(*grid)
// 		*grid = reverseRows(*grid)
// 	}
// }

// func transpose(grid []string) []string {
// 	if len(grid) == 0 {
// 		return grid
// 	}

// 	transposed := make([]string, len(grid[0]))
// 	for i := range transposed {
// 		var sb strings.Builder
// 		for _, row := range grid {
// 			if i < len(row) {
// 				sb.WriteByte(row[i])
// 			}
// 		}
// 		transposed[i] = sb.String()
// 	}
// 	return transposed
// }

// func processGrid(grid []string) []string {
// 	for i, row := range grid {
// 		groups := strings.Split(row, "#")
// 		for j, group := range groups {
// 			characters := strings.Split(group, "")
// 			sort.Sort(sort.Reverse(sort.StringSlice(characters)))
// 			groups[j] = strings.Join(characters, "")
// 		}
// 		grid[i] = strings.Join(groups, "#")
// 	}
// 	return grid
// }

// func reverseRows(grid []string) []string {
// 	for i := range grid {
// 		runes := []rune(grid[i])
// 		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
// 			runes[j], runes[k] = runes[k], runes[j]
// 		}
// 		grid[i] = string(runes)
// 	}
// 	return grid
// }

// func findFirstAppearance(array [][]string, target []string) int {
// 	for i, grid := range array {
// 		if equalGrids(grid, target) {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func equalGrids(grid1, grid2 []string) bool {
// 	if len(grid1) != len(grid2) {
// 		return false
// 	}

// 	for i := range grid1 {
// 		if grid1[i] != grid2[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func makeCopy(grid []string) []string {
// 	copyGrid := make([]string, len(grid))
// 	copy(copyGrid, grid)
// 	return copyGrid
// }

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
