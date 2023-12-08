package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EngineGearRationSum() {

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

	total := 0

	for r, row := range grid {
		for c, ch := range row {
			if ch != '*' {
				continue
			}

			cs := make(map[string]bool)

			for _, cr := range []int{r - 1, r, r + 1} {
				for _, cc := range []int{c - 1, c, c + 1} {
					if cr < 0 || cr >= len(grid) || cc < 0 || cc >= len(grid[cr]) || grid[cr][cc] < '0' || grid[cr][cc] > '9' {
						continue
					}
					for cc > 0 && grid[cr][cc-1] >= '0' && grid[cr][cc-1] <= '9' {
						cc--
					}
					cs[fmt.Sprintf("(%d,%d)", cr, cc)] = true
				}
			}

			if len(cs) != 2 {
				continue
			}

			var ns []int

			for key := range cs {
				var cr, cc int
				fmt.Sscanf(key, "(%d,%d)", &cr, &cc)
				s := ""
				for cc < len(grid[cr]) && grid[cr][cc] >= '0' && grid[cr][cc] <= '9' {
					s += string(grid[cr][cc])
					cc++
				}
				n, err := strconv.Atoi(s)
				if err == nil {
					ns = append(ns, n)
				}
			}

			total += ns[0] * ns[1]
		}
	}

	fmt.Println(total)
}
