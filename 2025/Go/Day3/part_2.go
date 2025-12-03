package main

import (
	"bufio"
	"fmt"
	"os"
)

func NewTotalOutputJoltage() uint64 {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var total uint64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		length := 0
		for i := 0; i < len(line) && line[i] >= '0' && line[i] <= '9'; i++ {
			length++
		}

		T := make([]uint64, length)
		R := make([]uint64, length)

		for i := 0; i < length; i++ {
			T[i] = uint64(line[i] - '0')
		}

		summands := 12
		for it := 0; it < summands; it++ {
			m := uint64(0)
			for i := 0; i < length; i++ {
				newVal := 10*m + T[i]
				if R[i] > m {
					m = R[i]
				}
				R[i] = newVal
			}
		}

		m := uint64(0)
		for i := 0; i < length; i++ {
			if R[i] > m {
				m = R[i]
			}
			R[i] = m
		}

		total += R[length-1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	return total
}
