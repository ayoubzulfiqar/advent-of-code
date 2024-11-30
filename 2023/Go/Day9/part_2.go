package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PreviousSumOfExtraPolatedValues() {
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
		values := parseIntegers(line)
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
			allZeroes = allZeroesInSlices(sublist)
			subLists = append(subLists, sublist)
		}

		subLists[len(subLists)-1] = append([]int{0}, subLists[len(subLists)-1]...)
		for i := len(subLists) - 2; i >= 0; i-- {
			extrapolatedValue := subLists[i][0] - subLists[i+1][0]
			subLists[i] = append([]int{extrapolatedValue}, subLists[i]...)
		}

		extrapolatedValues = append(extrapolatedValues, subLists[0][0])
	}

	ans := sums(extrapolatedValues)
	fmt.Println(ans)
}

func parseIntegers(input string) []int {
	var nums []int
	fields := strings.Fields(input)

	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		nums = append(nums, num)
	}

	return nums
}

func allZeroesInSlices(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func sums(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
