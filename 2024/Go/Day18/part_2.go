package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Complex struct {
	Real, Imag int
}

func (c Complex) Add(other Complex) Complex {
	return Complex{c.Real + other.Real, c.Imag + other.Imag}
}

const (
	Width  = 70
	Height = 70
	Front  = 1024
)

func parseInput(filename string) ([]Complex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bytes []Complex
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid input format")
		}

		bytes = append(bytes, Complex{Real: x, Imag: y})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return bytes, nil
}

func bfs(bytes []Complex, cut int) int {
	start := Complex{Real: 0, Imag: 0}
	goal := Complex{Real: Width, Imag: Height}
	steps := 0

	walls := make(map[Complex]bool)
	for _, wall := range bytes[:cut] {
		walls[wall] = true
	}

	front := map[Complex]bool{start: true}
	seen := map[Complex]bool{start: true}

	directions := []Complex{
		{Real: 1, Imag: 0},
		{Real: -1, Imag: 0},
		{Real: 0, Imag: 1},
		{Real: 0, Imag: -1},
	}

	for len(front) > 0 {
		newFront := make(map[Complex]bool)
		steps++

		for pos := range front {
			for _, d := range directions {
				newPos := pos.Add(d)

				if newPos == goal {
					return steps
				}

				if newPos.Real < 0 || newPos.Real > Width || newPos.Imag < 0 || newPos.Imag > Height {
					continue
				}

				if walls[newPos] || seen[newPos] {
					continue
				}

				seen[newPos] = true
				newFront[newPos] = true
			}
		}

		front = newFront
	}
	return 0
}

func firstByteReachablePosition(bytes []Complex) {
	low, high := Front, len(bytes)

	for high-low > 1 {
		mid := (low + high) / 2
		if bfs(bytes, mid) > 0 {
			low = mid
		} else {
			high = mid
		}
	}

	fmt.Printf("%d,%d\n", bytes[low].Real, bytes[low].Imag)
}

func main() {
	bytes, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	firstByteReachablePosition(bytes)
}
