package main

// blinkNTimes performs the blink operation iteratively for n times.
func blinkNTimes(iterations int) int {
	lines := readLinesToList()
	if len(lines) == 0 || len(lines[0]) == 0 {
		return 0
	}

	stones := make(map[int]int)
	for _, stone := range lines[0] {
		stones[stone]++
	}

	for i := 0; i < iterations; i++ {
		newStones := make(map[int]int)
		for rock, count := range stones {
			blinkResults := blink(rock)
			for _, blinkResult := range blinkResults {
				newStones[blinkResult] += count
			}
		}
		stones = newStones
	}

	total := 0
	for _, count := range stones {
		total += count
	}

	return total
}
