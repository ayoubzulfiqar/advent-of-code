package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func resultingFileChecksum() int {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the input
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	con := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create initial file system
	var fileSystem []int
	var fileValue int = 0

	for i, char := range con {
		count, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		for j := 0; j < count; j++ {
			if i%2 == 0 {
				fileSystem = append(fileSystem, fileValue)
			} else {
				fileSystem = append(fileSystem, -1)
			}
		}
		if i%2 == 0 {
			fileValue++
		}
	}

	// Keep swapping empty spaces from bottom to top
	bottom, top := 0, len(fileSystem)-1
	for bottom < top {
		if fileSystem[bottom] == -1 {
			for fileSystem[top] == -1 {
				top--
			}
			if top < bottom {
				break
			}
			fileSystem[bottom], fileSystem[top] = fileSystem[top], -1
		}
		bottom++
	}

	// Find checksum of the file system
	totalSum := 0
	for i, value := range fileSystem {
		if value == -1 {
			break
		}
		totalSum += value * i
	}

	fmt.Println(totalSum)
	return totalSum
}

