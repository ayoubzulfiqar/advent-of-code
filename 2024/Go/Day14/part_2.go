package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bots struct {
	x  int
	y  int
	dx int
	dy int
}

const (
	maxX = 101
	maxY = 103
)

func robotsToEasterEgg() {
	var bots []Bots

	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the lines from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var bs Bots
		// Parse the line into Bot structure (updated format: p=<x,y> v=<dx,dy>)
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &bs.x, &bs.y, &bs.dx, &bs.dy)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		bots = append(bots, bs)
	}

	// Loop for 100,000 iterations
	for i := 0; i < 1e5; i++ {
		// Initialize the grid and fill with '.'
		var grid [maxY][maxX]rune
		for y := range grid {
			for x := range grid[y] {
				grid[y][x] = '.'
			}
		}

		distinct := true

		// Update bot positions and check for collisions
		for j := 0; j < len(bots); j++ {
			bots[j].x += bots[j].dx
			bots[j].y += bots[j].dy

			// Wrap around the edges
			if bots[j].x < 0 {
				bots[j].x += maxX
			} else if bots[j].x >= maxX {
				bots[j].x -= maxX
			}

			if bots[j].y < 0 {
				bots[j].y += maxY
			} else if bots[j].y >= maxY {
				bots[j].y -= maxY
			}

			// Mark the grid cell
			if grid[bots[j].y][bots[j].x] == '.' {
				grid[bots[j].y][bots[j].x] = '#'
			} else {
				distinct = false
			}
		}

		// If all bots were distinct, print the grid
		if distinct {
			// in my case first Iter was the answer
			fmt.Printf("\nIter: %d\n", i+1)
			// for _, row := range grid {
			// 	for _, c := range row {
			// 		fmt.Print(string(c))
			// 	}
			// 	fmt.Println()
			// }
		}
	}
}
