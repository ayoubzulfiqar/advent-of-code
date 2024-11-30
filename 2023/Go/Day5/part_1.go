package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LowestLocationSeedNumber() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seeds []int
	var mappings []mapping
	currentMapping := mapping{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			seedsString := strings.Split(line, "seeds: ")
			seedList := strings.Split(seedsString[1], " ")
			for _, seedItem := range seedList {
				seed, err := strconv.Atoi(seedItem)
				if err != nil {
					fmt.Println("Error converting seed to integer:", err)
					return 0
				}
				seeds = append(seeds, seed)
			}
			continue
		}

		if strings.Contains(line, "-") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = mapping{}
			continue
		}

		values := strings.Fields(line)

		source, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Error converting source to integer:", err)
			return 0
		}

		size, err := strconv.Atoi(values[2])
		if err != nil {
			fmt.Println("Error converting size to integer:", err)
			return 0
		}

		offset := sti(values[0]) - sti(values[1])

		currentMapping = append(currentMapping, SubMapping{
			source: source,
			size:   size,
			offset: offset,
		})
	}

	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	lowest := -1
	for _, seed := range seeds {
		val := seed

		for _, mapping := range mappings {
			for _, subMapping := range mapping {
				if val >= subMapping.source && val <= subMapping.source+subMapping.size {
					val += subMapping.offset
					break
				}
			}
		}

		if lowest == -1 || val < lowest {
			lowest = val
		}
	}
	fmt.Println(lowest)
	return lowest
}

type mapping []SubMapping

type SubMapping struct {
	source, size, offset int
}

func sti(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	}
	return i
}
