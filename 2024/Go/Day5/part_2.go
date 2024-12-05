package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RuleSet map[[2]int]struct{}

func parseUpdateAndRules() (RuleSet, [][]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(RuleSet)
	var updates [][]int

	// Parse rules
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 1 {
			break
		}
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}
		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		rules[[2]int{a, b}] = struct{}{}
	}

	// Parse updates
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		update := make([]int, len(nums))
		for i, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}
			update[i] = val
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}

func addMiddlePageNumber() (int, error) {
	rules, updates, err := parseUpdateAndRules()
	if err != nil {
		return 0, err
	}

	isNotSorted := func(update []int) bool {
		for i := 0; i < len(update)-1; i++ {
			if _, exists := rules[[2]int{update[i+1], update[i]}]; exists {
				return true
			}
		}
		return false
	}

	customSort := func(update []int) []int {
		sortedUpdate := make([]int, len(update))
		copy(sortedUpdate, update)

		sort.Slice(sortedUpdate, func(i, j int) bool {
			a, b := sortedUpdate[i], sortedUpdate[j]
			if _, exists := rules[[2]int{a, b}]; exists {
				return true
			}
			if _, exists := rules[[2]int{b, a}]; exists {
				return false
			}
			return a < b
		})

		return sortedUpdate
	}

	total := 0
	for _, update := range updates {
		if isNotSorted(update) {
			sortedUpdate := customSort(update)
			total += sortedUpdate[len(sortedUpdate)/2]
		}
	}

	return total, nil
}
