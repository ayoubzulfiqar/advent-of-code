package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Positions struct {
	x, y int
}

func findTrailing(grid [][]int, starting Positions, part1 bool) int {
	width, height := len(grid[0]), len(grid) // Get grid dimensions
	queue := []Positions{starting}           // Initialize the queue with the starting position
	visited := make(map[string]bool)         // Map to track visited positions (used only in part1)

	paths := 0 // Initialize path counter
	for len(queue) > 0 {
		// Dequeue the first element
		current := queue[0]
		queue = queue[1:]

		// If the current position is the target (value 9), increment paths
		if grid[current.y][current.x] == 9 {
			paths++
			continue
		}

		// Directions to move: up, right, down, left
		directions := []Positions{
			{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0},
		}

		// Iterate over each direction
		for _, direction := range directions {
			// Calculate the new position
			positions := Positions{
				x: current.x + direction.x,
				y: current.y + direction.y,
			}

			// Validate the new position
			if positions.x < 0 || positions.x >= width || positions.y < 0 || positions.y >= height {
				continue // Out of bounds, skip
			}

			// Check if the position is visited and it's part1
			visitedKey := fmt.Sprintf("%d,%d", positions.x, positions.y)
			if part1 && visited[visitedKey] {
				continue
			}

			// Check elevation constraint
			if grid[positions.y][positions.x]-grid[current.y][current.x] != 1 {
				continue
			}

			// Add the position to the queue
			queue = append(queue, positions)

			// Mark as visited if part1
			if part1 {
				visited[visitedKey] = true
			}
		}
	}

	return paths
}

func ratingOfAllTrailHeads() int {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err) // Handle file opening errors
	}
	defer file.Close()

	// Read the input and construct the grid
	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, num := range strings.Split(line, "") {
			val, _ := strconv.Atoi(num) // Convert string to integer
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	width, height := len(grid[0]), len(grid) // Get grid dimensions

	// Initialize total trailHead ratings
	total := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Check if the current cell is a trailHead (value 0)
			if grid[y][x] == 0 {
				total += findTrailing(grid, Positions{x: x, y: y}, false)
			}
		}
	}

	fmt.Println(total) // Print the total
	return total
}
