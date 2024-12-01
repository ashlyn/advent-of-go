package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(1, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	list1, list2 := parseLists(input)

	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	for i := 0; i < len(list1); i++ {
		result += maths.Abs(list1[i] - list2[i])
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	list1, list2 := parseLists(input)

	frequencyMap := make(map[int]int)
	for i := 0; i < len(list1); i++ {
		frequencyMap[list2[i]]++
	}

	for i := 0; i < len(list1); i++ {
		value := list1[i]
		result += value * frequencyMap[value]
	}

	return result
}

func parseLists(input []string) ([]int, []int) {
	list1, list2 := make([]int, len(input)), make([]int, len(input))
	for i, line := range input {
		values := strings.Fields(line)
		list1[i], _ = strconv.Atoi(values[0])
		list2[i], _ = strconv.Atoi(values[1])
	}
	return list1, list2
}
