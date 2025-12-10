package str

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func CharAt(str string, pos int) (string, error) {
	if pos >= len(str) {
		return "", errors.New("invalid index")
	}

	return string(str[pos]), nil
}

func ReplaceCharAt(str string, replacement string, index int) string {
	ret := ""

	for i := 0; i < len(str); i++ {
		if i != index {
			ret += string(str[i])
		} else {
			ret += replacement
		}
	}

	return ret
}

func Reverse(str string) string {
	ret := ""

	for i := len(str) - 1; i >= 0; i-- {
		ret += string(str[i])
	}

	return ret
}

// IndexesAny returns the indexes of any of the characters in chars found in str
func IndexesAny(str string, chars string) []int {
	indexes := []int{}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(chars); j++ {
			if str[i] == chars[j] {
				indexes = append(indexes, i)
				break
			}
		}
	}

	return indexes
}

// ParseAllGroupsBetween returns all substrings found between left and right delimiters
func ParseAllGroupsBetween(left, right string, line string) []string {
	pattern := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
	matches := pattern.FindAllStringSubmatch(line, -1)
	captured := make([]string, len(matches))
	for i, match := range matches {
		captured[i] = match[1]
	}
	return captured
}

// ParseDelimitedStringToInts parses a string delimited by the specified delimiter into a slice of ints
func ParseDelimitedStringToInts(input string, delimiter string) []int {
	parts := strings.Split(input, delimiter)
	ints := make([]int, len(parts))
	for i, part := range parts {
		ints[i], _ = strconv.Atoi(part)
	}
	return ints
}
