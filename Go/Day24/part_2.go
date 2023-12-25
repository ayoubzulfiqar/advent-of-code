package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// type Hailstone struct {
// 	px, py, pz, vx, vy, vz int
// }

var hailstones []Hailstone
var velocitiesX = make(map[int][]int)
var velocitiesY = make(map[int][]int)
var velocitiesZ = make(map[int][]int)

func getRockVelocity(velocities map[int][]int) int {
	possibleV := make([]int, 0)
	for x := -1000; x <= 1000; x++ {
		possibleV = append(possibleV, x)
	}

	for vel, values := range velocities {
		if len(values) < 2 {
			continue
		}

		newPossibleV := make([]int, 0)
		for _, possible := range possibleV {
			// Add a check to ensure that the denominator is not zero
			if possible-vel != 0 && (values[0]-values[1])%(possible-vel) == 0 {
				newPossibleV = append(newPossibleV, possible)
			}
		}

		possibleV = newPossibleV
	}

	return possibleV[0]
}

func CoordinatesOfInitialPosition() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " @ ")
		positions := parts[0]
		velocity := parts[1]

		pos := strings.Split(positions, ", ")
		px, _ := strconv.Atoi(pos[0])
		py, _ := strconv.Atoi(pos[1])
		pz, _ := strconv.Atoi(pos[2])

		vel := strings.Split(velocity, ", ")
		vx, _ := strconv.Atoi(vel[0])
		vy, _ := strconv.Atoi(vel[1])
		vz, _ := strconv.Atoi(vel[2])

		if _, ok := velocitiesX[vx]; !ok {
			velocitiesX[vx] = []int{px}
		} else {
			velocitiesX[vx] = append(velocitiesX[vx], px)
		}

		if _, ok := velocitiesY[vy]; !ok {
			velocitiesY[vy] = []int{py}
		} else {
			velocitiesY[vy] = append(velocitiesY[vy], py)
		}

		if _, ok := velocitiesZ[vz]; !ok {
			velocitiesZ[vz] = []int{pz}
		} else {
			velocitiesZ[vz] = append(velocitiesZ[vz], pz)
		}

		hailstones = append(hailstones, Hailstone{px, py, pz, vx, vy, vz})
	}

	possibleVX := make([]int, 2001) // Create a slice with length 2001
	for x := -1000; x <= 1000; x++ {
		possibleVX[x+1000] = x
	}

	rvx := getRockVelocity(velocitiesX)
	rvy := getRockVelocity(velocitiesY)
	rvz := getRockVelocity(velocitiesZ)

	results := make(map[int]int)
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			stoneA := hailstones[i]
			stoneB := hailstones[j]

			ma := float64(stoneA.vy-rvy) / float64(stoneA.vx-rvx)
			mb := float64(stoneB.vy-rvy) / float64(stoneB.vx-rvx)

			ca := float64(stoneA.py) - ma*float64(stoneA.px)
			cb := float64(stoneB.py) - mb*float64(stoneB.px)

			rpx := int((cb - ca) / (ma - mb))
			rpy := int(ma*float64(rpx) + ca)

			time := int((rpx - stoneA.px) / int(float64(stoneA.vx-rvx)))
			rpz := stoneA.pz + (stoneA.vz-rvz)*time

			result := rpx + rpy + rpz
			if _, ok := results[result]; !ok {
				results[result] = 1
			} else {
				results[result]++
			}
		}
	}

	var keys []int
	for k := range results {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return results[keys[i]] > results[keys[j]]
	})

	fmt.Println(keys[0])
}
