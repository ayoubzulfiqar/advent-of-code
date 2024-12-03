package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func resultOfEnabledMultiply() uint32 {
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
	// Compile the regex pattern
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|don't\(\)|do\(\)`)

	// Set the enabled flag to true
	enabled := true
	total := uint32(0)

	// Iterate over the regex matches
	matches := re.FindAllStringSubmatch(corruptedMemory, -1)
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				// Parse the numbers and multiply them
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				total += uint32(x * y)
			}
		}
	}

	return total
}
