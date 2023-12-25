package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	px, py, pz, vx, vy, vz int
}

func getLinesIntersection(p1x, v1x, p1y, v1y, p2x, v2x, p2y, v2y int) map[string]float64 {
	x1, x2, y1, y2 := float64(p1x), float64(p1x+100000000000000*v1x), float64(p1y), float64(p1y+100000000000000*v1y)
	x3, x4, y3, y4 := float64(p2x), float64(p2x+100000000000000*v2x), float64(p2y), float64(p2y+100000000000000*v2y)

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		return nil
	}

	x := ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / denominator
	y := ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / denominator

	return map[string]float64{"x": x, "y": y}
}

func getTime(s, v, p int) float64 {
	return float64(p-s) / float64(v)
}

func InteractionWithinTestArea() {
	hailstones := make([]Hailstone, 0)
	min := 200000000000000
	max := 400000000000000
	count := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " @ ")
		position := strings.Split(parts[0], ", ")
		velocity := strings.Split(parts[1], ", ")

		px, _ := strconv.Atoi(position[0])
		py, _ := strconv.Atoi(position[1])
		pz, _ := strconv.Atoi(position[2])
		vx, _ := strconv.Atoi(velocity[0])
		vy, _ := strconv.Atoi(velocity[1])
		vz, _ := strconv.Atoi(velocity[2])

		hailstones = append(hailstones, Hailstone{px, py, pz, vx, vy, vz})
	}

	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[i]
			h2 := hailstones[j]

			intersection := getLinesIntersection(h1.px, h1.vx, h1.py, h1.vy, h2.px, h2.vx, h2.py, h2.vy)

			if intersection == nil {
				continue
			}

			if intersection["x"] < float64(min) || intersection["x"] > float64(max) ||
				intersection["y"] < float64(min) || intersection["y"] > float64(max) {
				continue
			}

			timeH1 := getTime(h1.px, h1.vx, int(intersection["x"]))
			timeH2 := getTime(h2.px, h2.vx, int(intersection["x"]))

			if timeH1 < 0 || timeH2 < 0 {
				continue
			}

			count++
		}
	}

	fmt.Println(count)
}
