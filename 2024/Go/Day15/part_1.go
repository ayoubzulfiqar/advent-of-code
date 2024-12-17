package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type Point struct {
// 	x int
// 	y int
// }

// // Define the directions as a map of characters to Points
// var DIRECTIONS = map[rune]Point{
// 	'^': {x: 0, y: -1},
// 	'>': {x: 1, y: 0},
// 	'v': {x: 0, y: 1},
// 	'<': {x: -1, y: 0},
// }

// Reads the input file and calculates the sum of GPS coordinates
func sumOfGPSCoordinates() int {
	// Read the file
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return 0
	}

	// Parse the input into grid and instructions
	parts := strings.Split(strings.TrimSpace(content.String()), "\n\n")
	grid := [][]rune{}
	for _, line := range strings.Split(parts[0], "\n") {
		grid = append(grid, []rune(line))
	}
	instructions := parts[1]

	width, height := len(grid[0]), len(grid)

	// Recursive function to move a box
	var moveBox func(position Point, direction Point) bool
	moveBox = func(position Point, direction Point) bool {
		next := Point{x: position.x + direction.x, y: position.y + direction.y}

		if grid[next.y][next.x] == '.' {
			// If the next spot is empty, swap positions
			grid[position.y][position.x], grid[next.y][next.x] = grid[next.y][next.x], grid[position.y][position.x]
			return true
		} else if grid[next.y][next.x] == '#' {
			// If the next spot is a wall, stop all boxes from moving
			return false
		} else {
			// Only move the current box if the next box can move
			if moveBox(next, direction) {
				grid[position.y][position.x], grid[next.y][next.x] = grid[next.y][next.x], grid[position.y][position.x]
				return true
			}
		}
		return false // This should never be reached
	}

	// Find the robot and clear its position
	robot := Point{x: 0, y: 0}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '@' {
				robot = Point{x: x, y: y}
				grid[y][x] = '.'
			}
		}
	}

	// Process each instruction
	for _, instruction := range instructions {
		direction := DIRECTIONS[instruction]
		position := Point{x: robot.x + direction.x, y: robot.y + direction.y}

		// If there is a wall, don't move
		if grid[position.y][position.x] != '#' {
			// If there is an empty spot, move without moving boxes
			if grid[position.y][position.x] == '.' {
				robot = position
			}
			// If there is a box, try to move all the boxes, then move
			if grid[position.y][position.x] == 'O' && moveBox(position, direction) {
				robot = position
			}
		}
	}

	// Tally all the box positions
	score := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' {
				score += y*100 + x
			}
		}
	}

	fmt.Println(score)
	return score
}
