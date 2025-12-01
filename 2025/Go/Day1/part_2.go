package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TurnDirection interface {
	isTurnDirection()
	getAmount() int
	withAmount(int) TurnDirection
}

type LeftTurn struct{ amount int }
type RightTurn struct{ amount int }

func (l LeftTurn) getAmount() int                  { return l.amount }
func (r RightTurn) getAmount() int                 { return r.amount }
func (LeftTurn) isTurnDirection()                  {}
func (RightTurn) isTurnDirection()                 {}
func (l LeftTurn) withAmount(a int) TurnDirection  { return LeftTurn{amount: a} }
func (r RightTurn) withAmount(a int) TurnDirection { return RightTurn{amount: a} }

func parseDirection(s string) (TurnDirection, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("empty string")
	}

	amount, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid number: %s", s[1:])
	}

	switch s[0] {
	case 'L':
		return LeftTurn{amount: amount}, nil
	case 'R':
		return RightTurn{amount: amount}, nil
	default:
		return nil, fmt.Errorf("malformed input: expected L<n> or R<n>")
	}
}

func applyTurn(pointer, zeroCount int, direction TurnDirection) (int, int) {
	return processRecursively(pointer, zeroCount, direction)
}

func processRecursively(p, z int, d TurnDirection) (int, int) {
	// Base case: if turn amount is 0, return current state
	if d.getAmount() == 0 {
		return p, z
	}

	// Handle the recursive step
	var nextPointer int
	var nextZeros int
	var nextDirection TurnDirection

	switch dir := d.(type) {
	case LeftTurn:
		if p == 0 {
			nextPointer = (p - 1) % 100
			if nextPointer < 0 {
				nextPointer += 100
			}
			nextZeros = z + 1
		} else {
			nextPointer = (p - 1) % 100
			if nextPointer < 0 {
				nextPointer += 100
			}
			nextZeros = z
		}
		nextDirection = dir.withAmount(dir.amount - 1)
	case RightTurn:
		if p == 0 {
			nextPointer = (p + 1) % 100
			nextZeros = z + 1
		} else {
			nextPointer = (p + 1) % 100
			nextZeros = z
		}
		nextDirection = dir.withAmount(dir.amount - 1)
	}

	return processRecursively(nextPointer, nextZeros, nextDirection)
}

func computeResult(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	directions := make([]TurnDirection, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		dir, err := parseDirection(line)
		if err != nil {
			return 0, err
		}
		directions = append(directions, dir)
	}

	pointer := 50
	zeroCount := 0

	for _, dir := range directions {
		pointer, zeroCount = applyTurn(pointer, zeroCount, dir)
	}

	return zeroCount, nil
}

func Method0x434C49434BToOpenTheDoor() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	output, err := computeResult(string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("result = %d\n", output)
}
