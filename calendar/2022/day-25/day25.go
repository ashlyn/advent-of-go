package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"fmt"
)

func main() {
	input := files.ReadFile(25, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {

	sum := 0
	for _, str := range input {
		sum += snafuToDecimal(str)
	}

	return decimalToSnafu(sum)
}

func solvePart2(input []string) string {
	return "Merry Christmas!"
}

func snafuToDecimal(snafu string) int {
	decimal := 0
	for i, char := range snafu {
		var digit int
		if char == '-' {
			digit = -1
		} else if char == '=' {
			digit = -2
		} else {
			digit = int(char) - 48
		}
		decimal += digit * maths.Pow(5, len(snafu) - (i + 1))
	}
	return decimal
}

func decimalToSnafu(decimal int) string {
	snafu := ""

	if decimal == 0 {
		return "0"
	}
	for decimal != 0 {
		rem := decimal % 5
		decimal /= 5
		if rem > 2 {
			decimal++
			rem -= 5
		}
		if rem == -2 {
			snafu = "=" + snafu
		} else if rem == -1 {
			snafu = "-" + snafu
		} else {
			snafu = fmt.Sprint(rem) + snafu
		}
	}
	return snafu
}
