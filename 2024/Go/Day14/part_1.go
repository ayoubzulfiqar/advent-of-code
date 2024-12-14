package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bot struct {
	x  int
	y  int
	dx int
	dy int
}

// const (
// 	maxX = 101
// 	maxY = 103
// )

func HundredSecondSafetyFactor() {
	var bots []Bot

	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the lines from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var b Bot
		// Parse the line into Bot structure (updated format: p=<x,y> v=<dx,dy>)
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &b.x, &b.y, &b.dx, &b.dy)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		bots = append(bots, b)
	}

	// Move the bots for 100 iterations
	for i := 0; i < 100; i++ {
		for j := 0; j < len(bots); j++ {
			bots[j].x += bots[j].dx
			bots[j].y += bots[j].dy

			// Wrap around the edges
			if bots[j].x < 0 {
				bots[j].x += maxX
			} else if bots[j].x >= maxX {
				bots[j].x -= maxX
			}

			if bots[j].y < 0 {
				bots[j].y += maxY
			} else if bots[j].y >= maxY {
				bots[j].y -= maxY
			}
		}
	}

	// Count how many bots are in each quadrant
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, bot := range bots {
		if bot.x < maxX/2 && bot.y < maxY/2 {
			q1++
		}
		if bot.x > maxX/2 && bot.y < maxY/2 {
			q2++
		}
		if bot.x < maxX/2 && bot.y > maxY/2 {
			q3++
		}
		if bot.x > maxX/2 && bot.y > maxY/2 {
			q4++
		}
	}

	// Output the result
	out := q1 * q2 * q3 * q4
	fmt.Println(out)
}
