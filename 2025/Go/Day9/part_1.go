package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Point struct {
	x, y uint32
}

func GetCoords(input string) []Point {
	var coords []Point
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		x, err1 := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 32)
		y, err2 := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 32)

		if err1 == nil && err2 == nil {
			coords = append(coords, Point{uint32(x), uint32(y)})
		}
	}

	return coords
}

func LargestAreaForRectangle(coords []Point) uint64 {
	var maxArea uint64 = 0

	for i, p1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			p2 := coords[j]

			var width, height uint64
			if p1.x > p2.x {
				width = uint64(p1.x-p2.x) + 1
			} else {
				width = uint64(p2.x-p1.x) + 1
			}

			if p1.y > p2.y {
				height = uint64(p1.y-p2.y) + 1
			} else {
				height = uint64(p2.y-p1.y) + 1
			}

			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// func main() {
// 	// Read input from file
// 	content, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	input := string(content)

// 	coords := GetCoords(input)
// 	part1 := LargestAreaForRectangle(coords)
// 	fmt.Printf("Part 1 Answer: %d\n", part1)

// }
