package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord [3]int

const MAX = 188_000_000

type Edge struct {
	a, b int
	dist int
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent, size}
}

func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x])
	}
	return dsu.parent[x]
}

func (dsu *DSU) Union(x, y int) {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)
	if rootX == rootY {
		return
	}
	if dsu.size[rootX] < dsu.size[rootY] {
		rootX, rootY = rootY, rootX
	}
	dsu.parent[rootY] = rootX
	dsu.size[rootX] += dsu.size[rootY]
}

func dist(a, b Coord) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	dz := a[2] - b[2]
	return dx*dx + dy*dy + dz*dz
}

func MultiplicationThreeLargestCircuits() {
	// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var coords []Coord
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		var coord Coord
		for i := 0; i < 3; i++ {
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				fmt.Printf("Error parsing number: %v\n", err)
				os.Exit(1)
			}
			coord[i] = val
		}
		coords = append(coords, coord)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Generate edges
	var edges []Edge
	n := len(coords)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := dist(coords[i], coords[j])
			if d < MAX {
				edges = append(edges, Edge{i, j, d})
			}
		}
	}

	// Sort edges by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	// Initialize DSU and process first 1000 edges
	dsu := NewDSU(n)
	limit := min(1000, len(edges))
	for i := 0; i < limit; i++ {
		dsu.Union(edges[i].a, edges[i].b)
	}

	// Count circuits
	circuitMap := make(map[int]int)
	for i := 0; i < n; i++ {
		root := dsu.Find(i)
		circuitMap[root]++
	}

	// Collect circuit sizes
	var circuits []int
	for _, size := range circuitMap {
		circuits = append(circuits, size)
	}

	// Sort in descending order
	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i] > circuits[j]
	})

	// Get product of top 3 circuits
	topCount := min(3, len(circuits))
	product := 1
	for i := 0; i < topCount; i++ {
		product *= circuits[i]
	}

	fmt.Println(product)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
