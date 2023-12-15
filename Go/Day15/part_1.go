package main

import (
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
	v := 0

	for _, ch := range s {
		v += int(ch)
		v *= 17
		v %= 256
	}

	return v
}

func HashSumOfResults() {
	input, _ := os.ReadFile("input.txt")
	inputList := strings.Split(string(input), ",")

	var result int
	for _, str := range inputList {
		result += hash(str)
	}

	fmt.Println(result)
}
