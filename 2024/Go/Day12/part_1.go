package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func totalPriceRegion() {
	// Read the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	shapeX, shapeY := len(grid), len(grid[0])
	neighborDict := make(map[Point]map[Point]struct{})

	// Populate neighborDict
	dirs := []Point{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for i := 0; i < shapeX; i++ {
		for j := 0; j < shapeY; j++ {
			curPoint := Point{i, j}
			if _, exists := neighborDict[curPoint]; !exists {
				neighborDict[curPoint] = make(map[Point]struct{})
			}
			for _, d := range dirs {
				ni, nj := i+d.x, j+d.y
				if ni >= 0 && ni < shapeX && nj >= 0 && nj < shapeY {
					if grid[i][j] == grid[ni][nj] {
						neighborDict[curPoint][Point{ni, nj}] = struct{}{}
					}
				}
			}
		}
	}

	// Helper to get region
	getRegion := func(point Point) map[Point]struct{} {
		region := make(map[Point]struct{})
		remaining := map[Point]struct{}{point: {}}

		for len(remaining) > 0 {
			var curPoint Point
			for p := range remaining {
				curPoint = p
				break
			}
			delete(remaining, curPoint)
			region[curPoint] = struct{}{}

			for neighbor := range neighborDict[curPoint] {
				if _, inRegion := region[neighbor]; !inRegion {
					remaining[neighbor] = struct{}{}
				}
			}
		}

		return region
	}

	// Find all regions
	remainingPoints := make(map[Point]struct{})
	for i := 0; i < shapeX; i++ {
		for j := 0; j < shapeY; j++ {
			remainingPoints[Point{i, j}] = struct{}{}
		}
	}

	var regions []map[Point]struct{}
	for len(remainingPoints) > 0 {
		var startPoint Point
		for p := range remainingPoints {
			startPoint = p
			break
		}
		region := getRegion(startPoint)
		regions = append(regions, region)

		for p := range region {
			delete(remainingPoints, p)
		}
	}

	// Calculate perimeter
	perimeter := func(region map[Point]struct{}) int {
		p := 0
		for point := range region {
			p += 4 - len(neighborDict[point])
		}
		return p
	}

	// Calculate area
	area := func(region map[Point]struct{}) int {
		return len(region)
	}

	// Compute the answer
	answer := 0
	for _, region := range regions {
		answer += perimeter(region) * area(region)
	}

	fmt.Println(answer)
}
