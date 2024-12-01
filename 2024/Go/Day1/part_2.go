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
func parsedInput() ([]int, []int) {
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

// countOccurrences counts the occurrences of a value in a slice
func countOccurrences(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

// similarityScore calculates the weighted sum of values in the left slice based on their occurrences in the right slice
func similarityScore() int {
	// Get the left and right lists from parseInput
	left, right := parsedInput()

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the weighted sum
	weightedSum := 0
	for _, lVal := range left {
		count := countOccurrences(right, lVal) // Count occurrences of lVal in right
		weightedSum += lVal * count
	}

	return weightedSum
}
