package main

import (
	"bufio"
	"fmt"
	"os"
)

// sum: 3059 power: 65371

// func CubeConundrum() int {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return 0
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	sum, powers := 0, 0
// 	for scanner.Scan() {
// 		game := strings.Split(scanner.Text(), ": ")
// 		mins := map[string]int{}
// 		for _, show := range strings.Split(game[1], ";") {
// 			for _, c := range strings.Split(show, ",") {
// 				cubes := strings.Split(strings.TrimSpace(c), " ")
// 				n, _ := strconv.ParseInt(cubes[0], 10, 32)
// 				mins[cubes[1]] = max(mins[cubes[1]], int(n))
// 			}
// 		}
// 		if mins["red"] <= 12 && mins["green"] <= 13 && mins["blue"] <= 14 {
// 			n, _ := strconv.ParseInt(strings.Split(game[0], " ")[1], 10, 32)
// 			sum += int(n)
// 		}
// 		powers += mins["red"] * mins["green"] * mins["blue"]
// 	}
// 	fmt.Println(sum, powers)
// 	return sum

// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func CubeConundrum() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	var bufferCount int

	for i := 0; scanner.Scan(); i++ {
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
				if bufferCount > 12 {
					goto exit
				}
			case 'g':
				if bufferCount > 13 {
					goto exit
				}
			case 'b':
				if bufferCount > 14 {
					goto exit
				}
			}

			bufferCount = 0
		}
		total += i + 1
	exit:
	}

	fmt.Println(total)
	return total
}
