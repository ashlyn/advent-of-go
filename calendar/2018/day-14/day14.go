package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(14, 2018, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	target, _ := strconv.Atoi(input[0])
	rounds := 10
	recipes := []int{3, 7}
	elf1, elf2 := 0, 1
	for len(recipes) < target+rounds {
		elf1, elf2, recipes = generateRecipes(elf1, elf2, recipes)
	}
	return recipesToString(recipes[target : target+rounds])
}

func solvePart2(input []string) int {
	target := input[0]
	recipes := []int{3, 7}
	elf1, elf2 := 0, 1
	i := 0
	for i < 100000000 {
		elf1, elf2, recipes = generateRecipes(elf1, elf2, recipes)
		targetIndex := strings.Index(recipesToString(recipes), target)
		if targetIndex != -1 {
			return targetIndex
		}
		if (i % 10000) == 0 {
			println(i, len(recipes))
		}
		i++
	}
	return -1
}

func generateRecipes(elf1, elf2 int, recipes []int) (int, int, []int) {
	digits := strconv.Itoa(recipes[elf1] + recipes[elf2])
	for i := 0; i < len(digits); i++ {
		r, _ := strconv.Atoi(digits[i:i+1])
		recipes = append(recipes, r)
	}
	elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
	elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	return elf1, elf2, recipes
}

func recipesToString(recipes []int) string {
	sb := strings.Builder{}
	for _, r := range recipes {
		sb.WriteString(strconv.Itoa(r))
	}
	return sb.String()
}
