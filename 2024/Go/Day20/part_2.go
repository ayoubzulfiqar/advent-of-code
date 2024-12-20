package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func saveHundredPicoSecondCheats() {
	// Read the input from the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Read the grid into a 2D slice
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	// Find start and end positions
	start := reallyFindPosition(grid, 'S')
	end := reallyFindPosition(grid, 'E')

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
			if isReallyValid(newI, newJ, grid) && !isReallyInTrack(track, [2]int{newI, newJ}) && isReallyValidChar(grid[newI][newJ]) {
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

	// Count valid pairs
	count := 0
	for coords := range track {
		potentials := cheatEndpoints(coords, track)
		for otherCoords := range potentials {
			if track[otherCoords]-track[coords]-manhattanDistance(coords, otherCoords) >= 100 {
				count++
			}
		}
	}

	fmt.Println(count)
}

// Helper function to find the position of a character in the grid
func reallyFindPosition(grid [][]rune, char rune) [2]int {
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
func isReallyValid(i, j int, grid [][]rune) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i])
}

// Check if a character is 'S', 'E' or '.'
func isReallyValidChar(c rune) bool {
	return c == 'S' || c == 'E' || c == '.'
}

// Check if a position is already in the track map
func isReallyInTrack(track map[[2]int]int, pos [2]int) bool {
	_, exists := track[pos]
	return exists
}

// Function to find possible endpoints in the neighborhood
func cheatEndpoints(coords [2]int, track map[[2]int]int) map[[2]int]struct{} {
	i, j := coords[0], coords[1]
	output := make(map[[2]int]struct{})
	for di := -20; di <= 20; di++ {
		djMax := 20 - int(math.Abs(float64(di)))
		for dj := -djMax; dj <= djMax; dj++ {
			newPos := [2]int{i + di, j + dj}
			if _, exists := track[newPos]; exists {
				output[newPos] = struct{}{}
			}
		}
	}
	return output
}

// Manhattan distance between two coordinates
func manhattanDistance(coord1, coord2 [2]int) int {
	return int(math.Abs(float64(coord1[0]-coord2[0]))) + int(math.Abs(float64(coord1[1]-coord2[1])))
}
