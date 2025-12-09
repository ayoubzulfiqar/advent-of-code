package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Points struct {
	x, y uint32
}

func GetCoordinates(input string) []Points {
	var coords []Points
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
			coords = append(coords, Points{uint32(x), uint32(y)})
		}
	}

	return coords
}

func LargetsAreaWithGreenAndRedTiles(coords []Points) uint64 {
	var maxArea uint64 = 0

	for i, p1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			p2 := coords[j]

			x1, x2 := p1.x, p2.x
			y1, y2 := p1.y, p2.y
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			width := uint64(x2-x1) + 1
			height := uint64(y2-y1) + 1
			area := width * height

			if area <= maxArea {
				continue
			}

			if isValidRect(x1, x2, y1, y2, coords) {
				maxArea = area
			}
		}
	}

	return maxArea
}

func isValidRect(x1, x2, y1, y2 uint32, poly []Points) bool {
	mx := uint64(x1) + uint64(x2)
	my := uint64(y1) + uint64(y2)

	if !isPointInPoly(mx, my, poly) {
		return false
	}

	n := len(poly)
	for i := 0; i < n; i++ {
		u := poly[i]
		v := poly[(i+1)%n]

		if u.x == v.x {
			ex := u.x
			eyMin, eyMax := u.y, v.y
			if eyMin > eyMax {
				eyMin, eyMax = eyMax, eyMin
			}

			if ex > x1 && ex < x2 {
				overlapStart := maxU32(y1, eyMin)
				overlapEnd := minU32(y2, eyMax)
				if overlapStart < overlapEnd {
					return false
				}
			}
		} else {
			// Horizontal edge
			ey := u.y
			exMin, exMax := u.x, v.x
			if exMin > exMax {
				exMin, exMax = exMax, exMin
			}

			if ey > y1 && ey < y2 {
				overlapStart := maxU32(x1, exMin)
				overlapEnd := minU32(x2, exMax)
				if overlapStart < overlapEnd {
					return false
				}
			}
		}
	}

	return true
}

func isPointInPoly(x, y uint64, poly []Points) bool {

	n := len(poly)

	for i := 0; i < n; i++ {
		u := poly[i]
		v := poly[(i+1)%n]

		u0_2 := uint64(u.x) * 2
		u1_2 := uint64(u.y) * 2
		v0_2 := uint64(v.x) * 2
		v1_2 := uint64(v.y) * 2

		if u.x == v.x {
			if u0_2 == x {
				minY := minU64(u1_2, v1_2)
				maxY := maxU64(u1_2, v1_2)
				if y >= minY && y <= maxY {
					return true
				}
			}
		} else {
			if u1_2 == y {
				minX := minU64(u0_2, v0_2)
				maxX := maxU64(u0_2, v0_2)
				if x >= minX && x <= maxX {
					return true
				}
			}
		}
	}

	intersections := 0
	for i := 0; i < n; i++ {
		u := poly[i]
		v := poly[(i+1)%n]

		if u.x == v.x {
			minY := uint64(minU32(u.y, v.y)) * 2
			maxY := uint64(maxU32(u.y, v.y)) * 2
			ex := uint64(u.x) * 2

			if y >= minY && y < maxY && ex > x {
				intersections++
			}
		}
	}

	return intersections%2 == 1
}

func minU32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func maxU32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func minU64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func maxU64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// func main() {
//
// 	content, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	input := string(content)
// 	coords := GetCoordinates(input)
// 	part2 := LargetsAreaWithGreenAndRedTiles(coords)
// 	fmt.Printf("Part 2 Answer: %d\n", part2)
// }
