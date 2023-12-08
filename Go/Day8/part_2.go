package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func processing(contents string) uint64 {
	scanner := bufio.NewScanner(strings.NewReader(contents))
	nodesMap := make(map[string]struct{ Left, Right string })
	scanner.Scan() // Read turns
	turns := scanner.Text()
	scanner.Scan() // Skip a line
	nodeRe := regexp.MustCompile(`([A-Z][A-Z][A-Z]) = \(([A-Z][A-Z][A-Z]), ([A-Z][A-Z][A-Z])\)`)
	for scanner.Scan() {
		line := scanner.Text()
		nodeCaps := nodeRe.FindStringSubmatch(line)
		source := nodeCaps[1]
		lDest := nodeCaps[2]
		rDest := nodeCaps[3]
		nodesMap[source] = struct{ Left, Right string }{Left: lDest, Right: rDest}
	}
	startNodes := []string{}
	for k := range nodesMap {
		if strings.HasSuffix(k, "A") {
			startNodes = append(startNodes, k)
		}
	}
	allSteps := []uint64{}
	for _, startNode := range startNodes {
		currentNode := startNode
		var steps uint64
		turnIdx := 0
		for !strings.HasSuffix(currentNode, "Z") {
			turn := turns[turnIdx]
			switch turn {
			case 'L':
				currentNode = nodesMap[currentNode].Left
			case 'R':
				currentNode = nodesMap[currentNode].Right
			default:
				panic("unexpected char in turns")
			}
			steps++
			turnIdx = (turnIdx + 1) % len(turns)
		}
		allSteps = append(allSteps, steps)
	}
	result := uint64(1)
	for _, x := range allSteps {
		result = lcm(result, x)
	}
	return result
}

// gcd calculates the greatest common divisor of two numbers using Euclid's algorithm.
func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the least common multiple of two numbers using the GCD.
func lcm(a, b uint64) uint64 {
	return a / gcd(a, b) * b
}

func OnlyNodesEndWithZ() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var contents string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents += scanner.Text() + "\n"
	}

	result := processing(contents)
	fmt.Printf("result = %d\n", result)
}
