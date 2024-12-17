package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ThreeBitMachine struct {
	registerA          int
	registerB          int
	registerC          int
	program            []int
	instructionPointer int
	outputValues       []int
}

func (tm *ThreeBitMachine) adv(operand int) {
	numerator := tm.registerA
	denominator := 1 << tm.getComboOperandValue(operand)
	tm.registerA = numerator / denominator
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) bdv(operand int) {
	numerator := tm.registerA
	denominator := 1 << tm.getComboOperandValue(operand)
	tm.registerB = numerator / denominator
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) cdv(operand int) {
	numerator := tm.registerA
	denominator := 1 << tm.getComboOperandValue(operand)
	tm.registerC = numerator / denominator
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) bxl(operand int) {
	tm.registerB ^= operand
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) bst(operand int) {
	tm.registerB = tm.getComboOperandValue(operand) % 8
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) jnz(operand int) {
	if tm.registerA == 0 {
		tm.incrementPointer()
		return
	}
	tm.instructionPointer = operand
}

func (tm *ThreeBitMachine) bxc(operand int) {
	tm.registerB ^= tm.registerC
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) out(operand int) {
	operandValue := tm.getComboOperandValue(operand)
	tm.outputValues = append(tm.outputValues, operandValue%8)
	tm.incrementPointer()
}

func (tm *ThreeBitMachine) incrementPointer() {
	tm.instructionPointer += 2
}

func (tm *ThreeBitMachine) getComboOperandValue(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}
	switch operand {
	case 4:
		return tm.registerA
	case 5:
		return tm.registerB
	case 6:
		return tm.registerC
	}
	panic("not a valid program")
}

func (tm *ThreeBitMachine) runInstruction(opcode int, operand int) {
	switch opcode {
	case 0:
		tm.adv(operand)
	case 6:
		tm.bdv(operand)
	case 7:
		tm.cdv(operand)
	case 1:
		tm.bxl(operand)
	case 2:
		tm.bst(operand)
	case 3:
		tm.jnz(operand)
	case 4:
		tm.bxc(operand)
	case 5:
		tm.out(operand)
	default:
		panic("not a valid opcode")
	}
}

func (tm *ThreeBitMachine) getOutput() {
	var output []string
	for _, value := range tm.outputValues {
		output = append(output, fmt.Sprintf("%d", value))
	}
	fmt.Println(strings.Join(output, ","))
}

func (tm *ThreeBitMachine) runProgram() {
	for tm.instructionPointer < len(tm.program)-1 {
		opcode := tm.program[tm.instructionPointer]
		operand := tm.program[tm.instructionPointer+1]
		tm.runInstruction(opcode, operand)
	}
	tm.getOutput()
}

func readInputFile(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}

func joinCommaOutPutString() {
	lines := readInputFile("input.txt")
	registerA, _ := strconv.Atoi(lines[0][12:])
	registerB, _ := strconv.Atoi(lines[1][12:])
	registerC, _ := strconv.Atoi(lines[2][12:])
	programStr := lines[4][9:]
	programStrs := strings.Split(programStr, ",")
	program := make([]int, len(programStrs))
	for i, str := range programStrs {
		program[i], _ = strconv.Atoi(str)
	}

	machine := &ThreeBitMachine{
		registerA:          registerA,
		registerB:          registerB,
		registerC:          registerC,
		program:            program,
		instructionPointer: 0,
	}

	machine.runProgram()
}
