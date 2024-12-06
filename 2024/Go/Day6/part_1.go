package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func guardDistinctPositions() {
	// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var inputData []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputData = append(inputData, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rows := len(inputData)
	cols := len(inputData[0])

	// Direction vectors (up, right, down, left)
	directions := [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	// Initialize variables for the guard's starting position and direction
	var guardRow, guardCol, guardDir int
	found := false
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if strings.ContainsRune("^>v<", rune(inputData[r][c])) {
				guardRow, guardCol = r, c
				guardDir = strings.Index("^>v<", string(inputData[r][c]))
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Set to track distinct visited positions
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d", guardRow, guardCol)] = true

	// Simulate guard movement
	for {
		dr, dc := directions[guardDir][0], directions[guardDir][1]
		nextRow, nextCol := guardRow+dr, guardCol+dc

		// Check if the next position is outside the grid
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			break // Guard leaves the grid
		}

		if inputData[nextRow][nextCol] == '#' {
			// Obstacle ahead: turn right (clockwise)
			guardDir = (guardDir + 1) % 4
		} else {
			// Move forward
			guardRow, guardCol = nextRow, nextCol
			visited[fmt.Sprintf("%d,%d", guardRow, guardCol)] = true
		}
	}

	// Output the number of distinct positions visited
	fmt.Printf("Distinct positions visited: %d\n", len(visited))
}
