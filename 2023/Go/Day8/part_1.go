package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func process(contents string) uint32 {
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
	currentNode := "AAA"
	var steps uint32
	turnIdx := 0
	for currentNode != "ZZZ" {
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
	return steps
}

func StepsToReachZ() {
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

	result := process(contents)
	fmt.Printf("result = %d\n", result)
}
