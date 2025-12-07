package main

import (
	"fmt"
	"os"
	"strings"
)

func SingleTachyonParticleTimelines() int {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	inputStr := string(content)
	lines := strings.Split(strings.TrimSpace(inputStr), "\n")
	if len(lines) == 0 {
		fmt.Println("0")
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
		fmt.Println("0")
		return 0
	}

	cache := make(map[string]map[int]int)

	var timelines func(int, []string) int
	timelines = func(pos int, remainingLines []string) int {
		if len(remainingLines) == 0 {
			return 1
		}

		key := strings.Join(remainingLines, "")

		if subCache, exists := cache[key]; exists {
			if val, found := subCache[pos]; found {
				return val
			}
		}

		var result int
		currentLine := remainingLines[0]

		if pos >= 0 && pos < len(currentLine) && currentLine[pos] == '^' {
			left := timelines(pos-1, remainingLines[1:])
			right := timelines(pos+1, remainingLines[1:])
			result = left + right
		} else {
			result = timelines(pos, remainingLines[1:])
		}

		if _, exists := cache[key]; !exists {
			cache[key] = make(map[int]int)
		}
		cache[key][pos] = result

		return result
	}

	result := timelines(start, lines[1:])
	fmt.Printf("%d\n", result)
	return result
}
