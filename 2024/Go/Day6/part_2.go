package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	rows, cols int
	inputData  []string
	directions = [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}
	startRow, startCol, startDir int
)

// Simulate guard movement with an optional extra obstacle
func simulateWithObstacle(obstacleRow, obstacleCol int) bool {
	guardRow, guardCol, guardDir := startRow, startCol, startDir
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d,%d", guardRow, guardCol, guardDir)] = true

	for {
		dr, dc := directions[guardDir][0], directions[guardDir][1]
		nextRow, nextCol := guardRow+dr, guardCol+dc

		// Check if the next position is outside the grid
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			return false // Guard exits the grid
		}

		// Treat the additional obstacle as if it's a `#`
		nextCell := "#"
		if !(nextRow == obstacleRow && nextCol == obstacleCol) {
			nextCell = string(inputData[nextRow][nextCol])
		}

		if nextCell == "#" {
			// Obstacle ahead, turn right
			guardDir = (guardDir + 1) % 4
		} else {
			// Move forward
			guardRow = nextRow
			guardCol = nextCol
		}

		state := fmt.Sprintf("%d,%d,%d", guardRow, guardCol, guardDir)
		if visited[state] {
			return true // Loop detected
		}
		visited[state] = true
	}
}

func differentPositionForObstruction() {
	// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputData = append(inputData, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rows = len(inputData)
	cols = len(inputData[0])

	// Locate the guard's initial position and direction
	found := false
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if strings.ContainsRune("^>v<", rune(inputData[r][c])) {
				startRow, startCol = r, c
				startDir = strings.Index("^>v<", string(inputData[r][c]))
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Count valid positions for the new obstruction
	validPositions := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Skip positions that are not empty or are the starting position
			if inputData[r][c] == '#' || (r == startRow && c == startCol) {
				continue
			}

			// Simulate guard movement with an obstacle at (r, c)
			if simulateWithObstacle(r, c) {
				validPositions++
			}
		}
	}

	fmt.Printf("Number of valid positions for a new obstruction: %d\n", validPositions)
}
