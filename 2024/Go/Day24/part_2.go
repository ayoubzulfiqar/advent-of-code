package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// parseInput reads the input file and splits it into parts.
func parseInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var parts [][]string
	scanner := bufio.NewScanner(file)
	var currentPart []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentPart) > 0 {
				parts = append(parts, currentPart)
				currentPart = []string{}
			}
		} else {
			currentPart = append(currentPart, line)
		}
	}
	if len(currentPart) > 0 {
		parts = append(parts, currentPart)
	}

	return parts
}

// find searches for a gate matching the given inputs and operator.
func find(a, b, operator string, gates []string) string {
	for _, gate := range gates {
		if strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", a, operator, b)) ||
			strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", b, operator, a)) {
			parts := strings.Split(gate, " -> ")
			return parts[len(parts)-1]
		}
	}
	return ""
}

// swapAndJoinWires performs the logic to swap and join wires.
func swapAndJoinWires() string {
	data := parseInput()
	gates := data[1]
	var swapped []string
	var c0 string

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		// Half adder logic
		m1 = find("x"+n, "y"+n, "XOR", gates)
		n1 = find("x"+n, "y"+n, "AND", gates)

		if c0 != "" {
			r1 = find(c0, m1, "AND", gates)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(c0, m1, "AND", gates)
			}

			z1 = find(c0, m1, "XOR", gates)

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}

			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}

			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(r1, n1, "OR", gates)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if c0 == "" {
			c0 = n1
		} else {
			c0 = c1
		}
	}

	// Sort and join swapped wires
	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}

// func main() {
// 	fmt.Println("Part 2:", swapAndJoinWires())
// }
