package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(11, 2024, "\n")
	sw := time.Now()
	println(solvePart1(input))
	fmt.Printf("Solved part 1 in %v\n", time.Since(sw))
	sw = time.Now()
	println(solvePart2(input))
	fmt.Printf("Solved part 2 in %v\n", time.Since(sw))
}

func solvePart1(input []string) int {
	stonesCounts := parseInput(input[0])
	return applyRulesToStonesNTimes(stonesCounts, 25)
}

func solvePart2(input []string) int {
	stonesCounts := parseInput(input[0])
	return applyRulesToStonesNTimes(stonesCounts, 75)
}

func parseInput(input string) map[string]int {
	stonesArr := strings.Fields(input)
	stones := map[string]int{}
	for i := 0; i < len(stonesArr); i++ {
		stones[stonesArr[i]]++
	}
	return stones
}

func applyRules(stones map[string]int) map[string]int {
	newStones := map[string]int{}
	for stone, count := range stones {
		if stone == "0" {
			newStones["1"] += count
		} else if len(stone) % 2 == 0 {
			firstHalf := stone[:len(stone) / 2]
			secondValue, _ := strconv.Atoi(stone[len(stone) / 2:])
			secondHalf := fmt.Sprintf("%d", secondValue)
			newStones[firstHalf] += count
			newStones[secondHalf] += count
		} else {
			newValue, _ := strconv.Atoi(stone)
			newStones[fmt.Sprintf("%d", newValue * 2024)] += count
		}
	}
	return newStones
}

func applyRulesToStonesNTimes(stonesCounts map[string]int, n int) int {
	for i := 0; i < n; i++ {
		stonesCounts = applyRules(stonesCounts)
	}

	result := 0
	for _, count := range stonesCounts {
		result += count
	}
	return result
}
