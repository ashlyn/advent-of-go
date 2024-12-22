package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"math"
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
			// println(code)
		}
		length, numeric := len(code), getNumericComponent(line)
		fmt.Printf("%s: %d * %d\n", line, length, numeric)
		result += length * numeric
	}

	return result
}

// 108790909750248 too low
//    897690107904 too low
// off by 13704 with r = 3 (6 per code)
// idek anymore
func solvePart2(input []string) int {
	result := 0
	robots := 24 // 2
	cache := make(map[string][]int)
	for _, line := range input {
		length := abstractMemoized(abstract("A" + line), robots, 0, cache)
		numeric := getNumericComponent(line)
		fmt.Printf("%s: %d * %d\n", line, length, numeric)
		if (result + (length - 1) * numeric) < result {
			panic("overflow")
		}
		result += (length - 1) * numeric
	}

	println(math.MaxInt)
	return result
}

func getNumericComponent(input string) int {
	value, _ := strconv.Atoi(strings.ReplaceAll(input, "A", ""))
	return value
}

// All the optimized moves
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

func translateMove(from string, to string) string {
	return moveMap[from + to] + "A"
}
func abstract(code string) string {
	sb := strings.Builder{}
	for i := 0; i < len(code) - 1; i++ {
		sb.WriteString(translateMove(code[i:i+1], code[i+1:i+2]))
		// sb.WriteString(moveMap[code[i:i+2]])
		// sb.WriteByte('A')
	}
	return sb.String()
}

func abstractMemoized(code string, totalRobots, currentRobot int, cache map[string][]int) int {
	if cached, ok := cache[code]; ok {
		if cached[currentRobot] != 0 {
			return cached[currentRobot]
		}
	} else {
		cache[code] = make([]int, totalRobots)
	}

	nextCode := abstract("A" + code)
	cache[code][0] = len(nextCode)

	if currentRobot == totalRobots - 1 {
		return len(nextCode)
	}

	length := 0

	steps := strings.Split(nextCode, "A")
	for i := 0; i < len(steps); i++ {
		s := steps[i] + "A"
		nextLength := abstractMemoized(s, totalRobots, currentRobot + 1, cache)
		if _, ok := cache[s]; !ok {
			cache[s] = make([]int, totalRobots)
		}
		cache[code][0] = nextLength
		length += nextLength
	}

	cache[code][currentRobot] = length

	return length
}
