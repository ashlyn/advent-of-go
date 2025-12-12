package main

import (
	"advent-of-go/utils/colors"
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(12, 2025, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	presents, zones := parseInput(input)

	willFit, wontFit, toCheck := 0, 0, 0
	for _, z := range zones {
		fit := z.IsValid(presents)
		if fit == yes {
			willFit++
		} else if fit == no {
			wontFit++
		} else {
			toCheck++
		}
	}

	if toCheck == 0 {
		return willFit
	}

	// The sample input requires tetris fitting the zones
	// The real input does not, so no generalized solution is implemented
	fmt.Println(colors.RedString("Part 1 requires manual checking of zones: "), toCheck)
	return willFit
}

func solvePart2(input []string) string {
	return colors.GreenString("Merry Christmas!")
}

type present []string
type zone struct {
	width, height int
	presentRequirements []int
}

func parseInput(input []string) ([]present, []zone) {
	presents := make([]present, len(input) - 1)

	for i := 0; i < len(input) - 1; i++ {
		lines := strings.Split(input[i], "\n")
		presents[i] = lines[1:]
	}

	zoneStrings := strings.Split(input[len(input) - 1], "\n")
	zones := make([]zone, len(zoneStrings))
	for i, line := range zoneStrings {
		zones[i] = parseZone(line)
	}

	return presents, zones
}

func parseZone(line string) zone {
	parts := strings.Fields(line)
	dimensionsPattern := regexp.MustCompile(`(\d+)x(\d+):`)
	matches := dimensionsPattern.FindStringSubmatch(parts[0])
	width, _ := strconv.Atoi(matches[1])
	height, _ := strconv.Atoi(matches[2])
	return zone{
		width: width,
		height: height,
		presentRequirements: slices.ParseIntsFromStrings(parts[1:]),
	}
}

var yes, no, unknown rune = 'Y', 'N', '?'
func (z zone) IsValid(presents []present) rune {
	if !z.HasEnoughVolume(presents) {
		return no
	}
	if z.CanFitUnstacked(presents) {
		return yes
	}
	return unknown
}

func (z zone) HasEnoughVolume(presents []present) bool {
	volume := z.width * z.height
	requiredVolume := 0
	for i := 0; i < len(z.presentRequirements); i++ {
		if z.presentRequirements[i] > 0 {
			requiredVolume += presents[i].Volume() * z.presentRequirements[i]
		}
	}
	return requiredVolume <= volume
}

func (z zone) CanFitUnstacked(presents []present) bool {
	dividedWidth, dividedHeight := z.width / 3, z.height / 3
	boxesAvailable := dividedWidth * dividedHeight
	requiredBoxes := slices.Sum(z.presentRequirements)
	return requiredBoxes <= boxesAvailable
}

func (p present) Volume() int {
	volume := 0
	for _, line := range p {
		volume += strings.Count(line, "#")
	}
	return volume
}