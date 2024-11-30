package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func overlaps(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[3], b[3]) && max(a[1], b[1]) <= min(a[4], b[4])
}

func SumOfBricksWouldFall() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var bricks [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var brick []int
		for _, s := range split(line, "~", ",") {
			val := parseInt(s)
			brick = append(brick, val)
		}
		bricks = append(bricks, brick)
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i][2] < bricks[j][2]
	})

	for index, brick := range bricks {
		maxZ := 1
		for _, check := range bricks[:index] {
			if overlaps(brick, check) {
				maxZ = max(maxZ, check[5]+1)
			}
		}
		brick[5] -= brick[2] - maxZ
		brick[2] = maxZ
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i][2] < bricks[j][2]
	})

	kSupportsV := make(map[int]map[int]struct{})
	vSupportsK := make(map[int]map[int]struct{})

	for i := range bricks {
		kSupportsV[i] = make(map[int]struct{})
		vSupportsK[i] = make(map[int]struct{})
	}

	for j, upper := range bricks {
		for i, lower := range bricks[:j] {
			if overlaps(lower, upper) && upper[2] == lower[5]+1 {
				kSupportsV[i][j] = struct{}{}
				vSupportsK[j][i] = struct{}{}
			}
		}
	}

	total := 0

	for i := range bricks {
		q := make([]int, 0)
		for j := range kSupportsV[i] {
			if len(vSupportsK[j]) == 1 {
				q = append(q, j)
			}
		}

		falling := make(map[int]struct{})
		for _, j := range q {
			falling[j] = struct{}{}
		}
		falling[i] = struct{}{}

		for len(q) > 0 {
			j := q[0]
			q = q[1:]
			for k := range kSupportsV[j] {
				if _, ok := falling[k]; !ok {
					if isSubset(vSupportsK[k], falling) {
						q = append(q, k)
						falling[k] = struct{}{}
					}
				}
			}
		}

		total += len(falling) - 1
	}

	fmt.Println(total)
}

func split(s, sep1, sep2 string) []string {
	s = strings.ReplaceAll(s, sep1, sep2)
	return strings.Split(s, sep2)
}

func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func isSubset(set1, set2 map[int]struct{}) bool {
	for key := range set1 {
		if _, ok := set2[key]; !ok {
			return false
		}
	}
	return true
}
