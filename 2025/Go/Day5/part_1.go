package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IDFreshIngrediants() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimRight(content.String(), "\n"), "\n")

	var ranges []struct{ Low, High int }
	var ids []int
	var i int
	parsingRanges := true

	for i = 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "Error: invalid range format: %s\n", line)
				os.Exit(1)
			}

			low, err1 := strconv.Atoi(parts[0])
			high, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Fprintf(os.Stderr, "Error: invalid numbers in range: %s\n", line)
				os.Exit(1)
			}

			ranges = append(ranges, struct{ Low, High int }{Low: low, High: high})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: invalid ID: %s\n", line)
				os.Exit(1)
			}
			ids = append(ids, id)
		}
	}

	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.Low && id <= r.High {
				count++
				break
			}
		}
	}

	fmt.Printf("result = %d\n", count)
}
