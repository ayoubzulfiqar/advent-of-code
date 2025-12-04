package main

import (
	"fmt"
	"strings"
	"time"
)

func ElevsAndForkLifts(input string) (string, time.Duration) {
	startTime := time.Now()

	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	height := len(grid)
	totalRemoved := 0

	for {
		cellsToRemove := []struct{ row, col int }{}

		for row := 0; row < height; row++ {
			currentLine := grid[row]
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

						neighborLine := grid[neighborRow]

						if neighborCol < 0 || neighborCol >= len(neighborLine) {
							continue
						}

						if neighborLine[neighborCol] == '@' {
							adjacentCount++
						}
					}
				}

				if adjacentCount < 4 {
					cellsToRemove = append(cellsToRemove, struct{ row, col int }{row, col})
				}
			}
		}

		if len(cellsToRemove) == 0 {
			break
		}

		for _, cell := range cellsToRemove {
			grid[cell.row][cell.col] = '.'
		}

		totalRemoved += len(cellsToRemove)
	}

	result := fmt.Sprintf("%d", totalRemoved)
	return result, time.Since(startTime)
}

// func main() {
// 	input, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		fmt.Printf("Error reading input.txt: %v\n", err)
// 		return
// 	}
// 	inputStr := strings.TrimSpace(string(input))

// 	part2Result, part2Time := ElevsAndForkLifts(inputStr)

// 	fmt.Printf("Part 2 Result: %s (Time: %v)\n", part2Result, part2Time)
// }
