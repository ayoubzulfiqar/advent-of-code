package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parseInput reads the input file and splits the integers into two lists
func parseInput() ([]int, []int) {
	filePath := "D:/Projects/advent-of-code/2024/Go/Day1/input.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Initialize slices for data, left, and right
	data := []int{}
	left := []int{}
	right := []int{}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line) // Split line into fields
		for _, n := range numbers {
			value, err := strconv.Atoi(n) // Convert string to int
			if err == nil {
				data = append(data, value)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Distribute elements into left and right slices
	for index, value := range data {
		if index%2 == 0 {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	return left, right
}

// totalDistance calculates the sum of absolute differences between two sorted lists
func totalDistance() int {
	// Get the left and right lists from parseInput
	left, right := parseInput()

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the total distance
	totalDifference := 0
	for i := 0; i < len(left); i++ {
		totalDifference += abs(left[i] - right[i])
	}

	return totalDifference
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
