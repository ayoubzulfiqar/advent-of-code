package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type State map[string]int

// ParseInput reads the input file and splits it into two parts.
func parsingInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [][]string
	current := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, strings.Split(strings.TrimSpace(current), "\n"))
			current = ""
		} else {
			current += line + "\n"
		}
	}
	if current != "" {
		result = append(result, strings.Split(strings.TrimSpace(current), "\n"))
	}

	return result
}

// simulateGates processes the gates and computes the final state.
func simulateGates(data [][]string) State {
	state := make(State)
	for _, wire := range data[0] {
		parts := strings.Split(wire, ": ")
		value, _ := strconv.Atoi(parts[1])
		state[parts[0]] = value
	}

	re := regexp.MustCompile(`^(.*) (AND|OR|XOR) (.*) -> (.*)$`)
	loop := true

	for loop {
		shouldLoopAgain := false
		for _, gate := range data[1] {
			match := re.FindStringSubmatch(gate)
			if match != nil {
				left, operator, right, output := match[1], match[2], match[3], match[4]
				leftVal, leftExists := state[left]
				rightVal, rightExists := state[right]

				if !leftExists || !rightExists {
					shouldLoopAgain = true
					continue
				}

				switch operator {
				case "AND":
					state[output] = leftVal & rightVal
				case "OR":
					state[output] = leftVal | rightVal
				case "XOR":
					state[output] = leftVal ^ rightVal
				}
			}
		}
		loop = shouldLoopAgain
	}

	return state
}

// decimalWireZ computes the decimal value for wires starting with "z".
func decimalWireZ() int {
	data := parsingInput()
	state := simulateGates(data)

	// Collect and sort keys starting with "z"
	var zKeys []string
	for name := range state {
		if strings.HasPrefix(name, "z") {
			zKeys = append(zKeys, name)
		}
	}

	// Sort keys in descending order
	sort.Slice(zKeys, func(i, j int) bool {
		return zKeys[i] > zKeys[j]
	})

	// Concatenate values in the sorted order
	var bits strings.Builder
	for _, key := range zKeys {
		bits.WriteString(strconv.Itoa(state[key]))
	}

	// Convert binary string to integer
	result, _ := strconv.ParseInt(bits.String(), 2, 64)
	return int(result)
}

// func main() {
// 	fmt.Println("Part 1", DecimalWireZ())
// }
