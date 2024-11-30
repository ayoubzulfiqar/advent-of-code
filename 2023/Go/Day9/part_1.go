package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NextSumOfExtraPolatedValues() {
	histories := make([][]int, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error reading input.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := parseInts(line)
		histories = append(histories, values)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading input.txt: %v\n", err)
		os.Exit(1)
	}

	extrapolatedValues := make([]int, 0)

	for _, history := range histories {
		subLists := [][]int{history}
		allZeroes := false
		for !allZeroes {
			sublist := make([]int, 0)
			for i := 0; i < len(subLists[len(subLists)-1])-1; i++ {
				difference := subLists[len(subLists)-1][i+1] - subLists[len(subLists)-1][i]
				sublist = append(sublist, difference)
			}
			allZeroes = allZeroesInSlice(sublist)
			subLists = append(subLists, sublist)
		}

		subLists[len(subLists)-1] = append(subLists[len(subLists)-1], 0)
		for i := len(subLists) - 2; i >= 0; i-- {
			extrapolatedValue := subLists[i][len(subLists[i])-1] + subLists[i+1][len(subLists[i+1])-1]
			subLists[i] = append(subLists[i], extrapolatedValue)
		}

		extrapolatedValues = append(extrapolatedValues, subLists[0][len(subLists[0])-1])
	}

	ans := sum(extrapolatedValues)
	fmt.Println(ans)
}

func parseInts(input string) []int {
	var nums []int
	fields := strings.Fields(input)

	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		nums = append(nums, num)
	}

	return nums
}

func allZeroesInSlice(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
