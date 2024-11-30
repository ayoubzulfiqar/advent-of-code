package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hashed(s string) int {
	v := 0

	for _, ch := range s {
		v += int(ch)
		v *= 17
		v %= 256
	}

	return v
}

func PowerOfResultLens() {
	boxes := make([][]string, 256)
	focalLengths := make(map[string]int)
	file, _ := os.ReadFile("input.txt")

	instructions := strings.Split(string(file), ",")

	for _, instruction := range instructions {
		if strings.Contains(instruction, "-") {
			label := instruction[:len(instruction)-1]
			index := hashed(label)

			for i, l := range boxes[index] {
				if l == label {
					boxes[index] = append(boxes[index][:i], boxes[index][i+1:]...)
					break
				}
			}
		} else {
			parts := strings.Split(instruction, "=")
			label := parts[0]
			length := parts[1]

			lengthValue, err := strconv.Atoi(length)
			if err != nil {
				fmt.Println("Error parsing length:", err)
				return
			}

			index := hashed(label)
			if !contains(boxes[index], label) {
				boxes[index] = append(boxes[index], label)
			}

			focalLengths[label] = lengthValue
		}
	}

	total := 0

	for boxNumber, box := range boxes {
		for lensSlot, label := range box {
			total += (boxNumber + 1) * (lensSlot + 1) * focalLengths[label]
		}
	}

	fmt.Println(total)
}

func contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
