package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction interface {
	isDirection()
}

type L struct{ n int }
type R struct{ n int }

func (L) isDirection() {}
func (R) isDirection() {}

func parseTurn(s string) (Direction, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("empty string")
	}

	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid number: %s", s[1:])
	}

	switch s[0] {
	case 'L':
		return L{n: n}, nil
	case 'R':
		return R{n: n}, nil
	default:
		return nil, fmt.Errorf("malformed input: expected L<n> or R<n>")
	}
}

func step(dial, zeros int, turn Direction) (int, int) {
	var dialPrime int

	switch t := turn.(type) {
	case L:
		dialPrime = (dial - t.n) % 100
		if dialPrime < 0 {
			dialPrime += 100
		}
	case R:
		dialPrime = (dial + t.n) % 100
	}

	zerosPrime := zeros
	if dialPrime == 0 {
		zerosPrime++
	}

	return dialPrime, zerosPrime
}

func process(content string) (int, error) {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	turns := make([]Direction, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		turn, err := parseTurn(line)
		if err != nil {
			return 0, err
		}
		turns = append(turns, turn)
	}

	dial := 50
	zeros := 0

	for _, turn := range turns {
		dial, zeros = step(dial, zeros, turn)
	}

	return zeros, nil
}

func ActualPasswordOfTheDoor() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	result, err := process(string(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("result = %d\n", result)
}
