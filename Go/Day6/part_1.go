package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func DesertIslandNumbers() int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	game := parseRecords(string(file))
	r := game.result()
	fmt.Println(r)
	return r
}

type Game struct {
	records []Race
}

func (g *Game) result() int {
	wins := 1
	for _, record := range g.records {
		wins *= record.findWinningTimes()
	}
	return wins
}

func parseRecords(input string) *Game {
	values := strings.Split(strings.TrimSpace(input), "\n")

	times := strings.Fields(values[0])[1:]
	distances := strings.Fields(values[1])[1:]

	records := make([]Race, len(times))

	for i, time := range times {
		timeInt, err := strconv.Atoi(time)
		if err != nil {
			log.Fatal(err)
		}
		records[i].time = timeInt

		distanceInt, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal(err)
		}
		records[i].distance = distanceInt
	}

	game := Game{records: records}
	return &game
}

type Race struct {
	time     int
	distance int
}

func (r *Race) findWinningTimes() int {
	wins := 0
	for holdTime := 1; holdTime < r.time; holdTime++ {
		distance := r.travelDistance(holdTime)
		if distance > r.distance {
			wins += 1
		} else if wins > 0 {
			break
		}
	}
	return wins
}

func (r *Race) travelDistance(holdTime int) int {
	return holdTime * (r.time - holdTime)
}
