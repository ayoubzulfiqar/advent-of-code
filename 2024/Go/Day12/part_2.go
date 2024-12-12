package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Points struct {
	x, y int
}

type Edge struct {
	x, y, k     int
	orientation string
}

func totalFencingPrice() {
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
	neighborDict := make(map[Points]map[Points]struct{})
	edgeDict := make(map[Points]map[Edge]struct{})

	// Populate neighborDict and edgeDict
	dirs := []Points{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for i := 0; i < shapeX; i++ {
		for j := 0; j < shapeY; j++ {
			curPoint := Points{i, j}
			if _, exists := neighborDict[curPoint]; !exists {
				neighborDict[curPoint] = make(map[Points]struct{})
			}
			if _, exists := edgeDict[curPoint]; !exists {
				edgeDict[curPoint] = make(map[Edge]struct{})
			}
			for k, d := range dirs {
				ni, nj := i+d.x, j+d.y
				if ni >= 0 && ni < shapeX && nj >= 0 && nj < shapeY && grid[i][j] == grid[ni][nj] {
					neighborDict[curPoint][Points{ni, nj}] = struct{}{}
				} else {
					orientation := "vh"[int(math.Abs(float64(d.x))):][:1]
					edgeDict[curPoint][Edge{i + max(d.x, 0), j + max(d.y, 0), k, orientation}] = struct{}{}
				}
			}
		}
	}

	// Helper to get region
	getRegion := func(point Points) map[Points]struct{} {
		region := make(map[Points]struct{})
		remaining := map[Points]struct{}{point: {}}

		for len(remaining) > 0 {
			var curPoint Points
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
	remainingPoints := make(map[Points]struct{})
	for i := 0; i < shapeX; i++ {
		for j := 0; j < shapeY; j++ {
			remainingPoints[Points{i, j}] = struct{}{}
		}
	}

	var regions []map[Points]struct{}
	for len(remainingPoints) > 0 {
		var startPoint Points
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

	// Helper to split edges
	splitWhen := func(edges []Edge, predicate func(Edge, Edge) bool) [][]Edge {
		var result [][]Edge
		var group []Edge
		for i, edge := range edges {
			if i > 0 && predicate(edges[i-1], edge) {
				result = append(result, group)
				group = nil
			}
			group = append(group, edge)
		}
		if len(group) > 0 {
			result = append(result, group)
		}
		return result
	}

	// Calculate number of sides
	numSides := func(region map[Points]struct{}) int {
		edges := make(map[Edge]struct{})
		for point := range region {
			for edge := range edgeDict[point] {
				edges[edge] = struct{}{}
			}
		}

		var horizontals, verticals []Edge
		for edge := range edges {
			if edge.orientation == "h" {
				horizontals = append(horizontals, edge)
			} else {
				verticals = append(verticals, edge)
			}
		}

		sort.Slice(horizontals, func(i, j int) bool {
			return horizontals[i].x < horizontals[j].x || (horizontals[i].x == horizontals[j].x && horizontals[i].y < horizontals[j].y)
		})
		sort.Slice(verticals, func(i, j int) bool {
			return verticals[i].y < verticals[j].y || (verticals[i].y == verticals[j].y && verticals[i].x < verticals[j].x)
		})

		horizontalsGroups := splitWhen(horizontals, func(x, y Edge) bool {
			return !(x.x == y.x && y.y-x.y == 1 && y.k == x.k)
		})
		verticalsGroups := splitWhen(verticals, func(x, y Edge) bool {
			return !(x.y == y.y && y.x-x.x == 1 && y.k == x.k)
		})

		// numHorizontalSides := len(horizontalsGroups)
		// numVerticalSides := len(verticalsGroups)

		return len(horizontalsGroups) + len(verticalsGroups)
	}

	// Calculate area
	area := func(region map[Points]struct{}) int {
		return len(region)
	}

	// Compute the answer
	answer := 0
	for _, region := range regions {
		answer += numSides(region) * area(region)
	}

	fmt.Println(answer)
}
