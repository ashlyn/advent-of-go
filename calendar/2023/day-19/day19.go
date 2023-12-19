package main

import (
	"advent-of-go/utils/files"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(19, 2023, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	rules, parts := parseInput(input)
	for _, part := range parts {
		if isPartAccepted(part, "in", rules) {
			result += calulateTotalPartRating(part)
		}
	}

	return result
}

func solvePart2(input []string) int {
	rules, _ := parseInput(input)
	return generateAcceptanceCriteria("in", 0, rules, map[string][2]int{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}})
}

func isPartAccepted(part map[string]int, currentWorkflow string, rules map[string][]string) bool {
	gt, lt := ">", "<"
	for i := 0; i < len(rules[currentWorkflow]); i++ {
		currentRule := rules[currentWorkflow]
		currentToken := currentRule[i]
		if currentToken == "A" {
			return true
		}
		if currentToken == "R" {
			return false
		}
		if strings.Contains(currentToken, gt) || strings.Contains(currentToken, lt) {
			label, operator := currentToken[0:1], currentToken[1:2]
			value, _ := strconv.Atoi(currentToken[2:])
			if !((operator	== lt  && part[label] < value) || (operator == gt && part[label] > value)) {
				i++
			}
		} else {
			return isPartAccepted(part, currentToken, rules)
		}
	}
	return false
}

func calculateValidCombinations(criteria map[string][2]int) int {
	result := 1
	for _, c := range criteria {
		result *= c[1] - c[0] + 1
	}
	return result
}

func generateAcceptanceCriteria(currentWorkflow string, currentTokenIndex int, rules map[string][]string, originalCriteria map[string][2]int) int {
	gt, lt := ">", "<"

	if currentWorkflow == "A" {
		return calculateValidCombinations(originalCriteria)
	}
	if currentWorkflow == "R" {
		return 0
	}
	currentRule := rules[currentWorkflow]
	currentToken := currentRule[currentTokenIndex]
	rejectedRange, acceptedRange := [2]int{}, [2]int{}
	if strings.Contains(currentToken, gt) || strings.Contains(currentToken, lt) {
		label, operator := currentToken[0:1], currentToken[1:2]
		value, _ := strconv.Atoi(currentToken[2:])
		currentRange := originalCriteria[label]
		if (operator == lt && currentRange[0] >= value) || (operator == gt && currentRange[1] <= value) {
			return 0
		} else if (operator == lt && currentRange[1] < value) || (operator == gt && currentRange[0] > value) {
			return generateAcceptanceCriteria(currentWorkflow, currentTokenIndex + 1, rules, originalCriteria)
		} else if operator == lt {
			rejectedRange = [2]int{ value, currentRange[1] }
			acceptedRange = [2]int{ currentRange[0], value - 1 }
		} else if operator == gt {
			rejectedRange = [2]int{ currentRange[0], value }
			acceptedRange = [2]int{ value + 1, currentRange[1] }
		}

		rejectedCriteria, acceptedCriteria := map[string][2]int{}, map[string][2]int{}
		for k, v := range originalCriteria {
			if k == label {
				rejectedCriteria[k] = rejectedRange
				acceptedCriteria[k] = acceptedRange
			} else {
				rejectedCriteria[k] = v
				acceptedCriteria[k] = v
			}
		}
		return generateAcceptanceCriteria(currentRule[currentTokenIndex + 1], 0, rules, acceptedCriteria) + generateAcceptanceCriteria(currentWorkflow, currentTokenIndex + 2, rules, rejectedCriteria)
	}
	
	return generateAcceptanceCriteria(currentToken, 0, rules, originalCriteria)
}

func calulateTotalPartRating(part map[string]int) int {
	total := 0
	for _, value := range part {
		total += value
	}
	return total
}

func parseInput(input []string) (map[string][]string, []map[string]int) {
	ruleInputs, partInputs := strings.Split(input[0], "\n"), strings.Split(input[1], "\n")
	rules := map[string][]string{}
	parts := make([]map[string]int, len(partInputs))

	for _, ruleInput := range ruleInputs {
		label, ruleParts := parseRule(ruleInput)
		rules[label] = ruleParts
	}

	for i, partInput := range partInputs {
		parts[i] = parsePart(partInput)
	}

	return rules, parts
}

func parseRule(input string) (string, []string) {
	characterPattern := regexp.MustCompile(`[:{},]`)
	parts := characterPattern.Split(input, -1)
	return parts[0], parts[1:len(parts)-1]
}

func parsePart(input string) map[string]int {
	part := map[string]int{}
	labelPattern := regexp.MustCompile(`([a-z]+)=([0-9]+)`)
	matches := labelPattern.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		value, _ := strconv.Atoi(match[2])
		part[match[1]] = value
	}
	return part
}
