package main

import (
	"bufio"
	"fmt"
	"os"
)

func addLines(mat [][]rune) [][]rune {
	size := len(mat)
	var nm [][]rune

	for i := 0; i < size; i++ {
		r := true
		for j := 0; j < size; j++ {
			r = r && mat[i][j] != '#'
		}
		nm = append(nm, mat[i])
		if r {
			nm = append(nm, make([]rune, size))
		}
	}

	return nm
}

func addColumn(mat [][]rune) [][]rune {
	size := len(mat)
	var nm [][]rune

	for i := 0; i < size; i++ {
		nm = append(nm, make([]rune, 0, len(mat[0])))
	}

	for i := 0; i < len(mat[0]); i++ {
		c := true
		for j := 0; j < size; j++ {
			c = c && mat[j][i] != '#'
			nm[j] = append(nm[j], mat[j][i])
		}
		if c {
			for j := 0; j < size; j++ {
				nm[j] = append(nm[j], '.')
			}
		}
	}

	return nm
}

// func findPoints(mat [][]rune) []point {
// 	var p []point

// 	for i, l := range mat {
// 		for j, c := range l {
// 			if c == '#' {
// 				p = append(p, point{i, j})
// 			}
// 		}
// 	}

// 	return p
// }

// type point struct {
// 	x, y int
// }

func ShortestPathGalaxyLength() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var mat [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, []rune(line))
	}

	mat = addLines(mat)
	mat = addColumn(mat)
	points := findPoints(mat)

	dists := make([]int, 0)

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			x := abs(p1.x-p2.x) + abs(p1.y-p2.y)
			dists = append(dists, x)
		}
	}

	total := 0
	for _, dist := range dists {
		total += dist
	}

	fmt.Println(total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
