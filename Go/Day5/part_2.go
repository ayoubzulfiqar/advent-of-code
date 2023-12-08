package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LowestInitialSeedNumber() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seedPairs [][2]int
	var mappings []Mapping

	currentMapping := Mapping{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			seedsString := strings.Split(line, "seeds: ")
			seedList := strings.Split(seedsString[1], " ")

			for i := 0; i < len(seedList); i += 2 {
				seedPairs = append(seedPairs, [2]int{sti(seedList[i]), sti(seedList[i]) + sti(seedList[i+1])})
			}
			continue
		}

		if strings.Contains(line, "-") {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = Mapping{}
			continue
		}

		values := strings.Fields(line)

		currentMapping = append(currentMapping, SubMapping{
			source: sti(values[1]),
			size:   sti(values[2]),
			offset: sti(values[0]) - sti(values[1]),
		})
	}
	if len(currentMapping) > 0 {
		mappings = append(mappings, currentMapping)
	}

	lowest := -1

	for _, pair := range seedPairs {
		ranges := [][2]int{pair}

		for _, mapping := range mappings {
			for _, subMapping := range mapping {
				ranges = splitRangesAt(ranges, subMapping.source)
			}

			for i := range ranges {
				ranges[i][0] = mapping.convert(ranges[i][0])
				ranges[i][1] = mapping.convert(ranges[i][1])
			}
		}

		for i := range ranges {
			if lowest == -1 {
				lowest = ranges[i][0]
			}

			lowest = min(lowest, ranges[i][0])
		}
	}
	fmt.Println(lowest)
	return lowest
}

type Mapping []SubMapping

func (m Mapping) convert(in int) int {
	for _, c := range m {
		if in >= c.source && in <= c.source+c.size {
			return in + c.offset
		}
	}
	return in
}

func splitRangesAt(s [][2]int, n int) [][2]int {
	for i, ss := range s {
		if n > ss[0] && n <= ss[1] {
			s[i][1] = n - 1
			s = append(s, [2]int{n, ss[1]})
			return s
		}
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
