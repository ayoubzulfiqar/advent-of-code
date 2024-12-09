package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func uniqueLocationsInBounds() int {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the grid
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	width, height := len(grid[0]), len(grid)
	antennas := make(map[rune][]map[string]int)

	// Find all antennas with the same frequency (a-z, A-Z, 0-9)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] != '.' {
				frequency := rune(grid[y][x])
				if _, exists := antennas[frequency]; !exists {
					antennas[frequency] = []map[string]int{}
				}
				antennas[frequency] = append(antennas[frequency], map[string]int{"x": x, "y": y})
			}
		}
	}

	// Count all unique antiNodes from each pair of antennas
	antiNodes := make(map[string]struct{})
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}

				dx := positions[j]["x"] - positions[i]["x"]
				dy := positions[j]["y"] - positions[i]["y"]

				// Try all positions in the back of a pair of antennas
				for k := -50; k <= 50; k++ {
					antinodeX := positions[i]["x"] + dx*k
					antinodeY := positions[i]["y"] + dy*k

					// Do bounds checking
					if antinodeX >= 0 && antinodeX < width && antinodeY >= 0 && antinodeY < height {
						key := fmt.Sprintf("%d,%d", antinodeX, antinodeY)
						antiNodes[key] = struct{}{}
					}
				}
			}
		}
	}

	// Output the result
	fmt.Println(len(antiNodes))
	return len(antiNodes)
}
