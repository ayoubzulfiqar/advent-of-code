package main

import (
	"bufio"
	"fmt"
	"os"
)

func picoSecondsCheats() {
	// Read the input from the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Read all lines from the file into grid
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	// Find start and end positions
	start := findPosition(grid, 'S')
	end := findPosition(grid, 'E')

	// Track visited positions and their step counts
	track := make(map[[2]int]int)
	track[start] = 0
	cur := start
	curStep := 0

	// Perform BFS-like traversal
	for cur != end {
		curStep++
		i, j := cur[0], cur[1]
		found := false
		for _, dir := range [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			newI, newJ := i+dir[0], j+dir[1]
			if isValid(newI, newJ, grid) && !isInTrack(track, [2]int{newI, newJ}) && isValidChar(grid[newI][newJ]) {
				cur = [2]int{newI, newJ}
				track[cur] = curStep
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	// Count pairs that satisfy the condition
	count := 0
	for p1 := range track {
		for _, dir := range [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			newI, newJ := p1[0]+dir[0], p1[1]+dir[1]
			if isInTrack(track, [2]int{newI, newJ}) {
				continue
			}
			otherPos := [2]int{newI + dir[0], newJ + dir[1]}
			if isInTrack(track, otherPos) && track[otherPos]-track[p1] >= 102 {
				count++
			}
		}
	}

	fmt.Println(count)
}

// Helper function to find the position of a character in the grid
func findPosition(grid []string, char rune) [2]int {
	for i, line := range grid {
		for j, cell := range line {
			if cell == char {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

// Check if the position is valid within the grid
func isValid(i, j int, grid []string) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i])
}

// Check if a character is 'S', 'E' or '.'
func isValidChar(c byte) bool {
	return c == 'S' || c == 'E' || c == '.'
}

// Check if a position is already in the track map
func isInTrack(track map[[2]int]int, pos [2]int) bool {
	_, exists := track[pos]
	return exists
}
