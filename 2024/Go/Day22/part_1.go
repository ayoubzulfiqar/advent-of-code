package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func randomNumbers(seed int) int {
	seed = ((seed << 6) ^ seed) % 16777216
	seed = ((seed >> 5) ^ seed) % 16777216
	seed = ((seed << 11) ^ seed) % 16777216
	return seed
}

func sumOfTwoThousandthNumber() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		numStr := strings.TrimSpace(scanner.Text()) // Remove any extra whitespace or carriage returns
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		seed := num
		for i := 0; i < 2000; i++ {
			seed = randomNumbers(seed)
		}
		total += seed
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

// func main() {
// 	fmt.Println(sumOfTwoThousandthNumber())
// }
