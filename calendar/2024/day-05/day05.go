package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(5, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	rules, updates := parseInput(input)

	for _, update := range updates {
		result += validateUpdate(update, rules)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	rules, updates := parseInput(input)
	rulesMap := rulesToMap(rules)
	for _, update := range updates {
		if validateUpdate(update, rules) == 0 {
			u := reorderUpdate(update, rulesMap)
			result += u[len(u) / 2]
		}
	}

	return result
}

func parseInput(input []string) ([][2]int, [][]int) {
	rules := [][2]int{}
	updates := [][]int{}

	i := 0;
	for ; i < len(input); i++ {
		if input[i] == "" {
			i++
			break
		}

		r := strings.Split(input[i], "|")
		x, _ := strconv.Atoi(r[0])
		y, _ := strconv.Atoi(r[1])
		rules = append(rules, [2]int{x, y})
	}

	for ; i < len(input); i++ {
		pages := slices.ParseIntsFromStrings(strings.Split(input[i], ","))
		updates = append(updates, pages)
	}

	return rules, updates
}

func validateUpdate(update []int, rules [][2]int) int {
	u := updateToMap(update)
	for _, rule := range rules {
		indexX, hasX := u[rule[0]]
		indexY, hasY := u[rule[1]]

		if hasX && hasY && indexX >= indexY {
			return 0
		}
	}
	return update[len(update) / 2]
}

func updateToMap(update []int) map[int]int {
	m := map[int]int{}
	for i, v := range update {
		m[v] = i
	}
	return m
}

func reorderUpdate(update []int, rulesMap map[int][]int) []int {
	sort.Slice(update, func(i, j int) bool {
		a, b := update[i], update[j]
		unordered := slices.Contains(rulesMap[b], a)
		return !unordered
	})

	return update
}

func rulesToMap(rules [][2]int) map[int][]int {
	m := map[int][]int{}
	for _, rule := range rules {
		_, hasX := m[rule[0]]
		if !hasX {
			m[rule[0]] = []int{}
		}
		m[rule[0]] = append(m[rule[0]], rule[1])
	}
	return m
}
