package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScratchcardsWorth() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var t int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		x := parts[1]
		x = strings.TrimSpace(x)
		arrays := strings.Split(x, " | ")
		if len(arrays) != 2 {
			continue
		}

		a := parseIntArray(arrays[0])
		b := parseIntArray(arrays[1])

		j := 0
		for _, q := range b {
			for _, value := range a {
				if q == value {
					j++
					break
				}
			}
		}

		if j > 0 {
			t += 1 << (j - 1)
		}
	}

	fmt.Println(t)
}

func parseIntArray(s string) []int {
	var result []int
	parts := strings.Fields(s)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}
