package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func count(cfg string, nums []int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if strings.Contains(cfg, "#") {
			return 0
		}
		return 1
	}

	result := 0

	if cfg[0] == '.' || cfg[0] == '?' {
		result += count(cfg[1:], nums)
	}

	if cfg[0] == '#' || cfg[0] == '?' {
		if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") && (nums[0] == len(cfg) || cfg[nums[0]] != '#') {
			if nums[0] == len(cfg) {
				result += count("", nums[1:])
			} else {
				result += count(cfg[nums[0]+1:], nums[1:])
			}
		}
	}

	return result
}

func SumOfBrokenSprings() {
	file, _ := os.Open("input.txt")
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		cfg := parts[0]
		numsStr := strings.Split(parts[1], ",")
		var nums []int
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		total += count(cfg, nums)
	}

	fmt.Println(total)
}
