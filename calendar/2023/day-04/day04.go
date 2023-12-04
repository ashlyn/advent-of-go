package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"strconv"
	"strings"
)

type card struct {
	number int
	numbers *sets.Set
	winningNumbers *sets.Set
	matches *sets.Set
}

func main() {
	input := files.ReadFile(4, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for _, line := range input {
		m := parseCard(line).matches
		result += maths.Pow(2, m.Size() - 1)
	}

	return result
}

func solvePart2(input []string) int {
	originalCards := []*card{}
	cardList := []int{}
	for _, card := range input {
		c := parseCard(card)
		originalCards = append(originalCards, parseCard(card))
		cardList = append(cardList, c.number)
	}

	for i := 0; i < len(cardList); i++ {
		c := originalCards[cardList[i] - 1]
		for j := 0; j < c.matches.Size(); j++ {
			cardList = append(cardList, originalCards[c.number + j].number)
		}
	}

	return len(cardList)
}

func parseCard(input string) *card {
	parts := strings.Split(input, " | ")
	cardParts := strings.Fields(parts[0])
	number, _ := strconv.Atoi(strings.Replace(cardParts[1], ":", "", -1))

	numbers, winningNumbers := sets.New(), sets.New()
	numbers.AddRange(cardParts[2:])
	winningNumbers.AddRange(strings.Fields(parts[1]))

	matches := numbers.Intersect(winningNumbers)
	return &card{ number: number, numbers: &numbers, winningNumbers: &winningNumbers, matches: &matches }
}
