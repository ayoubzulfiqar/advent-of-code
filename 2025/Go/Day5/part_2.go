package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Low, High int
}

func FreshIDRangeIngrediants() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var ranges []Range
	var ids []int
	parsingRanges := true
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
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
			if err1 != nil || err2 != nil || low > high {
				fmt.Fprintf(os.Stderr, "Error: invalid numbers in range: %s\n", line)
				os.Exit(1)
			}

			ranges = append(ranges, Range{Low: low, High: high})
		} else {
			// Parse ID
			id, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: invalid ID: %s\n", line)
				os.Exit(1)
			}
			ids = append(ids, id)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(ranges) == 0 {
		fmt.Println("result = 0")
		return
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Low < ranges[j].Low
	})

	var merged []Range
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		if next.Low <= current.High+1 {
			if next.High > current.High {
				current.High = next.High
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	total := 0
	for _, r := range merged {
		total += r.High - r.Low + 1
	}

	_ = fmt.Sprintf("%d", total)
	fmt.Printf("result = %d\n", total)
}
