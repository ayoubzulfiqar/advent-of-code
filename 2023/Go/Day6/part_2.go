package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func BeatTheLongerRace() int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	race := parseRace(string(file))
	wins := race.countWins()
	fmt.Println(wins)
	return wins
}

func parseRace(input string) *Race {
	values := strings.Split(strings.TrimSpace(input), "\n")

	time := strings.Join(strings.Fields(values[0])[1:], "")
	distance := strings.Join(strings.Fields(values[1])[1:], "")

	timeInt, err := strconv.Atoi(time)
	if err != nil {
		log.Fatal(err)
	}

	distanceInt, err := strconv.Atoi(distance)
	if err != nil {
		log.Fatal(err)
	}

	return &Race{time: timeInt, distance: distanceInt}
}

// type Race struct {
// 	time     int
// 	distance int
// }

func (r *Race) countWins() int {
	time := float64(r.time)
	distance := float64(r.distance)

	x1 := -0.5*math.Sqrt(math.Pow(time, 2)-4*distance) - time
	x2 := 0.5*math.Sqrt(math.Pow(time, 2)-4*distance) - time

	return int(math.Abs(math.Floor(x2) - math.Floor(x1)))
}
