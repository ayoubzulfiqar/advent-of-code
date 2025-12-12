package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RegionsCanFitPresents() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	readingPatterns := true
	var patterns []int
	var regionLines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if readingPatterns && builder.Len() > 0 {
				patterns = append(patterns, strings.Count(builder.String(), "#"))
				builder.Reset()
			}
			continue
		}

		if readingPatterns {
			if strings.Contains(line, ": ") {
				readingPatterns = false
				if builder.Len() > 0 {
					patterns = append(patterns, strings.Count(builder.String(), "#"))
					builder.Reset()
				}
				regionLines = append(regionLines, line)
			} else {
				if builder.Len() > 0 {
					builder.WriteString("\n")
				}
				builder.WriteString(line)
			}
		} else {
			regionLines = append(regionLines, line)
		}
	}

	if readingPatterns && builder.Len() > 0 {
		patterns = append(patterns, strings.Count(builder.String(), "#"))
	}

	var regions int = 0

	for _, line := range regionLines {
		colonIndex := strings.Index(line, ": ")
		if colonIndex == -1 {
			continue
		}

		areaStr := line[:colonIndex]
		numsStr := line[colonIndex+2:]

		xIndex := strings.Index(areaStr, "x")
		if xIndex == -1 {
			continue
		}

		width, err1 := strconv.Atoi(areaStr[:xIndex])
		height, err2 := strconv.Atoi(areaStr[xIndex+1:])
		if err1 != nil || err2 != nil {
			continue
		}

		area := width * height

		fields := strings.Fields(numsStr)
		size := 0

		for i, field := range fields {
			if i >= len(patterns) {
				break
			}
			if num, err := strconv.Atoi(field); err == nil {
				size += patterns[i] * num
			}
		}

		if size > area {
			continue
		}

		if float64(size)*(1.2) < float64(area) {
			regions++
		}
	}
	fmt.Printf("Regions: %d", regions)
	return regions
}

// func main() {
// 	RegionsCanFitPresents()

// }
