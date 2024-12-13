package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func formulas() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)

	}
	defer file.Close()

	// Read the entire file content into a string
	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)

	}
	// Split input by double newlines
	sections := strings.Split(input, "\n\n")
	var systems [][]int
	re := regexp.MustCompile(`\d+`)

	// Extract integers from each section
	for _, section := range sections {
		matches := re.FindAllString(section, -1)
		system := make([]int, len(matches))
		for i, match := range matches {
			system[i], _ = strconv.Atoi(match)
		}
		systems = append(systems, system)
	}
	return systems
}

// Function to solve the system of equations and compute the result
func largerRun(machines [][]int, shift int) int {
	result := 0

	for _, machine := range machines {
		ax, ay, bx, by, px, py := float64(machine[0]), float64(machine[1]), float64(machine[2]), float64(machine[3]), float64(machine[4]), float64(machine[5])
		px += float64(shift)
		py += float64(shift)

		// Avoid division by zero or invalid calculations
		if by == 0 || ax == ay*bx/by {
			continue
		}

		// Solve the system of equations
		a := (px - py*bx/by) / (ax - ay*bx/by)
		b := (py - a*ay) / by

		// Round the results using math.Round for better precision
		ra := math.Round(a)
		rb := math.Round(b)

		// Check if the solution is valid
		if int(ra)*int(ax)+int(rb)*int(bx) == int(px) && int(ra)*int(ay)+int(rb)*int(by) == int(py) && ra >= 0 && rb >= 0 {
			result += int(ra)*3 + int(rb)
		}
	}
	return result
}

func fewestTokenToSpend() {
	runMachine := formulas()
	fmt.Println("p2", run(runMachine, 10000000000000))
}
