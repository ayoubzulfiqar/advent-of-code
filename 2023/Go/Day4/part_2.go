package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TotalScratchcards() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	var s []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	total := 0
	multiString := map[int]int{}

	for i := range s {
		multiString[i] = 1
	}

	for i, line := range s {
		game := strings.Split(line, ": ")

		games := strings.Split(game[1], "|")

		left := strings.Split(games[0], " ")
		right := strings.Split(games[1], " ")

		leftNumbers := parseNumbers(left)
		rightNumbers := parseNumbers(right)

		matches := intersection(leftNumbers, rightNumbers)

		multiplier := multiString[i]

		total += multiplier

		for j := range matches {
			multiString[i+j+1] += multiplier
		}
	}
	fmt.Println(total)

	return total
}

func parseNumbers(nums []string) []int {
	var result []int
	for _, n := range nums {
		if n == "" {
			continue
		}
		n = strings.TrimSpace(n)
		val, err := strconv.Atoi(n)
		if err == nil {
			result = append(result, val)
		}
	}
	return result
}

func intersection(a, b []int) []int {
	var result []int
	seen := make(map[int]bool)
	for _, v := range a {
		seen[v] = true
	}
	for _, v := range b {
		if seen[v] {
			result = append(result, v)
		}
	}
	return result
}
