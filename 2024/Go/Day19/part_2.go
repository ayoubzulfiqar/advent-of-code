package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func differentWaysToMakeDesign() int {
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
	available := strings.Split(lines[0], ", ")
	targets := lines[1:]

	// Trim spaces around available pieces
	for i := range available {
		available[i] = strings.TrimSpace(available[i])
	}

	// Memoization map to store results of sub-problems
	memo := make(map[string]int)

	// Recursive function to count the number of ways to construct the target
	var numWays func(string) int
	numWays = func(target string) int {
		// Base case: empty target has exactly one way to be constructed (by doing nothing)
		if target == "" {
			return 1
		}

		// Check memoization map
		if val, exists := memo[target]; exists {
			return val
		}

		// Count all possible ways
		totalWays := 0
		for _, start := range available {
			if strings.HasPrefix(target, start) {
				totalWays += numWays(target[len(start):])
			}
		}

		// Store the result in the memo map
		memo[target] = totalWays
		return totalWays
	}

	// Compute the total number of ways for all targets
	// try answer = 0 in your case if it does not work
	answer := -1
	for _, target := range targets {
		answer += numWays(target)
	}

	// Print and return the result
	fmt.Println(answer)
	return answer
}

// func main() {
// 	differentWaysToMakeDesign()
// }
