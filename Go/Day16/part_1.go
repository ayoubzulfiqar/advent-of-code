package main

import (
	"bufio"
	"os"
)

type Line struct {
	Index int
	Text  string
}

// GetInputLines returns 1 line at a time when looped over
func GetInputLines(filename string) (channel chan Line) {
	channel = make(chan Line, 1)
	currentLine := 0
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			close(channel)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for {
			if scanner.Scan() {
				channel <- Line{currentLine, scanner.Text()}
				currentLine++
			} else {
				close(channel)
				return
			}
		}

	}()
	return channel
}

const (
	TOP = iota
	RIGHT
	BOTTOM
	LEFT
)

var LastCol int
var LastRow int

type Beam struct {
}

type Object struct {
	Char          rune
	Loc           Location
	EnergizedFrom int
}

// going a little different this time
// putting Location in a map rather than 2D array
type Location struct {
	X int
	Y int
}

func (o *Object) getNext(incomingDirection int) (outboundLocations []Location) {
	switch incomingDirection {
	case TOP:
		switch o.Char {
		case '|':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y + 1})
		case '-':
			outboundLocations = append(outboundLocations, Location{o.Loc.X - 1, o.Loc.Y})
			outboundLocations = append(outboundLocations, Location{o.Loc.X + 1, o.Loc.Y})
		case '/':
			outboundLocations = append(outboundLocations, Location{o.Loc.X - 1, o.Loc.Y})
		case '\\':
			outboundLocations = append(outboundLocations, Location{o.Loc.X + 1, o.Loc.Y})
		default:
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y + 1})
		}
	case RIGHT:
		switch o.Char {
		case '|':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y - 1})
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y + 1})
		case '-':
			outboundLocations = append(outboundLocations, Location{o.Loc.X - 1, o.Loc.Y})
		case '/':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y + 1})
		case '\\':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y - 1})
		default:
			outboundLocations = append(outboundLocations, Location{o.Loc.X - 1, o.Loc.Y})
		}
	case BOTTOM:
		switch o.Char {
		case '|':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y - 1})
		case '-':
			outboundLocations = append(outboundLocations, Location{o.Loc.X - 1, o.Loc.Y})
			outboundLocations = append(outboundLocations, Location{o.Loc.X + 1, o.Loc.Y})
		case '/':
			outboundLocations = append(outboundLocations, Location{o.Loc.X + 1, o.Loc.Y})
		case '\\':
			outboundLocations = append(outboundLocations, Location{o.Loc.X - 1, o.Loc.Y})
		default:
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y - 1})
		}
	default:
		switch o.Char {
		case '|':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y - 1})
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y + 1})
		case '-':
			outboundLocations = append(outboundLocations, Location{o.Loc.X + 1, o.Loc.Y})
		case '/':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y - 1})
		case '\\':
			outboundLocations = append(outboundLocations, Location{o.Loc.X, o.Loc.Y + 1})
		default:
			outboundLocations = append(outboundLocations, Location{o.Loc.X + 1, o.Loc.Y})
		}
	}
	return outboundLocations
}

func parseFile(filename string) map[Location]*Object {
	objectMap := make(map[Location]*Object)
	for l := range GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		if LastCol == 0 {
			LastCol = len(line) - 1
		}
		for x, char := range line {
			objectMap[Location{X: x, Y: l.Index}] = &Object{Char: char, Loc: Location{X: x, Y: l.Index}, EnergizedFrom: -1}
		}
		LastRow = l.Index
	}
	return objectMap
}

func traverse(objectMap map[Location]*Object, location Location, direction int) {
	object, ok := objectMap[location]
	if !ok {
		return
	}
	if object.EnergizedFrom == direction {
		return
	}
	object.EnergizedFrom = direction
	nextLocations := object.getNext(direction)
	for _, nextLocation := range nextLocations {
		if nextLocation.X > location.X {
			traverse(objectMap, nextLocation, LEFT)
		} else if nextLocation.X < location.X {
			traverse(objectMap, nextLocation, RIGHT)
		} else if nextLocation.Y > location.Y {
			traverse(objectMap, nextLocation, TOP)
		} else {
			traverse(objectMap, nextLocation, BOTTOM)
		}
	}
}

func countEnergized(objectMap map[Location]*Object) (count int) {
	for loc, obj := range objectMap {
		if obj.EnergizedFrom >= 0 {
			count++
		}
		objectMap[loc].EnergizedFrom = -1
	}
	return count
}

func Run() int {
	answer := 0
	objectMap := parseFile("input.txt")
	traverse(objectMap, Location{0, 0}, LEFT)
	answer = countEnergized(objectMap)
	// 7472
	println(answer)
	return answer
}
