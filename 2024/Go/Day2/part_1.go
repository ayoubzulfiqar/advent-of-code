package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// readInput reads the input file and converts it into a 2D slice of integers
func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Split line into fields and convert to integers
		strNumbers := strings.Fields(line)
		var intNumbers []int
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatalf("Failed to convert string to integer: %v", err)
			}
			intNumbers = append(intNumbers, num)
		}
		data = append(data, intNumbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return data
}

// isSafe checks if a line is "safe" based on its pattern
func isSafe(line []int) bool {
	if len(line) < 2 {
		return false // A single number cannot form a pattern
	}

	var allowedDiffs map[int]bool
	if line[1] < line[0] { // Decreasing
		allowedDiffs = map[int]bool{-1: true, -2: true, -3: true}
	} else if line[1] > line[0] { // Increasing
		allowedDiffs = map[int]bool{1: true, 2: true, 3: true}
	} else {
		return false // First two numbers are the same, not safe
	}

	// Check if all consecutive differences are allowed
	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]
		if !allowedDiffs[diff] {
			return false
		}
	}

	return true
}

// safeReports calculates and prints the total number of safe reports
func safeReports() int {
	filename := "D:/Projects/advent-of-code/2024/Go/Day2/input.txt"
	data := readInput(filename)

	totalSafe := 0
	for _, line := range data {
		if isSafe(line) {
			totalSafe++
		}
	}

	fmt.Println("Part 1:", totalSafe)
	return totalSafe
}
