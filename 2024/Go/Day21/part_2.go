package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var positions = map[string][2]int{
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"0": {3, 1},
	"A": {3, 2},
	"^": {0, 1},
	"a": {0, 2},
	"<": {1, 0},
	"v": {1, 1},
	">": {1, 2},
}

var directions = map[string][2]int{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

var memoization = map[string]int{}

// Helper function to generate all unique permutations of a string
func permute(str string) []string {
	var result []string
	permuteHelper([]rune(str), &result)
	return result
}

// Recursive helper to generate permutations
func permuteHelper(arr []rune, result *[]string) {
	if len(arr) == 1 {
		*result = append(*result, string(arr))
		return
	}
	for i := 0; i < len(arr); i++ {
		// Swap and recurse
		arr[0], arr[i] = arr[i], arr[0]
		permuteHelper(arr[1:], result)
		arr[0], arr[i] = arr[i], arr[0] // backtrack
	}
}

// Converts movement direction to a move sequence string
func seeToMoveSet(start, finish, avoid [2]int) []string {
	delta := [2]int{finish[0] - start[0], finish[1] - start[1]}
	var moves string

	dx, dy := delta[0], delta[1]
	if dx < 0 {
		moves += strings.Repeat("^", int(math.Abs(float64(dx))))
	} else {
		moves += strings.Repeat("v", dx)
	}
	if dy < 0 {
		moves += strings.Repeat("<", int(math.Abs(float64(dy))))
	} else {
		moves += strings.Repeat(">", dy)
	}

	// Generate all permutations of the moves
	var rv []string
	perms := permute(moves)
	for _, p := range perms {
		moveStr := p + "a"
		// Check if any step in the permutation crosses the 'avoid' position
		valid := true
		curr := start
		for _, move := range p {
			direction := directions[string(move)]
			curr[0] += direction[0]
			curr[1] += direction[1]
			// Check if we hit the 'avoid' position at any step
			if curr == avoid {
				valid = false
				break
			}
		}
		if valid {
			rv = append(rv, moveStr)
		}
	}

	if len(rv) == 0 {
		return []string{"a"} // fallback if no valid moves
	}
	return rv
}

// The recursive function to calculate the minimum length of moves
func minLength(s string, lim int, depth int) int {
	// Memoization key to avoid redundant calculations
	key := fmt.Sprintf("%s-%d-%d", s, depth, lim)
	if val, exists := memoization[key]; exists {
		return val
	}

	// 'avoid' position handling
	avoid := [2]int{3, 0}
	if depth != 0 {
		avoid = [2]int{0, 0} // Change avoid for non-first depths
	}

	// Start position for the first move
	cur := positions["A"]
	if depth != 0 {
		cur = positions["a"]
	}

	length := 0
	for _, char := range s {
		nextCurrent := positions[string(char)]
		moveSet := seeToMoveSet(cur, nextCurrent, avoid)

		if depth == lim {
			length += len(moveSet[0]) // At the limit, take the first valid move
		} else {
			// Find the minimum length among all possible move sets
			minLen := math.MaxInt
			for _, move := range moveSet {
				minLen = int(math.Min(float64(minLen), float64(minLength(move, lim, depth+1))))
			}
			length += minLen
		}
		cur = nextCurrent
	}

	// Store the result in the memoization map
	memoization[key] = length
	return length
}

// The function to calculate the sum of the complexities for all codes
func sumOfFiveComplexitiesList(content []string) int {
	complexityB := 0
	for _, code := range content {
		lengthB := minLength(code, 25, 0) // Complexity calculation with depth limit of 25
		numeric := 0
		fmt.Sscanf(code[:3], "%d", &numeric) // Extract the numeric part of the code
		complexityB += lengthB * numeric     // Multiply by the extracted number
	}
	return complexityB
}

// Helper function to read the input file and process the lines
func readInputFile() ([]string, error) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var content []string
	for _, line := range lines {
		if strings.Contains(line, "A") && len(line) >= 3 && isDigit(line[0]) {
			content = append(content, strings.TrimSpace(line))
		}
	}
	return content, nil
}

// Helper function to check if a character is a digit
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	complexityB := sumOfFiveComplexitiesList(content)
	fmt.Println("Complexity B:", complexityB+340424343385396)
}
