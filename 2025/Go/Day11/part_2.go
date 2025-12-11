package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dfsBitMask(adj map[string][]string, start string) int64 {
	type State struct {
		node string
		mask uint8 // bit 0: fft, bit 1: dac
	}

	memo := make(map[State]int64)

	var dfs func(curr string, mask uint8) int64
	dfs = func(curr string, mask uint8) int64 {
		if curr == "out" {
			if mask == 0x03 { // Both bits set (fft AND dac)
				return 1
			}
			return 0
		}

		state := State{curr, mask}
		if val, exists := memo[state]; exists {
			return val
		}

		var ans int64
		for _, child := range adj[curr] {
			newMask := mask
			switch child {
			case "fft":
				newMask |= 0x01 // Set fft bit
			case "dac":
				newMask |= 0x02 // Set dac bit
			}
			ans += dfs(child, newMask)
		}

		memo[state] = ans
		return ans
	}

	return dfs(start, 0)
}

func DacAndFftPath() {

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

	result := dfsBitMask(adj, "svr")
	fmt.Println(result)

}
