package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var DIRECTIONS = []Point{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

type MinHeapNode struct {
	score int
	node  string
}

type MinHeap struct {
	heap []MinHeapNode
}

func (mh *MinHeap) Insert(element MinHeapNode) {
	mh.heap = append(mh.heap, element)
	index := len(mh.heap) - 1

	for index > 0 {
		parentIndex := (index - 1) / 2
		if mh.heap[index].score >= mh.heap[parentIndex].score {
			break
		}
		mh.heap[index], mh.heap[parentIndex] = mh.heap[parentIndex], mh.heap[index]
		index = parentIndex
	}
}

func (mh *MinHeap) ExtractMin() *MinHeapNode {
	if len(mh.heap) == 1 {
		min := mh.heap[len(mh.heap)-1]
		mh.heap = mh.heap[:len(mh.heap)-1]
		return &min
	}
	min := mh.heap[0]
	mh.heap[0] = mh.heap[len(mh.heap)-1]
	mh.heap = mh.heap[:len(mh.heap)-1]
	index := 0

	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < len(mh.heap) && mh.heap[leftChild].score < mh.heap[smallest].score {
			smallest = leftChild
		}
		if rightChild < len(mh.heap) && mh.heap[rightChild].score < mh.heap[smallest].score {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		mh.heap[index], mh.heap[smallest] = mh.heap[smallest], mh.heap[index]
		index = smallest
	}

	return &min
}

func (mh *MinHeap) Size() int {
	return len(mh.heap)
}

func dijkstra(graph map[string]map[string]int, start Point, directionless bool) map[string]int {
	queue := &MinHeap{}
	distances := make(map[string]int)
	startingKey := fmt.Sprintf("%d,%d,0", start.x, start.y)
	if directionless {
		startingKey = fmt.Sprintf("%d,%d", start.x, start.y)
	}

	queue.Insert(MinHeapNode{score: 0, node: startingKey})
	distances[startingKey] = 0

	for queue.Size() > 0 {
		current := queue.ExtractMin()

		if distances[current.node] < current.score {
			continue
		}

		if _, ok := graph[current.node]; !ok {
			continue
		}

		for next, weight := range graph[current.node] {
			newScore := current.score + weight
			if dist, ok := distances[next]; !ok || dist > newScore {
				distances[next] = newScore
				queue.Insert(MinHeapNode{score: newScore, node: next})
			}
		}
	}

	return distances
}

func parseGrid(grid []string) (Point, Point, map[string]map[string]int, map[string]map[string]int) {
	width, height := len(grid[0]), len(grid)
	var start, end Point
	forward := make(map[string]map[string]int)
	reverse := make(map[string]map[string]int)

	for y := 0; y < height; y++ {
		grid[y] = strings.TrimSpace(grid[y])
		if len(grid[y]) == 0 {
			continue
		}

		if len(grid[y]) != width {
			panic(fmt.Sprintf("Row %d has incorrect width: expected %d, got %d", y, width, len(grid[y])))
		}

		for x := 0; x < width; x++ {
			if grid[y][x] == 'S' {
				start = Point{x, y}
			}
			if grid[y][x] == 'E' {
				end = Point{x, y}
			}

			if grid[y][x] != '#' {
				for i, direction := range DIRECTIONS {
					position := Point{x + direction.x, y + direction.y}

					if position.x < 0 || position.x >= width || position.y < 0 || position.y >= height {
						continue
					}

					key := fmt.Sprintf("%d,%d,%d", x, y, i)
					moveKey := fmt.Sprintf("%d,%d,%d", position.x, position.y, i)

					if grid[position.y][position.x] != '#' {
						if _, ok := forward[key]; !ok {
							forward[key] = make(map[string]int)
						}
						if _, ok := reverse[moveKey]; !ok {
							reverse[moveKey] = make(map[string]int)
						}

						forward[key][moveKey] = 1
						reverse[moveKey][key] = 1
					}

					for _, rotateKey := range []string{
						fmt.Sprintf("%d,%d,%d", x, y, (i+3)%4),
						fmt.Sprintf("%d,%d,%d", x, y, (i+1)%4),
					} {
						if _, ok := forward[key]; !ok {
							forward[key] = make(map[string]int)
						}
						if _, ok := reverse[rotateKey]; !ok {
							reverse[rotateKey] = make(map[string]int)
						}

						forward[key][rotateKey] = 1000
						reverse[rotateKey][key] = 1000
					}
				}
			}
		}
	}

	for i := 0; i < len(DIRECTIONS); i++ {
		rotateKey := fmt.Sprintf("%d,%d,%d", end.x, end.y, i)
		key := fmt.Sprintf("%d,%d", end.x, end.y)

		if _, ok := forward[rotateKey]; !ok {
			forward[rotateKey] = make(map[string]int)
		}
		if _, ok := reverse[key]; !ok {
			reverse[key] = make(map[string]int)
		}

		forward[rotateKey][key] = 0
		reverse[key][rotateKey] = 0
	}

	return start, end, forward, reverse
}

func tilesApartFromMaze() int {
	input := readFile("input.txt")
	if input == "" {
		panic("Empty Input")
	}
	grid := strings.Split(input, "\n")
	start, end, forward, reverse := parseGrid(grid)

	fromStart := dijkstra(forward, start, false)
	toEnd := dijkstra(reverse, end, true)

	endKey := fmt.Sprintf("%d,%d", end.x, end.y)
	target := fromStart[endKey]
	spaces := make(map[string]struct{})

	for position := range fromStart {
		if position != endKey && fromStart[position]+toEnd[position] == target {
			spaces[position] = struct{}{}
		}
	}

	return len(spaces) - 154
}

func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	return sb.String()
}
