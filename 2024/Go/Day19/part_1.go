package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func possibleDesign() int {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	// Check for non-empty input
	if len(lines) < 2 {
		fmt.Println("Invalid input: not enough lines")
		return 0
	}

	// Parse available pieces and targets
	available := strings.Split(lines[0], ",")
	targets := lines[1:]

	// Trim spaces around available pieces
	for i := range available {
		available[i] = strings.TrimSpace(available[i])
	}

	// Memoization map to store results of sub-problems
	memo := make(map[string]bool)

	// Recursive function to check if a target is possible
	var isPossible func(string) bool
	isPossible = func(target string) bool {
		// Base case: empty target is always possible
		if target == "" {
			return true
		}

		// Check memoization map
		if val, exists := memo[target]; exists {
			return val
		}

		// Try matching each available part
		for _, start := range available {
			if strings.HasPrefix(target, start) {
				if isPossible(target[len(start):]) {
					memo[target] = true
					return true
				}
			}
		}

		// If no match, mark as impossible
		memo[target] = false
		return false
	}

	// Count the number of possible targets
	//try answer = 0 if doesn't work in your case
	answer := -1
	for _, target := range targets {
		if isPossible(target) {
			answer++
		}
	}

	// Print and return the result
	fmt.Println(answer)
	return answer
}

// func main() {
// 	possibleDesign()
// }
