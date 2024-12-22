package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
)

func main() {
	input := files.ReadFile(22, 2024, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	initialNumbers := slices.ParseIntsFromStrings(input)
	for i := 0; i < len(initialNumbers); i++ {
		initial := initialNumbers[i]
		next := initial
		for times := 0; times < 2000; times++ {
			next = getNextSecretNumber(next)
		}
		result += next
	}

	return result
}

func solvePart2(input []string) int {
	initialNumbers := slices.ParseIntsFromStrings(input)
	sequences := map[[4]int][]int{}
	secrets := make([][]int, len(initialNumbers))
	maxSecrets := 2000
	for i := 0; i < len(initialNumbers); i++ {
		initial := initialNumbers[i]
		secrets[i] = make([]int, maxSecrets + 1)
		secrets[i][0] = initial
		for times := 0; times < maxSecrets; times++ {
			secrets[i][times + 1] = getNextSecretNumber(secrets[i][times])
		}
	}

	// quite possibly is not checking the very first or last sequences possible
	for i := 0; i < len(secrets); i++ {
		secretsChain := secrets[i]
		for j := 5; j < len(secretsChain); j++ {
			sequence := [4]int{}
			values := secretsChain[j - 5:j]
			for k := 0; k < 4; k++ {
				previous, current := values[k], values[k + 1]
				delta := (current % 10) - (previous % 10)
				sequence[k] = delta
			}
			price := secretsChain[j-1] % 10
			if _, ok := sequences[sequence]; !ok {
				sequences[sequence] = make([]int, len(secrets))
			}
			if sequences[sequence][i] == 0 {
				sequences[sequence][i] = price
			}
		}
	}

	maxBananas, bestSequence := 0, [4]int{}
	for sequence, prices := range sequences {
		sum := slices.Sum(prices)
		if sum > maxBananas {
			maxBananas = sum
			bestSequence = sequence
		}
	}

	fmt.Println(bestSequence)
	return maxBananas
}

func mix(value, secretNumber int) int {
	return value ^ secretNumber
}

func prune(secretNumber int) int {
	return secretNumber % 16777216
}

func stepOne(secretNumber int) int {
	return prune(mix(secretNumber, secretNumber * 64))
}

func stepTwo(secretNumber int) int {
	return prune(mix(secretNumber, secretNumber / 32))
}

func stepThree(secretNumber int) int {
	return prune(mix(secretNumber, secretNumber * 2048))
}

func getNextSecretNumber(secretNumber int) int {
	return stepThree(stepTwo(stepOne(secretNumber)))
}
