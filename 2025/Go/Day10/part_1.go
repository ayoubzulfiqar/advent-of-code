package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toBitmask(arr []int) int {
	mask := 0
	for _, v := range arr {
		mask |= 1 << v
	}
	return mask
}

func ConfigureTheIndicatorLights() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	ret := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		lights := parts[0]
		wiring := parts[1 : len(parts)-1]

		// Parse lights (remove brackets and convert to bool array)
		lights = strings.Trim(lights, "[]")
		start := []int{}
		for i, ch := range lights {
			if ch == '#' {
				start = append(start, i)
			}
		}

		// Parse buttons
		buttons := [][]int{}
		for _, wire := range wiring {
			wire = strings.Trim(wire, "()")
			if wire == "" {
				continue
			}
			nums := strings.Split(wire, ",")
			button := []int{}
			for _, numStr := range nums {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Printf("Error converting number: %v\n", err)
					continue
				}
				button = append(button, num)
			}
			buttons = append(buttons, button)
		}

		// Convert to bitmasks
		startMask := toBitmask(start)
		buttonMasks := []int{}
		for _, button := range buttons {
			buttonMasks = append(buttonMasks, toBitmask(button))
		}
		endMask := 0

		// fmt.Printf("Start: %d, Buttons: %v, End: %d\n", startMask, buttonMasks, endMask)

		// BFS through states
		current := map[int]bool{startMask: true}
		iterations := 0
		found := false

		for !found {
			if _, exists := current[endMask]; exists {
				break
			}

			nextSet := make(map[int]bool)
			for currentMask := range current {
				for _, buttonMask := range buttonMasks {
					nextSet[currentMask^buttonMask] = true
				}
			}
			current = nextSet
			iterations++

			// Safety check to prevent infinite loop
			if iterations > 100000 {
				fmt.Println("Warning: Too many iterations, breaking")
				break
			}
		}

		// fmt.Printf("Iterations: %d\n", iterations)
		ret += iterations
	}

	fmt.Printf("Total: %d\n", ret)
}
