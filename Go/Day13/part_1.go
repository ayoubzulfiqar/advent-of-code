package main

import (
	"os"
	"strings"
)



func ReflectionSummarizingAllNotes() int {
	b, _ := os.ReadFile("input.txt")
	input := string(b)
	grids := parse(input)
	res := 0
	for i := range grids {
		j := verticalReflection(grids[i], false, -1)
		if j != -1 {
			res += j + 1
		} else {
			k := horizontalReflection(grids[i], false, -1)
			if k != -1 {
				res += (k + 1) * 100
			}
		}
	}

	return res
}

func verticalReflection(grid []string, smudge bool, notAllowed int) int {
	n := len(grid[0])
	for j := 0; j < n-1; j++ {
		left := j
		right := j + 1
		mismatch := false
		done := false

		for left >= 0 && right < n {
			leftStr := getColStr(grid, left)
			rightStr := getColStr(grid, right)
			delta := getDelta(leftStr, rightStr)

			if leftStr != rightStr && (!smudge || (delta > 1 || done)) {
				mismatch = true
				break
			}
			if delta == 1 {
				done = true
			}
			left -= 1
			right += 1
		}

		if !mismatch && j != notAllowed {
			return j
		}

	}

	return -1
}

func horizontalReflection(grid []string, smudge bool, notAllowed int) int {
	m := len(grid)
	for i := 0; i < m-1; i++ {
		top := i
		bottom := i + 1
		mismatch := false
		done := false

		for top >= 0 && bottom < m {
			delta := getDelta(grid[top], grid[bottom])
			if grid[top] != grid[bottom] && (!smudge || (delta > 1 || done)) {
				mismatch = true
				break
			}
			if delta == 1 {
				done = true
			}

			top -= 1
			bottom += 1
		}

		if !mismatch && i != notAllowed {
			return i
		}
	}

	return -1
}

func getDelta(s1, s2 string) int {
	delta := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			delta += 1
		}
	}

	return delta
}

func getColStr(grid []string, j int) string {
	var sb strings.Builder
	for i := 0; i < len(grid); i++ {
		sb.WriteByte(grid[i][j])
	}

	return sb.String()
}

func parse(input string) [][]string {
	var grids [][]string
	for _, block := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n") {
		grids = append(grids, strings.Split(block, "\n"))
	}

	return grids
}
