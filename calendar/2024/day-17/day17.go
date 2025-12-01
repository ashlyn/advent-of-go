package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/slices"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(17, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	c, program := parseInput(input)
	c.runProgram(program)
	return strings.Join(slices.IntsToStrings(c.output), ",")
}

func solvePart2(input []string) int {
	c, program := parseInput(input)
	return findStartingValue(0, 0, c, program)
}

func findStartingValue(a, outputLength int, c computer, program []int) int {
	if outputLength == len(program) {
		return a
	}

	for i := 0; i < 8; i++ {
		aCandidate := (a * 8) + i
		c.reset(aCandidate)
		c.runProgram(program)

		// work backwords through the program to find the next base 8 "digit" to include
		if c.output[0] == program[len(program) - 1 - outputLength] {
			result := findStartingValue(aCandidate, outputLength + 1, c, program)
			if result != -1 {
				return result
			}
		}
	}

	return -1
}

func numberToDigits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append([]int{number % 10}, digits...)
		number /= 10
	}
	return digits
}

type computer struct {
	registers [3]int
	output []int
}

func parseInput(input []string) (computer, []int) {
	a, _ := strconv.Atoi(strings.Fields(input[0])[2])
	b, _ := strconv.Atoi(strings.Fields(input[1])[2])
	c, _ := strconv.Atoi(strings.Fields(input[2])[2])
	program := slices.ParseIntsFromStrings(strings.Split(strings.Fields(input[4])[1], ","))
	return computer{
		registers: [3]int{a, b, c},
		output: []int{},
	}, program
}

type opcode int
const (
	adv opcode = iota
	bxl opcode = iota
	bst opcode = iota
	jnz opcode = iota
	bxc opcode = iota
	out opcode = iota
	bdv opcode = iota
	cdv opcode = iota
)
func (c *computer) runProgram(program []int) {
	currentInstruction := 0
	for currentInstruction < len(program) {
		instruction := opcode(program[currentInstruction])
		operand := program[currentInstruction + 1]
		jumped := false
		switch instruction {
		case adv:
			c.registers[0] = c.divide(operand)
		case bxl:
			c.registers[1] = xor(c.registers[1], operand)
		case bst:
			c.registers[1] = c.modulo(operand)
		case jnz:
			if c.registers[0] != 0 {
				currentInstruction = operand
				jumped = true
			}
		case bxc:
			c.registers[1] = xor(c.registers[1], c.registers[2])
		case out:
			c.output = append(c.output, c.modulo(operand))
		case bdv:
			c.registers[1] = c.divide(operand)
		case cdv:
			c.registers[2] = c.divide(operand)
		}

		if !jumped {
			currentInstruction += 2
		}
	}
}

func (c *computer) divide(operand int) int {
	return c.registers[0] / maths.Pow(2, c.calculateComboOperand(operand))
}

func (c *computer) modulo(operand int) int {
	return c.calculateComboOperand(operand) % 8
}

func (c *computer) calculateComboOperand(operand int) int {
	if operand <= 3 {
		return operand
	}
	if operand <= 6 {
		return c.registers[operand - 4]
	}
	panic(fmt.Sprintf("Invalid operand: %d", operand))
}

func (c *computer) reset(a int) {
	c.output = []int{}
	c.registers = [3]int{a, 0, 0}
}

func xor(a, b int) int {
	return a ^ b
}
