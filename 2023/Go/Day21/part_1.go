package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	R int
	C int
	S int
}

func SixtyFourElfSteps() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	sr, sc := 0, 0
	for r, row := range grid {
		for c, ch := range row {
			if ch == 'S' {
				sr, sc = r, c
				break
			}
		}
	}

	ans := make(map[Point]bool)
	seen := make(map[Point]bool)
	q := []Point{{sr, sc, 64}}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		r, c, s := current.R, current.C, current.S

		if s%2 == 0 {
			ans[current] = true
		}
		if s == 0 {
			continue
		}

		moves := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, move := range moves {
			nr, nc := r+move[0], c+move[1]

			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '#' || seen[Point{nr, nc, 0}] {
				continue
			}

			seen[Point{nr, nc, 0}] = true
			q = append(q, Point{nr, nc, s - 1})
		}
	}

	fmt.Println(len(ans) - 1)
}
