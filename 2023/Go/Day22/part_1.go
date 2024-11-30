package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func BrickSafelyChosen() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var bricks [][]int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(strings.ReplaceAll(line, "~", ","), ",")
		var brick []int
		for _, v := range values {
			num := 0
			fmt.Sscanf(v, "%d", &num)
			brick = append(brick, num)
		}
		bricks = append(bricks, brick)
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i][2] < bricks[j][2]
	})

	overlaps := func(a, b []int) bool {
		return max(a[0], b[0]) <= min(a[3], b[3]) && max(a[1], b[1]) <= min(a[4], b[4])
	}

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

	sort.Slice(bricks, func(i, j int) bool {
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
		satisfies := true
		for j := range kSupportsV[i] {
			if len(vSupportsK[j]) < 2 {
				satisfies = false
				break
			}
		}
		if satisfies {
			total++
		}
	}

	fmt.Println(total)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
