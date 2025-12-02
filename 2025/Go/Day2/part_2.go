package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MultiRepeatRange struct {
	LowerBound uint64
	UpperBound uint64
}

func parseMultiRepeatRanges(inputData string) []MultiRepeatRange {
	var rangeCollection []MultiRepeatRange
	pairStrings := strings.Split(inputData, ",")

	for _, rangeStr := range pairStrings {
		bounds := strings.Split(rangeStr, "-")
		if len(bounds) != 2 {
			continue
		}

		lowerValue, parseErr1 := strconv.ParseUint(strings.TrimSpace(bounds[0]), 10, 64)
		upperValue, parseErr2 := strconv.ParseUint(strings.TrimSpace(bounds[1]), 10, 64)

		if parseErr1 == nil && parseErr2 == nil {
			rangeCollection = append(rangeCollection, MultiRepeatRange{LowerBound: lowerValue, UpperBound: upperValue})
		}
	}

	return rangeCollection
}

func digitCount(value uint64) uint32 {
	if value == 0 {
		return 1
	}
	return uint32(math.Log10(float64(value))) + 1
}

func isEvenlyDivisible(dividend, divisor uint32) bool {
	return dividend%divisor == 0
}

func discoverMultiRepeatingIDs(minID, maxID uint64) map[uint64]struct{} {
	foundIDs := make(map[uint64]struct{})

	minDigits := digitCount(minID)
	maxDigits := digitCount(maxID)

	for repeatTimes := uint32(2); repeatTimes <= maxDigits; repeatTimes++ {
		var segmentMinLength uint32
		if isEvenlyDivisible(minDigits, repeatTimes) {
			segmentMinLength = minDigits / repeatTimes
			if segmentMinLength < 1 {
				segmentMinLength = 1
			}
		} else {
			segmentMinLength = minDigits/repeatTimes + 1
		}
		segmentMaxLength := maxDigits / repeatTimes

		for segmentLength := segmentMinLength; segmentLength <= segmentMaxLength; segmentLength++ {
			minSegmentValue := uint64(math.Pow10(int(segmentLength - 1)))
			if minSegmentValue < minID/uint64(math.Pow10(int(segmentLength*(repeatTimes-1)))) {
				minSegmentValue = minID / uint64(math.Pow10(int(segmentLength*(repeatTimes-1))))
			}

			maxSegmentValue := uint64(math.Pow10(int(segmentLength))) - 1
			if maxSegmentValue > maxID/uint64(math.Pow10(int(segmentLength*(repeatTimes-1)))) {
				maxSegmentValue = maxID / uint64(math.Pow10(int(segmentLength*(repeatTimes-1))))
			}

			if minSegmentValue > maxSegmentValue {
				continue
			}

			for segmentNumber := minSegmentValue; segmentNumber <= maxSegmentValue; segmentNumber++ {
				constructedID := segmentNumber
				for repetition := uint32(1); repetition < repeatTimes; repetition++ {
					constructedID = constructedID*uint64(math.Pow10(int(segmentLength))) + segmentNumber
				}

				if constructedID >= minID && constructedID <= maxID {
					foundIDs[constructedID] = struct{}{}
				}
			}
		}
	}

	return foundIDs
}

func computeMultiRepeatIDSum(ranges []MultiRepeatRange) uint64 {
	var cumulativeSum uint64 = 0

	for _, currentRange := range ranges {
		multiRepeatIDs := discoverMultiRepeatingIDs(currentRange.LowerBound, currentRange.UpperBound)
		for id := range multiRepeatIDs {
			cumulativeSum += id
		}
	}

	return cumulativeSum
}

func SolvePart2() {
	fileData, fileError := os.ReadFile("input.txt")
	if fileError != nil {
		panic(fileError)
	}
	rangeCollection := parseMultiRepeatRanges(string(fileData))
	totalResult := computeMultiRepeatIDSum(rangeCollection)
	fmt.Printf("Part 2 Result: %d\n", totalResult)
}
