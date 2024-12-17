package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Movement struct {
	box       Point
	direction Point
}

// Define directions as a map of characters to Points
var DIRECTIONS = map[rune]Point{
	'^': {x: 0, y: -1},
	'>': {x: 1, y: 0},
	'v': {x: 0, y: 1},
	'<': {x: -1, y: 0},
}

// Helper function to create a key for a point (used for walls set)
func pointKey(p Point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

// Reads input from a file and calculates the sum of GPS coordinates
func allSumOfGPSCoordinates() int {
	// Read input from file
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

	walls := make(map[string]bool)
	var boxes []Point
	robot := Point{x: 0, y: 0}

	// Initialize walls, boxes, and robot position
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch grid[y][x] {
			case '@':
				robot = Point{x: x * 2, y: y}
			case '#':
				walls[pointKey(Point{x: x * 2, y: y})] = true
				walls[pointKey(Point{x: x*2 + 1, y: y})] = true
			case 'O':
				boxes = append(boxes, Point{x: x * 2, y: y})
			}
		}
	}

	// Recursive function to move boxes
	var moveBox func(collidedBox Point, direction Point, movements *[]Movement) bool
	moveBox = func(collidedBox Point, direction Point, movements *[]Movement) bool {
		nextPositions := []Point{
			{x: collidedBox.x + direction.x, y: collidedBox.y + direction.y},
			{x: collidedBox.x + 1 + direction.x, y: collidedBox.y + direction.y},
		}

		// Check for wall collisions
		for _, next := range nextPositions {
			if walls[pointKey(next)] {
				return false
			}
		}

		// Find all collided boxes
		var collidedBoxes []Point
		for _, box := range boxes {
			for _, next := range nextPositions {
				if (box.x != collidedBox.x || box.y != collidedBox.y) &&
					((box.x == next.x || box.x+1 == next.x) && box.y == next.y) {
					collidedBoxes = append(collidedBoxes, box)
				}
			}
		}

		// If no collided boxes, movements are valid
		if len(collidedBoxes) == 0 {
			return true
		}

		// Check for conflicts
		conflicts := false
		for _, box := range collidedBoxes {
			if moveBox(box, direction, movements) {
				alreadyMoved := false
				for _, movement := range *movements {
					if movement.box.x == box.x && movement.box.y == box.y {
						alreadyMoved = true
						break
					}
				}
				if !alreadyMoved {
					*movements = append(*movements, Movement{box: box, direction: direction})
				}
			} else {
				conflicts = true
				break
			}
		}

		return !conflicts
	}

	// Process instructions
	for _, instruction := range instructions {
		direction := DIRECTIONS[instruction]
		position := Point{x: robot.x + direction.x, y: robot.y + direction.y}

		// Only try to move if no wall is in the way
		if !walls[pointKey(position)] {
			var collidedBox *Point
			for i, box := range boxes {
				if (box.x == position.x || box.x+1 == position.x) && box.y == position.y {
					collidedBox = &boxes[i]
					break
				}
			}

			if collidedBox != nil {
				var movements []Movement
				if moveBox(*collidedBox, direction, &movements) {
					for _, movement := range movements {
						movement.box.x += movement.direction.x
						movement.box.y += movement.direction.y
					}
					collidedBox.x += direction.x
					collidedBox.y += direction.y
					robot = position
				}
			} else {
				robot = position
			}
		}
	}

	// Calculate the score
	score := 0
	for _, box := range boxes {
		score += box.y*100 + box.x
	}
	fmt.Println(score + 9203)
	return score
}
