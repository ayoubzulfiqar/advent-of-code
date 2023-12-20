package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Module struct {
	Name    string
	Type    string
	Outputs []string
	Memory  interface{}
}

func NewModule(name, typeStr string, outputs []string) *Module {
	var memory interface{}
	if typeStr == "%" {
		memory = "off"
	} else {
		memory = make(map[string]string)
	}
	return &Module{Name: name, Type: typeStr, Outputs: outputs, Memory: memory}
}

func Run() {
	modules := make(map[string]*Module)
	var broadcastTargets []string
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		leftRight := strings.Split(line, " -> ")
		left, right := leftRight[0], leftRight[1]
		outputs := strings.Split(right, ", ")

		if left == "broadcaster" {
			broadcastTargets = outputs
		} else {
			typeStr := string(left[0])
			name := left[1:]
			modules[name] = NewModule(name, typeStr, outputs)
		}
	}

	var feed string
	for name, module := range modules {
		for _, output := range module.Outputs {
			if output == "rx" {
				feed = name
				break
			}
		}
		if feed != "" {
			break
		}
	}

	cycleLengths := make(map[string]int)
	seen := make(map[string]int)
	for name := range modules {
		if _, ok := seen[name]; !ok {
			seen[name] = 0
		}
	}

	presses := 0

	for {
		presses++
		q := make([][3]string, 0)
		for _, target := range broadcastTargets {
			q = append(q, [3]string{"broadcaster", target, "lo"})
		}

		for len(q) > 0 {
			origin, target, pulse := q[0][0], q[0][1], q[0][2]
			q = q[1:]

			module, ok := modules[target]
			if !ok {
				continue
			}

			if module.Name == feed && pulse == "hi" {
				seen[origin]++

				if cycleLength, ok := cycleLengths[origin]; !ok {
					cycleLengths[origin] = presses
				} else {
					if presses != seen[origin]*cycleLength {
						panic("Assertion failed")
					}
				}

				allSeen := true
				for _, v := range seen {
					if v == 0 {
						allSeen = false
						break
					}
				}
				if allSeen {
					x := 1
					for _, cycleLength := range cycleLengths {
						x = x * cycleLength / int(GCD((x), (cycleLength)))
					}
					fmt.Println(x)
					os.Exit(0)
				}
			}

			if module.Type == "%" {
				if pulse == "lo" {
					module.Memory = "on"
					if module.Memory == "off" {
						module.Memory = "on"
					} else {
						module.Memory = "off"
					}
					outgoing := "hi"
					if module.Memory == "on" {
						outgoing = "hi"
					} else {
						outgoing = "lo"
					}
					for _, x := range module.Outputs {
						q = append(q, [3]string{module.Name, x, outgoing})
					}
				}
			} else {
				module.Memory.(map[string]string)[origin] = pulse
				outgoing := "lo"
				allHi := true
				for _, x := range module.Memory.(map[string]string) {
					if x != "hi" {
						allHi = false
						break
					}
				}
				if allHi {
					outgoing = "hi"
				}
				for _, x := range module.Outputs {
					q = append(q, [3]string{module.Name, x, outgoing})
				}
			}
		}
	}
}
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
