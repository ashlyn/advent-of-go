package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(21, 2024, "\n")
	println(solvePart1(input))
	println()
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	robots := 3
	for _, line := range input {
		code := line
		for i := 0; i < robots; i++ {
			code = abstract("A" + code)
		}
		length, numeric := len(code), getNumericComponent(line)
		fmt.Printf("%s: %d * %d\n", line, length, numeric)
		result += length * numeric
	}

	return result
}

func solvePart2(input []string) int {
	result := 0
	robots := 25

	for _, line := range input {
		cache := make(map[cachedSequence]int)
		firstDirectional := abstract("A" + line)
		length := shortestSequence(firstDirectional, robots, cache)
		numeric := getNumericComponent(line)
		fmt.Printf("%s: %d * %d\n", line, length, numeric)
		result += length * numeric
	}
	return result
}

func getNumericComponent(input string) int {
	value, _ := strconv.Atoi(strings.ReplaceAll(input, "A", ""))
	return value
}

// All the optimized moves, pre-calculated by hand
// Prefers left to vertical to right
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
	sb := strings.Builder{}
	for i := 0; i < len(code) - 1; i++ {
		sb.WriteString(moveMap[code[i:i+2]])
		sb.WriteByte('A')
	}
	return sb.String()
}

type cachedSequence struct {
	sequence string
	depth int
}
func shortestSequence(moves string, depth int, cache map[cachedSequence]int) int {
	if depth == 0 {
		cache[cachedSequence{moves, depth}] = len(moves)
		return len(moves)
	}

	if cached, ok := cache[cachedSequence{moves, depth}]; ok {
		return cached
	}

	// length starts at -1 because the function will add 1 at each layer due to the implicit "A" start
	length := -1
	presses := strings.Split(moves, "A")
	for i := 0; i < len(presses); i++ {
		move := presses[i] + "A"
		nextMove := abstract("A" + move)
		length += shortestSequence(nextMove, depth - 1, cache)
	}
	cache[cachedSequence{moves, depth}] = length
	return length
}
