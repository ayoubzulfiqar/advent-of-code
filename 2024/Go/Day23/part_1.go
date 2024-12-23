package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func computerNameT() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse connections
	links := make(map[string]map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, end := parts[0], parts[1]
			if links[start] == nil {
				links[start] = make(map[string]bool)
			}
			if links[end] == nil {
				links[end] = make(map[string]bool)
			}
			links[start][end] = true
			links[end][start] = true
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Find possibilities
	possibilities := make(map[[3]string]bool)
	for c1 := range links {
		for c2 := range links[c1] {
			if c2 == c1 {
				continue
			}
			for c3 := range links[c2] {
				if c3 == c1 || c3 == c2 {
					continue
				}
				if links[c3][c1] {
					// Create a sorted triplet
					triplet := [3]string{c1, c2, c3}
					if triplet[0] > triplet[1] {
						triplet[0], triplet[1] = triplet[1], triplet[0]
					}
					if triplet[1] > triplet[2] {
						triplet[1], triplet[2] = triplet[2], triplet[1]
					}
					if triplet[0] > triplet[1] {
						triplet[0], triplet[1] = triplet[1], triplet[0]
					}
					possibilities[triplet] = true
				}
			}
		}
	}

	// Count triplets where any item starts with 't'
	answer := 0
	for triplet := range possibilities {
		for _, item := range triplet {
			if strings.HasPrefix(item, "t") {
				answer++
				break
			}
		}
	}

	fmt.Println(answer)
}
