package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CubicMetersOfLagoonLava() {
	inp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inp.Close()

	scanner := bufio.NewScanner(inp)
	var pos, nextPos, posCorr, nextPosCorr [2]int
	var sum1, sum2, sumDir, sum1Corr, sum2Corr, sumDirCorr int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		d, n, col := parts[0], parts[1], parts[2]

		var direction, directionCorr [2]int

		switch d {
		case "L":
			direction = [2]int{-1, 0}
		case "R":
			direction = [2]int{1, 0}
		case "U":
			direction = [2]int{0, -1}
		case "D":
			direction = [2]int{0, 1}
		}

		length, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		nextPos[0] = pos[0] + direction[0]*length
		nextPos[1] = pos[1] + direction[1]*length

		sum1 += pos[0] * nextPos[1]
		sum2 += pos[1] * nextPos[0]
		sumDir += length
		pos = nextPos

		cleaned := strings.TrimSuffix(strings.TrimPrefix(col, "(#"), ")")
		corrLen, err := strconv.ParseInt(cleaned[:5], 16, 64)
		if err != nil {
			panic(err)
		}

		switch cleaned[5] {
		case '0':
			directionCorr = [2]int{1, 0}
		case '1':
			directionCorr = [2]int{0, 1}
		case '2':
			directionCorr = [2]int{-1, 0}
		case '3':
			directionCorr = [2]int{0, -1}
		}

		nextPosCorr[0] = posCorr[0] + directionCorr[0]*int(corrLen)
		nextPosCorr[1] = posCorr[1] + directionCorr[1]*int(corrLen)

		sum1Corr += posCorr[0] * nextPosCorr[1]
		sum2Corr += posCorr[1] * nextPosCorr[0]
		sumDirCorr += int(corrLen)
		posCorr = nextPosCorr
	}

	// area := abs(sum1-sum2) / 2
	// fmt.Println(area + sumDir/2 + 1)

	areaCorr := abs(sum1Corr-sum2Corr) / 2
	fmt.Println("Part-2:", areaCorr+sumDirCorr/2+1)
}
