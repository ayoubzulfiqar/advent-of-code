package main

import (
	"bufio"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

var (
	DOWN   = Point{0, 1}
	UP     = Point{0, -1}
	ORIGIN = Point{0, 0}
)

func parse() string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return strings.TrimSpace(sb.String())
}

func parseGrid(block string) [][]rune {
	lines := strings.Split(block, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func uniqueLockANDKey() int {
	data := parse()
	locks := []int{}
	keys := []int{}
	result := 0

	blocks := strings.Split(data, "\n\n")
	for _, block := range blocks {
		grid := parseGrid(block)
		heights := 0

		if grid[ORIGIN.y][ORIGIN.x] == '#' {
			for x := 0; x < 5; x++ {
				position := Point{x, 1}
				for position.y < len(grid) && grid[position.y][position.x] == '#' {
					position = position.Add(DOWN)
				}
				heights = (heights << 4) + (position.y - 1)
			}
			locks = append(locks, heights)
		} else {
			for x := 0; x < 5; x++ {
				position := Point{x, 5}
				for position.y >= 0 && grid[position.y][position.x] == '#' {
					position = position.Add(UP)
				}
				heights = (heights << 4) + (5 - position.y)
			}
			keys = append(keys, heights)
		}
	}

	for _, lock := range locks {
		for _, key := range keys {
			if (lock+key+0x22222)&0x88888 == 0 {
				result++
			}
		}
	}

	return result
}

// func main() {
// 	fmt.Println(uniqueLockANDKey())
// }
