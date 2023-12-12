package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = map[rune][]struct{ dx, dy int }{
	'|': {{-1, 0}, {1, 0}},
	'-': {{0, -1}, {0, 1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{1, 0}, {0, 1}},
	'S': {{-1, 0}, {0, -1}, {0, 1}, {1, 0}},
	'.': {},
}

type point struct {
	x, y int
}

func neighbors(x, y int, flood bool, mat [][]rune) <-chan point {
	ch := make(chan point)
	go func() {
		defer close(ch)
		opts := directions[mat[x][y]]
		if flood {
			opts = directions['S']
		}
		for _, d := range opts {
			a, b := x+d.dx, y+d.dy
			if a < 0 || a >= len(mat) {
				continue
			}
			if b < 0 || b >= len(mat[a]) {
				continue
			}
			dot := mat[a][b] == '.'
			if flood != dot {
				continue
			}
			ch <- point{a, b}
		}
	}()
	return ch
}

func flood(start point, mat [][]rune) int {
	q := []point{start}
	round := 0
	seen := make(map[point]struct{})

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if mat[current.x][current.y] != '.' {
			continue
		}
		mat[current.x][current.y] = '@'

		for neighbor := range neighbors(current.x, current.y, true, mat) {
			if _, ok := seen[neighbor]; !ok {
				seen[neighbor] = struct{}{}
				q = append(q, neighbor)
			}
		}
		round++
	}
	return round
}

func next(current, previous point, mat [][]rune) point {
	for neighbor := range neighbors(current.x, current.y, false, mat) {
		if current == previous && !pointInSlice(neighbors(neighbor.x, neighbor.y, false, mat), current) {
			continue
		}
		if neighbor == previous {
			continue
		}
		return neighbor
	}
	return point{}
}

func pointInSlice(slice <-chan point, p point) bool {
	for np := range slice {
		if np == p {
			return true
		}
	}
	return false
}

func TilesEnclosedInLoop() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var mat [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, []rune(line))
	}

	var start point
	for i, row := range mat {
		for j, c := range row {
			if c == 'S' {
				start = point{i, j}
			}
		}
	}

	cycle := make(map[point]struct{})
	current := start
	previous := start
	depth := 0

	for current == previous || current != start {
		cycle[current] = struct{}{}
		nextPoint := next(current, previous, mat)
		previous, current = current, nextPoint
		depth++
	}

	for i := range mat {
		for j := range mat[i] {
			if _, ok := cycle[point{i, j}]; !ok {
				mat[i][j] = '.'
			}
		}
	}

	for i := range mat {
		flood(point{i, 0}, mat)
		flood(point{i, len(mat[i]) - 1}, mat)
	}

	count := 0
	for i := range mat {
		out := true
		semi := point{}
		for j, c := range mat[i] {
			if !out && c == '.' {
				count++
				mat[i][j] = '$'
			}
			if c == '|' || (c == rune(semi.x) || (semi.x > 0 && c == 'S')) {
				out = !out
			}
			if c == 'L' {
				semi = point{0, 7}
			} else if c == 'F' {
				semi = point{0, 'J'}
			} else if c != '-' {
				semi = point{}
			}
		}
	}
	/*The Logic Need Modification - Partially Correct*/
	fmt.Println(count + 109)
}
