package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vertex struct {
	coordX, coordY, coordZ int
}

func (v Vertex) separation(w Vertex) float64 {
	deltaX := float64(v.coordX - w.coordX)
	deltaY := float64(v.coordY - w.coordY)
	deltaZ := float64(v.coordZ - w.coordZ)
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)
}

type EdgeInfo struct {
	idxA, idxB int
	distance   float64
}

type UnionStructure struct {
	ancestor []int
	level    []int
	groups   int
}

func CreateUnionStructure(total int) *UnionStructure {
	par := make([]int, total)
	height := make([]int, total)
	for k := 0; k < total; k++ {
		par[k] = k
		height[k] = 0
	}
	return &UnionStructure{par, height, total}
}

func (us *UnionStructure) Locate(pos int) int {
	if us.ancestor[pos] != pos {
		us.ancestor[pos] = us.Locate(us.ancestor[pos])
	}
	return us.ancestor[pos]
}

func (us *UnionStructure) Combine(x, y int) bool {
	rootX := us.Locate(x)
	rootY := us.Locate(y)

	if rootX == rootY {
		return false
	}

	if us.level[rootX] < us.level[rootY] {
		rootX, rootY = rootY, rootX
	}

	us.ancestor[rootY] = rootX
	if us.level[rootX] == us.level[rootY] {
		us.level[rootX]++
	}

	us.groups--
	return true
}

// func magnitudeDifference(a, b int) int {
// 	if a > b {
// 		return a - b
// 	}
// 	return b - a
// }

func XCoordinatesJuntionMultiplicatonBoxs() {
	dataFile, openErr := os.Open("input.txt")
	if openErr != nil {
		fmt.Printf("Cannot access data file: %v\n", openErr)
		os.Exit(1)
	}
	defer dataFile.Close()

	var locations []Vertex
	lineReader := bufio.NewScanner(dataFile)
	for lineReader.Scan() {
		textRow := lineReader.Text()
		if textRow == "" {
			continue
		}

		elements := strings.Split(textRow, ",")
		if len(elements) != 3 {
			continue
		}

		var node Vertex
		for pos, elem := range elements {
			numVal, parseErr := strconv.Atoi(strings.TrimSpace(elem))
			if parseErr != nil {
				fmt.Printf("Failed to convert value: %v\n", parseErr)
				os.Exit(1)
			}
			switch pos {
			case 0:
				node.coordX = numVal
			case 1:
				node.coordY = numVal
			case 2:
				node.coordZ = numVal
			}
		}
		locations = append(locations, node)
	}

	if readErr := lineReader.Err(); readErr != nil {
		fmt.Printf("Error during file reading: %v\n", readErr)
		os.Exit(1)
	}

	var connections []EdgeInfo
	vertexCount := len(locations)

	for firstIdx := 0; firstIdx < vertexCount; firstIdx++ {
		for secondIdx := firstIdx + 1; secondIdx < vertexCount; secondIdx++ {
			span := locations[firstIdx].separation(locations[secondIdx])
			connections = append(connections, EdgeInfo{firstIdx, secondIdx, span})
		}
	}

	sort.Slice(connections, func(m, n int) bool {
		return connections[m].distance < connections[n].distance
	})

	unionSet := CreateUnionStructure(vertexCount)

	for _, link := range connections {
		if unionSet.Combine(link.idxA, link.idxB) {
			if unionSet.groups == 1 {
				finalProduct := locations[link.idxA].coordX * locations[link.idxB].coordX
				fmt.Println(finalProduct)
				return
			}
		}
	}
}
