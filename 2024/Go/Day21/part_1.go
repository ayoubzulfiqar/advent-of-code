package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var BFS_DIRECTIONS = map[string]struct {
	x, y int
}{
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
}

var KEYPAD = map[string]struct {
	x, y int
}{
	"7": {0, 0},
	"8": {1, 0},
	"9": {2, 0},
	"4": {0, 1},
	"5": {1, 1},
	"6": {2, 1},
	"1": {0, 2},
	"2": {1, 2},
	"3": {2, 2},
	"X": {0, 3},
	"0": {1, 3},
	"A": {2, 3},
}

var DIRECTIONS = map[string]struct {
	x, y int
}{
	"X": {0, 0},
	"^": {1, 0},
	"A": {2, 0},
	"<": {0, 1},
	"v": {1, 1},
	">": {2, 1},
}

func readInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// generate all paths from one button to another
func getCommand(input map[string]struct{ x, y int }, start, end string) []string {
	queue := []struct {
		x, y int
		path string
	}{{input[start].x, input[start].y, ""}}
	distances := make(map[string]int)
	allPaths := []string{}

	if start == end {
		return []string{"A"}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// find all paths
		if current.x == input[end].x && current.y == input[end].y {
			allPaths = append(allPaths, current.path+"A")
		}
		if distances[fmt.Sprintf("%d,%d", current.x, current.y)] != 0 && distances[fmt.Sprintf("%d,%d", current.x, current.y)] < len(current.path) {
			continue
		}

		for direction, vector := range BFS_DIRECTIONS {
			position := struct{ x, y int }{current.x + vector.x, current.y + vector.y}

			// don't allow traversal into the blank areas
			if input["X"].x == position.x && input["X"].y == position.y {
				continue
			}

			// only traverse if there is a button to hit
			for _, button := range input {
				if button.x == position.x && button.y == position.y {
					newPath := current.path + direction
					if distances[fmt.Sprintf("%d,%d", position.x, position.y)] == 0 || distances[fmt.Sprintf("%d,%d", position.x, position.y)] >= len(newPath) {
						queue = append(queue, struct {
							x, y int
							path string
						}{position.x, position.y, newPath})
						distances[fmt.Sprintf("%d,%d", position.x, position.y)] = len(newPath)
					}
				}
			}
		}
	}

	// sort from smallest to largest paths
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return allPaths
}

// find the smallest amount of button presses, given the robot and code to enter
func getKeyPresses(input map[string]struct{ x, y int }, code string, robot int, memo map[string]int) int {
	key := fmt.Sprintf("%s,%d", code, robot)
	if val, exists := memo[key]; exists {
		return val
	}

	current := "A"
	length := 0
	for i := 0; i < len(code); i++ {
		// find the smallest move for each transition
		moves := getCommand(input, current, string(code[i]))
		if robot == 0 {
			length += len(moves[0])
		} else {
			minLength := math.MaxInt
			for _, move := range moves {
				minLength = int(math.Min(float64(minLength), float64(getKeyPresses(DIRECTIONS, move, robot-1, memo))))
			}
			length += minLength
		}
		current = string(code[i])
	}

	memo[key] = length
	return length
}

// Part 1 solution
func sumOfFiveCodeComplexities(input string) int {
	keycodes := strings.Split(input, "\n")
	memo := make(map[string]int)
	total := 0

	for _, code := range keycodes {
		code = strings.TrimSpace(code)
		numerical := 0
		for _, char := range code {
			if char >= '0' && char <= '9' {
				numerical = numerical*10 + int(char-'0')
			}
		}
		total += numerical * getKeyPresses(KEYPAD, code, 2, memo)
	}

	return total
}

// func main() {
// 	// Example of how to use the functions
// 	// input := "540A\n839A\n682A\n826A\n974A"

// 	input := readingInput("input.txt")
// 	fmt.Println("Part 1:", sumOfFiveCodeComplexities(input))
// }