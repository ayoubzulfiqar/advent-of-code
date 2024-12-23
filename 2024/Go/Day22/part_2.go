package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func randomNumber(seed int) int {
	seed = ((seed << 6) ^ seed) % 16777216
	seed = ((seed >> 5) ^ seed) % 16777216
	seed = ((seed << 11) ^ seed) % 16777216
	return seed
}

func getMostBananas() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := make(map[string][]int)

	for scanner.Scan() {
		numStr := strings.TrimSpace(scanner.Text())
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		seed := num
		visited := make(map[string]struct{})
		changes := []int{}

		for i := 0; i < 2000; i++ {
			nextSeed := randomNumber(seed)
			changes = append(changes, (nextSeed%10)-(seed%10))
			seed = nextSeed

			if len(changes) == 4 {
				key := strings.Join(intSliceToStringSlice(changes), ",")
				if _, found := visited[key]; !found {
					if _, exists := ranges[key]; !exists {
						ranges[key] = []int{}
					}
					ranges[key] = append(ranges[key], nextSeed%10)
					visited[key] = struct{}{}
				}
				changes = changes[1:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	maxSum := 0
	for _, rangeValues := range ranges {
		sum := 0
		for _, val := range rangeValues {
			sum += val
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func intSliceToStringSlice(slice []int) []string {
	stringSlice := make([]string, len(slice))
	for i, val := range slice {
		stringSlice[i] = strconv.Itoa(val)
	}
	return stringSlice
}

// func main() {
// 	fmt.Println(getMostBananas())
// }
