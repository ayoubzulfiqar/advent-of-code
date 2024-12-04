package main

import (
	"bufio"
	"os"
)

func xMasAppear() int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	count := 0
	XMAS := "XMAS"

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'X' {
				// horizontal
				found := true
				for k := 0; k < 4; k++ {
					if j+k >= len(grid[0]) || grid[i][j+k] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// horizontal reverse
				found = true
				for k := 0; k < 4; k++ {
					if j-k < 0 || grid[i][j-k] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// vertical
				found = true
				for k := 0; k < 4; k++ {
					if i+k >= len(grid) || grid[i+k][j] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// vertical reverse
				found = true
				for k := 0; k < 4; k++ {
					if i-k < 0 || grid[i-k][j] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// diagonal
				found = true
				for k := 0; k < 4; k++ {
					if i+k >= len(grid) || j+k >= len(grid[0]) || grid[i+k][j+k] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// diagonal reverse
				found = true
				for k := 0; k < 4; k++ {
					if i-k < 0 || j-k < 0 || grid[i-k][j-k] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// off-diagonal
				found = true
				for k := 0; k < 4; k++ {
					if i-k < 0 || j+k >= len(grid[0]) || grid[i-k][j+k] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}

				// off-diagonal reverse
				found = true
				for k := 0; k < 4; k++ {
					if i+k >= len(grid) || j-k < 0 || grid[i+k][j-k] != rune(XMAS[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}
			}
		}
	}

	return count
}
