package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CubeConundrumPower() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	power := 0
	var bufferCount int

	for scanner.Scan() {
		var redCount, greenCount, blueCount int = 0, 0, 0
		lines := scanner.Text()

		for j := 0; j < len(lines); j++ {
			if lines[j] >= '0' && lines[j] <= '9' {
				bufferCount *= 10
				bufferCount += int(lines[j] - '0')
				continue
			}
			if lines[j] == ' ' {
				continue
			}
			switch lines[j] {
			case 'r':
				redCount = max(redCount, bufferCount)
			case 'g':
				greenCount = max(greenCount, bufferCount)
			case 'b':
				blueCount = max(blueCount, bufferCount)
			}

			bufferCount = 0
		}

		power += redCount * greenCount * blueCount
	}

	fmt.Println(power)
	return power
}
