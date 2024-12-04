package main

import (
	"bufio"
	"os"
)

func timesXMasAppears() int {
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

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 'A' {
				// Check pattern M.M
				//              .A.
				//              S.S
				cond := grid[i-1][j-1] == 'M' &&
					grid[i-1][j+1] == 'M' &&
					grid[i+1][j+1] == 'S' &&
					grid[i+1][j-1] == 'S'
				if cond {
					count++
				}

				// Check pattern S.S
				//              .A.
				//              M.M
				cond = grid[i-1][j-1] == 'S' &&
					grid[i-1][j+1] == 'S' &&
					grid[i+1][j+1] == 'M' &&
					grid[i+1][j-1] == 'M'
				if cond {
					count++
				}

				// Check pattern M.S
				//              .A.
				//              M.S
				cond = grid[i-1][j-1] == 'M' &&
					grid[i-1][j+1] == 'S' &&
					grid[i+1][j+1] == 'S' &&
					grid[i+1][j-1] == 'M'
				if cond {
					count++
				}

				// Check pattern S.M
				//              .A.
				//              S.M
				cond = grid[i-1][j-1] == 'S' &&
					grid[i-1][j+1] == 'M' &&
					grid[i+1][j+1] == 'M' &&
					grid[i+1][j-1] == 'S'
				if cond {
					count++
				}
			}
		}
	}

	return count
}
