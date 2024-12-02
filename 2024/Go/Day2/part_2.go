// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // readInput reads the input file and converts it into a 2D slice of integers
// func readInput(filename string) [][]int {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	var data [][]int
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.TrimSpace(line) == "" {
// 			continue
// 		}

// 		// Split line into fields and convert to integers
// 		strNumbers := strings.Fields(line)
// 		var intNumbers []int
// 		for _, strNum := range strNumbers {
// 			num, err := strconv.Atoi(strNum)
// 			if err != nil {
// 				log.Fatalf("Failed to convert string to integer: %v", err)
// 			}
// 			intNumbers = append(intNumbers, num)
// 		}
// 		data = append(data, intNumbers)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Error reading file: %v", err)
// 	}

// 	return data
// }

// // test checks if removing one element makes a sequence safe
// func test(line []int, index int) bool {
// 	for i := range line {
// 		// Create a copy of the line without the ith element
// 		tempLine := append([]int{}, line[:i]...)
// 		tempLine = append(tempLine, line[i+1:]...)

// 		// Check if the modified line is safe
// 		if isSafe(tempLine) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // isSafe determines if a line is "safe" based on its pattern
// func isSafe(line []int) bool {
// 	if len(line) < 2 {
// 		return false // A single number cannot form a pattern
// 	}

// 	var allowedDiffs map[int]bool
// 	if line[1] < line[0] { // Decreasing
// 		allowedDiffs = map[int]bool{-1: true, -2: true, -3: true}
// 	} else if line[1] > line[0] { // Increasing
// 		allowedDiffs = map[int]bool{1: true, 2: true, 3: true}
// 	} else {
// 		return false // First two numbers are the same, not safe
// 	}

// 	// Check if all consecutive differences are allowed
// 	for i := 1; i < len(line); i++ {
// 		diff := line[i] - line[i-1]
// 		if !allowedDiffs[diff] {
// 			return false
// 		}
// 	}

// 	return true
// }

// // singleLevelSafeReports calculates the total number of safe reports
// func singleLevelSafeReports(filename string) int {
// 	data := readInput(filename)
// 	totalSafe := 0
// 	var problems []int
// 	var problemIndexes [][]int

// 	for _, line := range data {
// 		problemCount := 0
// 		var currentIndexes []int

// 		if len(line) > 1 {
// 			if line[1] < line[0] { // Decreasing
// 				for j := 1; j < len(line); j++ {
// 					if (line[j] - line[j-1]) < -3 || (line[j] - line[j-1]) > -1 {
// 						problemCount++
// 						currentIndexes = append(currentIndexes, j)
// 					}
// 				}
// 			} else if line[1] > line[0] { // Increasing
// 				for j := 1; j < len(line); j++ {
// 					if (line[j] - line[j-1]) > 3 || (line[j] - line[j-1]) < 1 {
// 						problemCount++
// 						currentIndexes = append(currentIndexes, j)
// 					}
// 				}
// 			} else { // Equal
// 				problemCount++
// 				currentIndexes = append(currentIndexes, 0)
// 			}
// 		}

// 		problems = append(problems, problemCount)
// 		problemIndexes = append(problemIndexes, currentIndexes)
// 	}

// 	// Check problem indexes and test lines
// 	for i, line := range data {
// 		if problems[i] == 0 {
// 			totalSafe++
// 		} else {
// 			if test(line, problemIndexes[i][0]) {
// 				totalSafe++
// 			}
// 		}
// 	}

// 	return totalSafe
// }

// // main function to execute the program
//
//	func main() {
//		filename := "D:/Projects/advent-of-code/2024/Python/Day2/input.txt"
//		result := singleLevelSafeReports(filename)
//		fmt.Println("Total Safe Reports:", result)
//	}
package main
