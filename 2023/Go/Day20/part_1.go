// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// // type Type int

// // const (
// // 	None Type = iota
// // 	Flip
// // 	Conj
// // )

// func LowImpulseMultiplyHighImpulse() {
// 	dst := make(map[string][]string)
// 	src := make(map[string][]string)
// 	mem := make(map[string]map[string]int)
// 	state := make(map[string]int)
// 	t := make(map[string]Type)
// 	file, _ := os.Open("input.txt")
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		curType := None
// 		if strings.HasPrefix(line, "%") {
// 			curType = Flip
// 			line = line[1:]
// 		}
// 		if strings.HasPrefix(line, "&") {
// 			curType = Conj
// 			line = line[1:]
// 		}

// 		temp := strings.Split(line, " -> ")
// 		name := temp[0]
// 		right := strings.Split(temp[1], ", ")

// 		t[name] = curType
// 		for _, cAND := range right {
// 			dst[name] = append(dst[name], cAND)
// 			src[cAND] = append(src[cAND], name)
// 			if mem[cAND] == nil {
// 				mem[cAND] = make(map[string]int)
// 			}
// 			mem[cAND][name] = 0
// 		}
// 		state[name] = 0
// 	}

// 	v := [2]int{}
// 	for step := 0; step < 1000; step++ {
// 		q := make([][3]interface{}, 0)

// 		push := func(name, cAND string, value int) {
// 			q = append(q, [3]interface{}{name, cAND, value})
// 			v[value]++
// 		}

// 		push("button", "broadcaster", 0)
// 		for len(q) > 0 {
// 			entry := q[0]
// 			q = q[1:]

// 			prev := entry[0].(string)
// 			name := entry[1].(string)
// 			value := entry[2].(int)

// 			if _, ok := t[name]; !ok {
// 				continue
// 			}
// 			if t[name] == Flip && value == 1 {
// 				continue
// 			}
// 			if t[name] == Flip && value == 0 {
// 				state[name] ^= 1
// 				value = state[name]
// 			}
// 			if t[name] == Conj {
// 				mem[name][prev] = value
// 				value = 0
// 				for _, memValue := range mem[name] {
// 					if memValue == 0 {
// 						value = 1
// 						break
// 					}
// 				}
// 			}
// 			for _, cAND := range dst[name] {
// 				push(name, cAND, value)
// 			}
// 		}
// 	}

// 	result := v[0] * v[1]
// 	fmt.Println(result)
// }

package main
