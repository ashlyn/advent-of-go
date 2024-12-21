package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
)

func main() {
	input := files.ReadFile(21, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	robots := 3
	for _, line := range input {
		code := line
		println(line)
		for i := 0; i < robots; i++ {
			code = abstract("A" + code)
			println(code)
		}
		fmt.Printf("%d * %d\n", len(code), getNumericComponent(line))
		println()
		result += getNumericComponent(line) * len(code)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func getNumericComponent(input string) int {
	value, _ := strconv.Atoi(input[:len(input)-1])
	return value
}

var moveMap = map[string]string {
	"A0": "<",
	"A1": "^<<",
	"A2": "<^",
	"A3": "^",
	"A4": "^^<<",
	"A5": "<^^",
	"A6": "^^",
	"A7": "^^^<<",
	"A8": "<^^^",
	"A9": "^^^",

	"0A": ">",
	"01": "^<",
	"02": "^",
	"03": "^>",
	"04": "^^>",
	"05": "^^",
	"06": "^^>",
	"07": "^^^<",
	"08": "^^^",
	"09": "^^^>",

	"1A": ">>v",
	"10": ">v",
	"12": ">",
	"13": ">>",
	"14": "^",
	"15": "^>",
	"16": "^>>",
	"17": "^^",
	"18": "^^>",
	"19": "^^>>",

	"2A": "v>",
	"20": "v",
	"21": "<",
	"23": ">",
	"24": "<^",
	"25": "^",
	"26": "^>",
	"27": "<^^",
	"28": "^^",
	"29": "^^>",

	"3A": "v",
	"30": "<v",
	"31": "<<",
	"32": "<",
	"34": "<<^",
	"35": "<^",
	"36": "^",
	"37": "<<^^",
	"38": "<^^",
	"39": "^^",

	"4A": ">>vv",
	"40": ">vv",
	"41": "v",
	"42": "v>",
	"43": "v>>",
	"45": ">",
	"46": ">>",
	"47": "^",
	"48": "^>",
	"49": "^>>",

	"5A": "vv>",
	"50": "vv",
	"51": "<v",
	"52": "v",
	"53": "v>",
	"54": "<",
	"56": ">",
	"57": "<^",
	"58": "^",
	"59": "^>",

	"6A": "vv",
	"60": "<vv",
	"61": "<<v",
	"62": "<v",
	"63": "v",
	"64": "<<",
	"65": "<",
	"67": "<<^",
	"68": "<^",
	"69": "^",

	"7A": ">>vvv",
	"70": ">vvv",
	"71": "vv",
	"72": "vv>",
	"73": "vv>>",
	"74": "v",
	"75": "v>",
	"76": "v>>",
	"78": ">",
	"79": ">>",

	"8A": "vvv>",
	"80": "vvv",
	"81": "<vv",
	"82": "vv",
	"83": "vv>",
	"84": "<v",
	"85": "v",
	"86": "v>",
	"87": "<",
	"89": ">",

	"9A": "vvv",
	"90": "<vvv",
	"91": "<<vv",
	"92": "<vv",
	"93": "vv",
	"94": "<<v",
	"95": "<v",
	"96": "v",
	"97": "<<",
	"98": "<",

	"A^": "<",
	"A>": "v",
	"Av": "<v",
	"A<": "v<<",

	"^A": ">",
	"^>": "v>",
	"^v": "v",
	"^<": "v<",

	">A": "^",
	">^": "<^",
	">v": "<",
	"><": "<<",

	"vA": "^>",
	"v>": ">",
	"v^": "^",
	"v<": "<",

	"<A": ">>^",
	"<^": ">^",
	"<v": ">",
	"<>": ">>",
}
func abstract(code string) string {
	result := ""

	for i := 0; i < len(code) - 1; i++ {
		if code[i] == code[i+1] {
			result += "A"
			continue
		}
		_, ok := moveMap[code[i:i+2]]
		if !ok {
			println("Not found ", code[i:i+2])
		}
		result += moveMap[code[i:i+2]] + "A"
	}

	return result
}
