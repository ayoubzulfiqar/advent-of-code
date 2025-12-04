package main

import (
	"fmt"
	"strings"
	"time"
)

/*  With Grid Padding Appraoch


func ForkListRolls(input string) (string, time.Duration) {
	startTime := time.Now()

	lines := strings.Split(strings.TrimSpace(input), "\n")

	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, maxWidth)
		// Copy characters and fill remaining with null character or default
		for j := 0; j < maxWidth; j++ {
			if j < len(line) {
				grid[i][j] = rune(line[j])
			} else {
				grid[i][j] = '.'
			}
		}
	}

	height := len(grid)
	width := len(grid[0])
	accessibleCount := 0

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if grid[row][col] != '@' {
				continue
			}

			adjacentCount := 0

			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					// Skip the cell itself
					if dr == 0 && dc == 0 {
						continue
					}

					neighborRow := row + dr
					neighborCol := col + dc

					if neighborRow < 0 || neighborRow >= height || neighborCol < 0 || neighborCol >= width {
						continue
					}

					if grid[neighborRow][neighborCol] == '@' {
						adjacentCount++
					}
				}
			}

			if adjacentCount < 4 {
				accessibleCount++
			}
		}
	}

	result := fmt.Sprintf("%d", accessibleCount)
	return result, time.Since(startTime)
}




*/

func ForkListRolls(input string) (string, time.Duration) {
	startTime := time.Now()

	lines := strings.Split(strings.TrimSpace(input), "\n")

	height := len(lines)
	accessibleCount := 0

	for row := 0; row < height; row++ {
		currentLine := lines[row]
		width := len(currentLine)

		for col := 0; col < width; col++ {
			if currentLine[col] != '@' {
				continue
			}

			adjacentCount := 0

			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}

					neighborRow := row + dr
					neighborCol := col + dc

					if neighborRow < 0 || neighborRow >= height {
						continue
					}

					neighborLine := lines[neighborRow]

					if neighborCol < 0 || neighborCol >= len(neighborLine) {
						continue
					}

					if neighborLine[neighborCol] == '@' {
						adjacentCount++
					}
				}
			}

			if adjacentCount < 4 {
				accessibleCount++
			}
		}
	}

	result := fmt.Sprintf("%d", accessibleCount)
	return result, time.Since(startTime)
}

// func main() {
//
// 	input, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		fmt.Printf("Error reading input.txt: %v\n", err)
// 		return
// 	}
// 	inputStr := strings.TrimSpace(string(input))

//
// 	part1Result, part1Time := ForkListRolls(inputStr)

//
// 	fmt.Printf("Part 1 Result: %s (Time: %v)\n", part1Result, part1Time)
// }
