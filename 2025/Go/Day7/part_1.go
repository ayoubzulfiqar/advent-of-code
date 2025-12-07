package main

import (
	"fmt"
	"os"
	"strings"
)

func BeamSplitTime() int {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	inputStr := string(content)
	lines := strings.Split(strings.TrimSpace(inputStr), "\n")
	if len(lines) == 0 {
		fmt.Println("Part 1: 0")
		return 0
	}

	firstLine := lines[0]
	start := -1
	for i, ch := range firstLine {
		if ch == 'S' {
			start = i
			break
		}
	}
	if start == -1 {
		fmt.Println("Part 1: 0")
		return 0
	}

	beams := make(map[int]bool)
	beams[start] = true
	count := 0

	for _, line := range lines[1:] {
		splitters := make(map[int]bool)
		for i, ch := range line {
			if ch == '^' {
				splitters[i] = true
			}
		}

		splits := make(map[int]bool)
		splitCount := 0
		for pos := range beams {
			if splitters[pos] {
				splits[pos] = true
				splitCount++
			}
		}

		newBeams := make(map[int]bool)
		for pos := range splits {
			if pos+1 < len(line) {
				newBeams[pos+1] = true
			}
			if pos-1 >= 0 {
				newBeams[pos-1] = true
			}
		}

		remainingBeams := make(map[int]bool)
		for pos := range beams {
			if !splits[pos] {
				remainingBeams[pos] = true
			}
		}

		finalBeams := make(map[int]bool)
		for pos := range remainingBeams {
			finalBeams[pos] = true
		}
		for pos := range newBeams {
			finalBeams[pos] = true
		}

		count += splitCount
		beams = finalBeams
	}
	fmt.Printf("Part 1: %d\n", count)
	return count

}
