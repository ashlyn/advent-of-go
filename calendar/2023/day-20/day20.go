package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"sort"
	"strings"
)

type pulseType int
const (
	none pulseType = iota - 1
	low pulseType = iota - 1
	high pulseType = iota - 1
)

type moduleType string
const (
	broadcast = "broadcaster"
	flipflop = "flipflop"
	conjunction = "conjunction"
	untyped = "untyped"
)

func main() {
	input := files.ReadFile(20, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	low, high := 0, 0
	modules := parseInput(input)
	for i := 1; i <= 1000; i++ {
		l, h, _ := pressButton(modules, "")
		low += l
		high += h
	}
	return low * high
}

func solvePart2(input []string) int {
	result := 1
	// "reverse engineered" the conjunction modules feeding into rx
	criticalModules := []string{ "kb", "ks", "sx", "jt" }
	modules := parseInput(input)
	i := 0
	for _, m := range criticalModules {
		for isOn := false; !isOn && i < 99999999; i++ {
			_, _, isOn = pressButton(modules, m)
		}
		factors := maths.PrimeFactorsSlice(i)
		sort.Ints(factors)
		result = maths.Lcm(result, factors[len(factors)-1])
	}
	return result
}

type module struct {
	name string
	destinations []string
	typeName moduleType
	flipFlopIsOn bool
	conjunctionState map[string]pulseType
}

const (
	separator = " -> "
	flipflopPrefix = '%'
	conjunctionPrefix = '&'
)
func parseInput(input []string) map[string]*module {
	modules := map[string]*module{}
	for _, line := range input {
		parts := strings.Split(line, separator)
		name, destinations := parts[0], strings.Split(parts[1], ", ")
		var mt moduleType = untyped
		if name[0]	== flipflopPrefix {
			name = name[1:]
			mt = flipflop
		} else if name[0] == conjunctionPrefix {
			name = name[1:]
			mt = conjunction
		} else if name == broadcast {
			mt = broadcast
		}
		m := module{ name: name, destinations: destinations, typeName: mt, conjunctionState: map[string]pulseType{} }
		modules[name] = &m
	}

	for name, m := range modules {
		for _, d := range m.destinations {
			dest, destExists := modules[d]
			if destExists && dest.typeName == conjunction {
				dest.conjunctionState[name] = low
			}
		}
	}
	return modules
}

func receivePulse(pulse pulseType, toModule string, fromModule string, modules map[string]*module) pulseType {
	m, moduleExists := modules[toModule]
	if !moduleExists {
		return none
	}
	if m.typeName == broadcast {
		return pulse
	} else if m.typeName == flipflop {
		if pulse == high {
			return none
		}
		m.flipFlopIsOn = !m.flipFlopIsOn
			if m.flipFlopIsOn {
				return high
			}
			return low
	} else if m.typeName == conjunction {
		m.conjunctionState[fromModule] = pulse
		for _, state := range m.conjunctionState {
			if state == low {
				return high
			}
		}
		return low
	}
	return none
}

type instruction struct {
	fromModule string
	toModule string
	pulse pulseType
}
func pressButton(modules map[string]*module, moduleToCheck string) (int, int, bool) {
	isOn := false
	pulseCounts := map[pulseType]int{ low: 1 }
	queue := []instruction{}
	for _, d := range modules[broadcast].destinations {
		queue = append(queue, instruction{ fromModule: broadcast, toModule: d, pulse: pulseType(low) })
	}
	for len(queue) > 0 {
		currentInstruction := queue[0]
		queue = queue[1:]
		pulseCounts[currentInstruction.pulse]++
		toSendNext := receivePulse(currentInstruction.pulse, currentInstruction.toModule, currentInstruction.fromModule, modules)

		if toSendNext == none {
			continue
		}
		for _, d := range modules[currentInstruction.toModule].destinations {
			if moduleToCheck != "" && d == moduleToCheck && toSendNext == low {
				isOn = true
			}
			queue = append(queue, instruction{ fromModule: currentInstruction.toModule, toModule: d, pulse: toSendNext })
		}
	}
	return pulseCounts[low], pulseCounts[high], isOn
}
