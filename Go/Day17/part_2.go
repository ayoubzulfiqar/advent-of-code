package main

import (
	"container/heap"
	"log"
	"os"
	"slices"
	"strings"
)

type direction struct {
	row, col int
}

var (
	left  = direction{row: 0, col: -1}
	right = direction{row: 0, col: 1}
	up    = direction{row: -1, col: 0}
	down  = direction{row: 1, col: 0}
)

var rotations = map[direction][]direction{
	left:  {up, down},
	right: {up, down},
	up:    {left, right},
	down:  {left, right},
}

var reverse = map[direction]direction{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type state struct {
	row   int
	col   int
	dir   direction
	moves int
}

type item struct {
	heatLoss int
	state    state
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func LeastHeatLossIncur() int {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(b)
	grid := parse(input)
	return dijkstra(grid, 10, 4)
}

func dijkstra(grid []string, maxConsecutive, movesNeededBeforeTurn int) int {
	m := len(grid)
	n := len(grid[0])
	startRight := state{row: 0, col: 0, dir: right, moves: 0}
	startDown := state{row: 0, col: 0, dir: down, moves: 0}
	pq := priorityQueue{
		&item{heatLoss: 0, state: startRight, index: 0},
		&item{heatLoss: 0, state: startDown, index: 1},
	}

	minCost := map[state]int{startRight: 0, startDown: 0}
	heap.Init(&pq)

	for len(pq) > 0 {
		current := heap.Pop(&pq).(*item)
		if minCost[current.state] < current.heatLoss {
			continue
		}

		if current.state.row == m-1 && current.state.col == n-1 && current.state.moves >= movesNeededBeforeTurn {
			return current.heatLoss
		}

		for _, dir := range [4]direction{left, right, up, down} {
			if current.state.moves == maxConsecutive && !slices.Contains(rotations[current.state.dir], dir) || dir == reverse[current.state.dir] {
				continue
			}

			ni, nj := current.state.row+dir.row, current.state.col+dir.col
			nextMoves := current.state.moves

			if current.state.moves < movesNeededBeforeTurn {
				if dir != current.state.dir {
					continue
				}
				nextMoves += 1
			} else {
				if dir != current.state.dir {
					nextMoves = 1
				} else {
					nextMoves = nextMoves%maxConsecutive + 1
				}
			}

			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}

			nextState := state{row: ni, col: nj, moves: nextMoves, dir: dir}
			nextHeatLoss := int(rune(grid[ni][nj]) - '0')
			if _, ok := minCost[nextState]; ok && minCost[nextState] <= current.heatLoss+nextHeatLoss {
				continue
			}

			minCost[nextState] = current.heatLoss + nextHeatLoss
			heap.Push(&pq, &item{heatLoss: current.heatLoss + nextHeatLoss, state: nextState})
		}
	}

	return -1
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
