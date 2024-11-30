package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EngineSchematicSum() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	cs := make(map[string]bool)

	for r, row := range grid {
		for c, ch := range row {
			if ch >= '0' && ch <= '9' || ch == '.' {
				continue
			}
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					nr, nc := r+dr, c+dc
					if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[nr]) || grid[nr][nc] < '0' || grid[nr][nc] > '9' {
						continue
					}
					for nc > 0 && grid[nr][nc-1] >= '0' && grid[nr][nc-1] <= '9' {
						nc--
					}
					cs[fmt.Sprintf("(%d,%d)", nr, nc)] = true
				}
			}
		}
	}

	var ns []int

	for key := range cs {
		var r, c int
		fmt.Sscanf(key, "(%d,%d)", &r, &c)
		s := ""
		for c < len(grid[r]) && grid[r][c] >= '0' && grid[r][c] <= '9' {
			s += string(grid[r][c])
			c++
		}
		n, err := strconv.Atoi(s)
		if err == nil {
			ns = append(ns, n)
		}
	}

	sum := 0
	for _, num := range ns {
		sum += num
	}

	fmt.Println(sum)
}
