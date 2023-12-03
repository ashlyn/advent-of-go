package main

import (
	"advent-of-go/utils/files"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

type set struct {
	red int
	green int
	blue int
}

type game struct {
	id int
	sets []set
}

func solvePart1(input []string) int {
	result := 0

	games := parseGames(input)
	bag := set{red: 12, green: 13, blue: 14}
	for _, g := range games {
		if gameIsPossible(g, bag) {
			result += g.id
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	games := parseGames(input)
	for _, g := range games {
		s := findMinimumBag(g)
		result += calculatePower(s)
	}

	return result
}

func parseGames(input []string) []game {
	games := make([]game, len(input))
	for i, line := range input {
		games[i] = parseGame(line)
	}
	return games
}

func parseGame(line string) game {
	gameParts := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.Replace(gameParts[0], "Game ", "", 1))
	return game{ id: id, sets: parseSets(strings.Split(gameParts[1], "; ")) }
}

func parseSets(setStrings []string) []set {
	sets := make([]set, len(setStrings))
	for i, s := range setStrings {
		sets[i] = parseSet(s)
	}
	return sets
}

func parseSet(setString string) set {
	redPattern := regexp.MustCompile(`([0-9]+) red`)
	greenPattern := regexp.MustCompile(`([0-9]+) green`)
	bluePattern := regexp.MustCompile(`([0-9]+) blue`)

	red, green, blue := parseCount(setString, redPattern), parseCount(setString, greenPattern), parseCount(setString, bluePattern)

	return set{ red: red, green: green, blue: blue }
}

func parseCount(setString string, pattern *regexp.Regexp) int {
	count, _ := strconv.Atoi(strings.Split(pattern.FindString(setString), " ")[0])
	return count
}

func setIsPossible(set set, bag set) bool {
	return set.red <= bag.red && set.green <= bag.green && set.blue <= bag.blue
}

func gameIsPossible(game game, bag set) bool {
	for _, set := range game.sets {
		if !setIsPossible(set, bag) {
			return false
		}
	}
	return true
}

func findMinimumBag(game game) set {
	red, green, blue := 0, 0, 0
	for _, set := range game.sets {
		if set.red > red {
			red = set.red
		}
		if set.green > green {
			green = set.green
		}
		if set.blue > blue {
			blue = set.blue
		}
	}
	return set{ red: red, green: green, blue: blue }
}

func calculatePower(set set) int {
	return set.red * set.green * set.blue
}
