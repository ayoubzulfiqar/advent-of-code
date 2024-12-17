package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func combo(num int, registers map[string]int) int {
	if num <= 3 {
		return num
	} else if num == 4 {
		return registers["A"]
	} else if num == 5 {
		return registers["B"]
	} else if num == 6 {
		return registers["C"]
	}
	return 0
}

func adv(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	registers["A"] = registers["A"] / (1 << combo(num, registers))
	return registers, instr + 2, outputs
}

func bxl(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	registers["B"] ^= num
	return registers, instr + 2, outputs
}

func bst(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	registers["B"] = combo(num, registers) % 8
	return registers, instr + 2, outputs
}

func jnz(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	if registers["A"] == 0 {
		return registers, instr + 2, outputs
	} else {
		return registers, num, outputs
	}
}

func bxc(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	registers["B"] ^= registers["C"]
	return registers, instr + 2, outputs
}

func out(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	outputs = append(outputs, combo(num, registers)%8)
	return registers, instr + 2, outputs
}

func bdv(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	registers["B"] = registers["A"] / (1 << combo(num, registers))
	return registers, instr + 2, outputs
}

func cdv(num int, registers map[string]int, instr int, outputs []int) (map[string]int, int, []int) {
	registers["C"] = registers["A"] / (1 << combo(num, registers))
	return registers, instr + 2, outputs
}

func getOutput(a int, registersOriginal map[string]int, programs []int) []int {
	outputs := []int{}
	registers := make(map[string]int)
	for k, v := range registersOriginal {
		registers[k] = v
	}
	registers["A"] = a
	length := len(programs)
	instr := 0
	for instr < length {
		opcode := programs[instr]
		var function func(int, map[string]int, int, []int) (map[string]int, int, []int)
		switch opcode {
		case 0:
			function = adv
		case 1:
			function = bxl
		case 2:
			function = bst
		case 3:
			function = jnz
		case 4:
			function = bxc
		case 5:
			function = out
		case 6:
			function = bdv
		case 7:
			function = cdv
		}

		// Check if we can access the next index in the programs slice
		if instr+1 >= length {
			break
		}

		num := programs[instr+1]
		registers, instr, outputs = function(num, registers, instr, outputs)
	}
	return outputs
}

func lowestRegisterValueA() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var registersInput string
	var programsInput string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		registersInput += line + "\n"
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Program:") {
			programsInput = strings.TrimPrefix(line, "Program: ")
			break
		}
	}

	registersOriginal := make(map[string]int)
	registerLines := strings.Split(registersInput, "\n")
	for _, line := range registerLines {
		if line != "" {
			parts := strings.Split(line, ": ")
			registerName := strings.Split(parts[0], " ")[1]
			value, _ := strconv.Atoi(parts[1])
			registersOriginal[registerName] = value
		}
	}

	// Parse the program part (now correctly formatted)
	programs := []int{}
	programParts := strings.Split(programsInput, ",")
	for _, p := range programParts {
		num, _ := strconv.Atoi(strings.TrimSpace(p))
		programs = append(programs, num)
	}

	valid := []int{0}
	for length := 1; length < len(programs)+1; length++ {
		oldValid := valid
		valid = []int{}
		for _, num := range oldValid {
			for offset := 0; offset < 8; offset++ {
				newNum := 8*num + offset
				if fmt.Sprintf("%v", getOutput(newNum, registersOriginal, programs)) == fmt.Sprintf("%v", programs[len(programs)-length:]) {
					valid = append(valid, newNum)
				}
			}
		}
	}

	answer := valid[0]
	for _, v := range valid {
		if v < answer {
			answer = v
		}
	}
	fmt.Println(answer)
}
