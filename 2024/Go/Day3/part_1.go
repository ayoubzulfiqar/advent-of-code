package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func additionOfMultiplication() int {
	// Open the input file
	file, err := os.Open("./input.txt")
	// fmt.Println(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	// Read the file content
	scanner := bufio.NewScanner(file)
	var corruptedMemory string
	for scanner.Scan() {
		corruptedMemory += scanner.Text()
	}

	// Define the regex pattern
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(corruptedMemory, -1)

	// Calculate the sum of the products
	total := 0
	for _, match := range matches {
		if len(match) == 3 {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total += x * y
		}
	}

	// Print and return the result
	fmt.Println(total)
	return total
}
