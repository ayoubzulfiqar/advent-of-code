package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// type RuleSet map[[2]int]struct{}

func parseRulesAndUpdates() (RuleSet, [][]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(RuleSet)
	updates := [][]int{}

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

func middlePageNumber() (int, error) {
	rules, updates, err := parseRulesAndUpdates()
	if err != nil {
		return 0, err
	}

	isSorted := func(update []int) bool {
		for i := 0; i < len(update)-1; i++ {
			if _, exists := rules[[2]int{update[i+1], update[i]}]; exists {
				return false
			}
		}
		return true
	}

	total := 0
	for _, update := range updates {
		if isSorted(update) {
			total += update[len(update)/2]
		}
	}

	return total, nil
}
