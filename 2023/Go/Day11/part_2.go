package main

import (
	"bufio"
	"fmt"
	"os"
)

func empty(mat [][]rune) (map[int]struct{}, map[int]struct{}) {
	size := len(mat)
	er := make(map[int]struct{})
	ec := make(map[int]struct{})

	for i := 0; i < size; i++ {
		r := true
		c := true
		for j := 0; j < size; j++ {
			r = r && mat[i][j] != '#'
			c = c && mat[j][i] != '#'
		}
		if r {
			er[i] = struct{}{}
		}
		if c {
			ec[i] = struct{}{}
		}
	}
	return er, ec
}

func findPoints(mat [][]rune) []point {
	var p []point

	for i, l := range mat {
		for j, c := range l {
			if c == '#' {
				p = append(p, point{i, j})
			}
		}
	}

	return p
}

type point struct {
	x, y int
}

func ShortestPathLengthInPairGalaxies() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var mat [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, []rune(line))
	}

	el, ec := empty(mat)
	points := findPoints(mat)
	distS := make([]int, 0)
	weight := 1000000

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			x := 0

			for k := min(p1.x, p2.x); k < max(p1.x, p2.x); k++ {
				if _, exists := el[k]; exists {
					x += weight
				} else {
					x++
				}
			}

			for k := min(p1.y, p2.y); k < max(p1.y, p2.y); k++ {
				if _, exists := ec[k]; exists {
					x += weight
				} else {
					x++
				}
			}

			distS = append(distS, x)
		}
	}

	total := 0
	for _, dist := range distS {
		total += dist
	}

	fmt.Println(total)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
