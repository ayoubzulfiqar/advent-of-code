package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	StartID uint64
	EndID   uint64
}

func parseRanges(input string) []IDRange {
	var ranges []IDRange
	linePairs := strings.Split(input, ",")

	for _, pair := range linePairs {
		ids := strings.Split(pair, "-")
		if len(ids) != 2 {
			continue
		}

		startID, err1 := strconv.ParseUint(strings.TrimSpace(ids[0]), 10, 64)
		endID, err2 := strconv.ParseUint(strings.TrimSpace(ids[1]), 10, 64)

		if err1 == nil && err2 == nil {
			ranges = append(ranges, IDRange{StartID: startID, EndID: endID})
		}
	}

	return ranges
}

func countDigits(num uint64) uint32 {
	if num == 0 {
		return 1
	}
	return uint32(math.Log10(float64(num))) + 1
}

func isDivisibleBy(dividend, divisor uint32) bool {
	return dividend%divisor == 0
}

func findInvalidIDs(startID, endID uint64) []uint64 {
	var invalidIDs []uint64

	startDigits := countDigits(startID)
	endDigits := countDigits(endID)

	var minHalfDigits uint32
	if isDivisibleBy(startDigits, 2) {
		minHalfDigits = startDigits / 2
		if minHalfDigits < 1 {
			minHalfDigits = 1
		}
	} else {
		minHalfDigits = startDigits/2 + 1
	}
	maxHalfDigits := endDigits / 2

	for halfLength := minHalfDigits; halfLength <= maxHalfDigits; halfLength++ {
		firstHalfLowerBound := uint64(math.Pow10(int(halfLength - 1)))
		if firstHalfLowerBound < startID/uint64(math.Pow10(int(halfLength))) {
			firstHalfLowerBound = startID / uint64(math.Pow10(int(halfLength)))
		}

		firstHalfUpperBound := uint64(math.Pow10(int(halfLength))) - 1
		if firstHalfUpperBound > endID/uint64(math.Pow10(int(halfLength))) {
			firstHalfUpperBound = endID / uint64(math.Pow10(int(halfLength)))
		}

		if firstHalfLowerBound > firstHalfUpperBound {
			continue
		}

		for firstHalf := firstHalfLowerBound; firstHalf <= firstHalfUpperBound; firstHalf++ {
			repeatingID := firstHalf*uint64(math.Pow10(int(halfLength))) + firstHalf

			if repeatingID >= startID && repeatingID <= endID {
				invalidIDs = append(invalidIDs, repeatingID)
			}
		}
	}

	return invalidIDs
}

func calculateInvalidIDSum(idRanges []IDRange) uint64 {
	var totalSum uint64 = 0

	for _, idRange := range idRanges {
		invalidIDs := findInvalidIDs(idRange.StartID, idRange.EndID)
		for _, id := range invalidIDs {
			totalSum += id
		}
	}

	return totalSum
}

func SolvePart1() {
	data, readError := os.ReadFile("input.txt")
	if readError != nil {
		panic(readError)
	}
	idRanges := parseRanges(string(data))
	solution := calculateInvalidIDSum(idRanges)
	fmt.Printf("Part 1 Result: %d\n", solution)
}
