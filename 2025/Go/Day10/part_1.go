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

		lights = strings.Trim(lights, "[]")
		start := []int{}
		for i, ch := range lights {
			if ch == '#' {
				start = append(start, i)
			}
		}

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

		startMask := toBitmask(start)
		buttonMasks := []int{}
		for _, button := range buttons {
			buttonMasks = append(buttonMasks, toBitmask(button))
		}
		endMask := 0

		// fmt.Printf("Start: %d, Buttons: %v, End: %d\n", startMask, buttonMasks, endMask)

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

/*

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
		if len(line) == 0 {
			continue
		}

	
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

	
		lightsPart := parts[0]
		if len(lightsPart) < 2 || lightsPart[0] != '[' || lightsPart[len(lightsPart)-1] != ']' {
			continue
		}
		lightsStr := lightsPart[1 : len(lightsPart)-1]
		n := len(lightsStr)

		wiringParts := parts[1 : len(parts)-1]

	
		startState := 0
		for i, char := range lightsStr {
			if char == '#' {
				startState |= (1 << uint(i))
			}
		}

	
		if startState == 0 {
			continue
		}

		
		buttons := []int{}
		for _, wp := range wiringParts {
			if len(wp) < 2 || wp[0] != '(' || wp[len(wp)-1] != ')' {
				continue
			}
			inner := wp[1 : len(wp)-1]
			if inner == "" {
				buttons = append(buttons, 0)
				continue
			}

			indices := strings.Split(inner, ",")
			mask := 0
			for _, idxStr := range indices {
				if idx, err := strconv.Atoi(strings.TrimSpace(idxStr)); err == nil && idx >= 0 && idx < n {
					mask |= (1 << uint(idx))
				}
			}
			buttons = append(buttons, mask)
		}

		
		visited := make(map[int]bool)
		queue := []int{startState}
		visited[startState] = true
		steps := 0
		found := false

		
		for !found && len(queue) > 0 {
			nextQueue := []int{}

			for _, state := range queue {
				
				for _, btn := range buttons {
					newState := state ^ btn

					if newState == 0 {
						steps++
						found = true
						break
					}

					if !visited[newState] {
						visited[newState] = true
						nextQueue = append(nextQueue, newState)
					}
				}
				if found {
					break
				}
			}

			if found {
				break
			}

			queue = nextQueue
			steps++
		}

		totalSteps += steps
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(totalSteps)
}


*/
