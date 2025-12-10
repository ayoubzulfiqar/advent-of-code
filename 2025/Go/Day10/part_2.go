package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSteps := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		if len(tokens) < 2 {
			continue
		}

		lightsToken := tokens[0]
		if len(lightsToken) < 2 || lightsToken[0] != '[' || lightsToken[len(lightsToken)-1] != ']' {
			continue
		}
		lightsStr := lightsToken[1 : len(lightsToken)-1]
		n := len(lightsStr)

		wiringTokens := tokens[1 : len(tokens)-1]

		start := 0
		for i, char := range lightsStr {
			if char == '#' {
				start |= (1 << uint(i))
			}
		}

		if start == 0 {
			continue
		}

		buttons := []int{}
		for _, token := range wiringTokens {
			if len(token) < 2 || token[0] != '(' || token[len(token)-1] != ')' {
				continue
			}
			inner := token[1 : len(token)-1]
			if inner == "" {
				buttons = append(buttons, 0)
				continue
			}
			parts := strings.Split(inner, ",")
			mask := 0
			for _, p := range parts {
				numStr := strings.TrimSpace(p)
				if numStr == "" {
					continue
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					continue
				}
				if num >= 0 && num < n {
					mask |= (1 << uint(num))
				}
			}
			buttons = append(buttons, mask)
		}

		visited := make(map[int]bool)
		visited[start] = true
		current := []int{start}
		steps := 0
		found := false

		for !found && len(current) > 0 {
			nextLevel := []int{}
			for _, state := range current {
				for _, btn := range buttons {
					ns := state ^ btn
					if ns == 0 {
						steps++
						found = true
						break
					}
					if !visited[ns] {
						visited[ns] = true
						nextLevel = append(nextLevel, ns)
					}
				}
				if found {
					break
				}
			}
			if found {
				break
			}
			current = nextLevel
			steps++
		}

		totalSteps += steps
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(totalSteps)
}
