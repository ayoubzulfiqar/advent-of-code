package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

func findTrails(grid [][]int, starting Position, part1 bool) int {
	width, height := len(grid[0]), len(grid) // Get grid dimensions
	queue := []Position{starting}            // Initialize the queue with the starting position
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
		directions := []Position{
			{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0},
		}

		// Iterate over each direction
		for _, direction := range directions {
			// Calculate the new position
			position := Position{
				x: current.x + direction.x,
				y: current.y + direction.y,
			}

			// Validate the new position
			if position.x < 0 || position.x >= width || position.y < 0 || position.y >= height {
				continue // Out of bounds, skip
			}

			// Check if the position is visited and it's part1
			visitedKey := fmt.Sprintf("%d,%d", position.x, position.y)
			if part1 && visited[visitedKey] {
				continue
			}

			// Check elevation constraint
			if grid[position.y][position.x]-grid[current.y][current.x] != 1 {
				continue
			}

			// Add the position to the queue
			queue = append(queue, position)

			// Mark as visited if part1
			if part1 {
				visited[visitedKey] = true
			}
		}
	}

	return paths
}

func trailHeadsOnTopographicMap() int {
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

	// Initialize total trailHead scores
	total := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Check if the current cell is a trailHead (value 0)
			if grid[y][x] == 0 {
				total += findTrails(grid, Position{x: x, y: y}, true)
			}
		}
	}

	fmt.Println(total) // Print the total
	return total
}
