package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var dirs = map[string]string{"0": "R", "1": "D", "2": "L", "3": "U"}
var move = map[string]struct{ dx, dy int }{
	"R": {0, 1},
	"L": {0, -1},
	"U": {-1, 0},
	"D": {1, 0},
}

func CubicMetersOfLava() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data := strings.Split(strings.TrimSpace(string(content)), "\n")

	// Part 1
	x, y := 0, 0
	var borders []struct{ x, y int }
	for _, row := range data {
		parts := strings.Fields(row)
		d, m := parts[0], parts[1]
		dx, dy := move[d].dx, move[d].dy
		for i := 0; i < atoi(m); i++ {
			borders = append(borders, struct{ x, y int }{x, y})
			x, y = x+dx, y+dy
		}
	}

	// Shoelace formula
	area := 0
	for i := 0; i < len(borders)-1; i++ {
		x1, y1 := borders[i].x, borders[i].y
		x2, y2 := borders[i+1].x, borders[i+1].y
		area += x1*y2 - x2*y1
	}

	// Pick's theorem
	perimeter := len(borders)
	interiorArea := abs(area)/2 - perimeter/2 + 1
	fmt.Println("Part 1:", interiorArea+perimeter)

	// Part 2
	x, y = 0, 0
	borders = nil
	for _, row := range data {
		parts := strings.Fields(row)
		c := parts[2][1 : len(parts[2])-1]
		d, m := dirs[string(c[len(c)-1])], atoi(c[:len(c)-1])
		dx, dy := move[d].dx, move[d].dy
		for i := 0; i < m; i++ {
			borders = append(borders, struct{ x, y int }{x, y})
			x, y = x+dx, y+dy
		}
	}

	// Shoelace formula
	area = 0
	for i := 0; i < len(borders)-1; i++ {
		x1, y1 := borders[i].x, borders[i].y
		x2, y2 := borders[i+1].x, borders[i+1].y
		area += x1*y2 - x2*y1
	}

	// // Pick's theorem
	// perimeter = len(borders)
	// interiorArea = abs(area)/2 - perimeter/2 + 1
	// fmt.Println("Part 2:", interiorArea+perimeter)
}

// Helper function to convert string to integer
func atoi(s string) int {
	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}

// Helper function to get absolute value of an integer

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
