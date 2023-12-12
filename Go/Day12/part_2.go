package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func SumOfPossibleArrangementCount() int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	records, groups := parse(string(input))
	res := 0
	for i := range records {
		res += solve(unfoldRecord(records[i]), unfoldGroup(groups[i]))
	}
	println(res)
	return res
}

func unfoldRecord(record string) string {
	var res strings.Builder
	for i := 0; i < len(record)*5; i++ {
		if i != 0 && i%len(record) == 0 {
			res.WriteByte('?')
		}
		res.WriteByte(record[i%len(record)])
	}

	return res.String()
}

func unfoldGroup(group []int) []int {
	var res []int
	for i := 0; i < len(group)*5; i++ {
		res = append(res, group[i%len(group)])
	}

	return res
}

func solve(record string, group []int) int {
	var cache [][]int
	for i := 0; i < len(record); i++ {
		cache = append(cache, make([]int, len(group)+1))
		for j := 0; j < len(group)+1; j++ {
			cache[i][j] = -1
		}
	}

	return dp(0, 0, record, group, cache)
}

func dp(i, j int, record string, group []int, cache [][]int) int {
	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	res := 0
	if record[i] == '.' {
		res = dp(i+1, j, record, group, cache)
	} else {
		if record[i] == '?' {
			res += dp(i+1, j, record, group, cache)
		}
		if j < len(group) {
			count := 0
			for k := i; k < len(record); k++ {
				if count > group[j] || record[k] == '.' || count == group[j] && record[k] == '?' {
					break
				}
				count += 1
			}

			if count == group[j] {
				if i+count < len(record) && record[i+count] != '#' {
					res += dp(i+count+1, j+1, record, group, cache)
				} else {
					res += dp(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = res
	return res
}

func parse(input string) ([]string, [][]int) {
	var records []string
	var groups [][]int

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		parts := strings.Split(line, " ")
		records = append(records, parts[0])
		var group []int
		for _, num := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(num)
			group = append(group, num)
		}
		groups = append(groups, group)
	}

	return records, groups
}
