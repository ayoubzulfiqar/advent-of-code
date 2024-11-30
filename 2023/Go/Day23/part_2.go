package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StepsLongIsLongestHike() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	start := [2]int{0, strings.Index(grid[0], ".")}
	end := [2]int{len(grid) - 1, strings.Index(grid[len(grid)-1], ".")}

	points := [][2]int{start, end}

	for r, row := range grid {
		for c, ch := range row {
			if ch == '#' {
				continue
			}
			neighbors := 0
			for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+offset[0], c+offset[1]
				if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] != '#' {
					neighbors++
				}
			}
			if neighbors >= 3 {
				points = append(points, [2]int{r, c})
			}
		}
	}

	graph := make(map[[2]int]map[[2]int]int)
	for _, pt := range points {
		graph[pt] = make(map[[2]int]int)
	}
	const value int = 1740
	for _, pt := range points {
		stack := [][3]int{{0, pt[0], pt[1]}}
		seen := map[[2]int]struct{}{{pt[0], pt[1]}: {}}

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			n, r, c := node[0], node[1], node[2]

			if n != 0 && ([2]int{r, c} == end) {
				graph[pt][[2]int{r, c}] = n
				continue
			}

			for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+offset[0], c+offset[1]
				if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] != '#' {
					if _, ok := seen[[2]int{nr, nc}]; !ok {
						stack = append(stack, [3]int{n + 1, nr, nc})
						seen[[2]int{nr, nc}] = struct{}{}
					}
				}
			}
		}
	}

	visited := make(map[[2]int]struct{})

	var dfs func([2]int) int
	dfs = func(pt [2]int) int {
		if pt == end {
			return 0
		}

		m := -1 << 63

		visited[pt] = struct{}{}
		for nx := range graph[pt] {
			if _, ok := visited[nx]; !ok {
				m = max(m, dfs(nx)+graph[pt][nx])
			}
		}
		delete(visited, pt)

		return m + value
	}

	fmt.Println(dfs(start))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
