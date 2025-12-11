package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfs(adj map[string][]string, curr string, memo map[string]int64, visited map[string]bool) int64 {
	if curr == "out" {
		return 1
	}

	if val, exists := memo[curr]; exists {
		return val
	}

	if visited[curr] {
		return 0
	}

	visited[curr] = true
	var ans int64
	for _, child := range adj[curr] {
		ans += dfs(adj, child, memo, visited)
	}
	visited[curr] = false

	memo[curr] = ans
	return ans
}

func YouPathOut() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	adj := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSuffix(parts[0], ":")

		values := parts[1:]
		adj[key] = values
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	memo := make(map[string]int64, len(adj))
	visited := make(map[string]bool, len(adj))

	result := dfs(adj, "you", memo, visited)

	fmt.Println(result)

}

/*

// Iterative Approach





func countPathsIterative(adj map[string][]string, start string) int64 {
	if start == "out" {
		return 1
	}

	memo := make(map[string]int64, len(adj))
	memo["out"] = 1

	stack := []string{start}
	visited := make(map[string]bool, len(adj))
	processing := make(map[string]bool, len(adj))

	for len(stack) > 0 {
		curr := stack[len(stack)-1]

		if processing[curr] {
			var total int64
			for _, child := range adj[curr] {
				total += memo[child]
			}
			memo[curr] = total
			processing[curr] = false
			stack = stack[:len(stack)-1]
			continue
		}

		if visited[curr] {
			stack = stack[:len(stack)-1]
			continue
		}

		allChildrenComputed := true
		for _, child := range adj[curr] {
			if _, ok := memo[child]; !ok && child != "out" {
				allChildrenComputed = false
				if !visited[child] {
					stack = append(stack, child)
				}
			}
		}

		if allChildrenComputed {
			var total int64
			for _, child := range adj[curr] {
				total += memo[child]
			}
			memo[curr] = total
			visited[curr] = true
			stack = stack[:len(stack)-1]
		} else {
			processing[curr] = true
		}
	}

	return memo[start]
}

func YouPathOut() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	adj := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSuffix(parts[0], ":")
		adj[key] = parts[1:]
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	result := countPathsIterative(adj, "you")
	fmt.Println(result)
}





*/
