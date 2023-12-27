package main

import (
	"advent-of-go/utils/files"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bid int
}

type handSet []hand

func main() {
	input := files.ReadFile(7, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	hands := parseHands(input)
	return getWinnings(hands, false)
}

func solvePart2(input []string) int {
	hands := parseHands(input)
	return getWinnings(hands, true)
}

func parseHands(input []string) []hand {
	hands := make([]hand, len(input))
	for i, line := range input {
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		hands[i] = hand{ cards: parts[0], bid: bid }
	}
	return hands
}

func configuredLess(h handSet, jokersWild bool) func(i, j int) bool {
	cardMap := map[rune]int{
		'2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9,
		'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
	}
	if jokersWild {
		cardMap['J'] = 0
	}

	return func(i, j int) bool {
		cardI, cardJ := h[i], h[j]
		iScore, jScore := cardI.score(jokersWild), cardJ.score(jokersWild)

		if iScore != jScore {
			return iScore < jScore
		}

		for tieCard := 0; tieCard < len(cardI.cards); tieCard++ {
			currentIValue, currentJValue := cardMap[rune(cardI.cards[tieCard])], cardMap[rune(cardJ.cards[tieCard])]
			if currentIValue != currentJValue {
				return currentIValue < currentJValue
			}
		}

		return false
	}
}

func (h hand) score(jokersWild bool) int {
	cardSet := map[rune]int{}
	for _, c := range h.cards {
		cardSet[c]++
	}

	jokerCount, hasJokers := cardSet['J']
	if jokersWild && hasJokers {
		if jokerCount == 5 {
			return 7
		}
		delete(cardSet, 'J')
		mostCommon, mostCommonCount := '0', 0
		for c, count := range cardSet {
			if count > mostCommonCount {
				mostCommon, mostCommonCount = c, count
			}
		}
		cardSet[mostCommon] += jokerCount
	}

	switch len(cardSet) {
	case 5:
		// high card
		return 1
	case 4:
			// one pair
			return 2
	case 3:
		for _, count := range cardSet {
			// three of a kind
			if count == 3 {
				return 4
			}
		}
		// two pair
		return 3
	case 2:
		for _, count := range cardSet {
			if count == 4 {
				// four of a kind
				return 6
			}
		}
		// full house
		return 5
	case 1:
		// five of a kind
		return 7
	default:
		return -1
	}
}

func getWinnings(hands []hand, jokersWild bool) int {
	winnings := 0
	sort.Slice(handSet(hands), configuredLess(handSet(hands), jokersWild))
	for i := 0; i < len(hands); i++ {
		winnings += (i + 1) * hands[i].bid
	}
	return winnings
}
