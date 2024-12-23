package slices

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func IntsToStrings(slice []int) []string {
	strings := make([]string, len(slice))
	for i, num := range slice {
		strings[i] = strconv.Itoa(num)
	}
	return strings
}

func ParseIntsFromStrings(slice []string) []int {
	ints := make([]int, len(slice))
	for i, str := range slice {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func Unpack(slice []string, vars ...*string) {
	for i, str := range slice {
		*vars[i] = str
	}
}

func ParseLine(line string, splitOn string, vars ...*string) {
	regex := regexp.MustCompile(splitOn)
	Unpack(regex.Split(line, -1), vars...)
}

func Filter[T comparable](slice []T, filter func(T) bool) []T {
	retSlice := make([]T, 0)

	for _, element := range slice {
		if filter(element) {
			retSlice = append(retSlice, element)
		}
	}

	return retSlice
}

func Contains[T comparable](slice []T, word T) bool {
	for _, element := range slice {
		if element == word {
			return true
		}
	}

	return false
}

func Mode(input []int) int {
	counts := make(map[int]int)
	for _, value := range input {
		counts[value]++
	}
	maxCount, maxValue := 0, 0
	for value, count := range counts {
		if count > maxCount {
			maxCount = count
			maxValue = value
		}
	}
	return maxValue
}

func Max(input []int) int {
	max := math.MinInt
	for _, element := range input {
		if element > max {
			max = element
		}
	}

	return max
}

func Min(input []int) int {
	min := math.MaxInt
	for _, element := range input {
		if element < min {
			min = element
		}
	}

	return min
}

func Frame(slice []string) []string {
	framed := make([]string, len(slice)+1)
	padding := strings.Repeat(".", len(slice[0])+2)

	framed = append(framed, padding)
	framed[0] = padding

	for i := 1; i < len(slice)+1; i++ {
		framed[i] = fmt.Sprintf(".%s.", slice[i-1])
	}

	return framed
}

func Equals[T comparable](first []T, second []T) bool {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func CountCharInSlice(slice []string, char string) int {
	numChars := 0
	for _, element := range slice {
		numChars += strings.Count(element, char)
	}

	return numChars
}

func Sum(slice []int) int {
	sum := 0
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

func Swap(slice interface{}, i int, j int) {
	if reflect.TypeOf(slice).Kind() == reflect.Slice {
		reflect.Swapper(slice)(i, j)
	}
}

func Fill(value int, count int, arr *[]int) {
	for i := 0; i <= count; i++ {
		*arr = append(*arr, value)
	}
}

func GeneratePermutations(items []int) [][]int {
	length := len(items)

	initial, itemsCopy := make([]int, length), make([]int, length)
	copy(initial, items)
	copy(itemsCopy, items)

	permutations := [][]int{initial}

	indexes := make([]int, length)

	i := 0
	for i < length {
		if indexes[i] < i {
			if i%2 == 0 {
				Swap(itemsCopy, 0, i)
			} else {
				Swap(itemsCopy, indexes[i], i)
			}
			permutation := make([]int, length)
			copy(permutation, itemsCopy)
			permutations = append(permutations, permutation)
			indexes[i] = indexes[i] + 1
			i = 0
		} else {
			indexes[i] = 0
			i++
		}
	}

	return permutations
}

func GenerateCombinationsLengthNChannel(items []int, n int, abort <-chan []int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		length := len(items)
		itemsCopy := make([]int, length)
		copy(itemsCopy, items)

		select {
		case <-abort:
			return
		default:
			if length == 0 || n > length || n == 0 {
				c <- []int{}
				return
			} else if n == length {
				initial := make([]int, length)
				copy(initial, itemsCopy)
				c <- initial
				return
			}
	
			if n == length {
				for _, element := range itemsCopy {
					c <- []int{element}
				}
			}
	
			first := itemsCopy[0]
			for combo := range GenerateCombinationsLengthNChannel(itemsCopy[1:], n-1, abort) {
				c <- append([]int{first}, combo...)
			}
			for combo := range GenerateCombinationsLengthNChannel(itemsCopy[1:], n, abort) {
				c <- combo
			}
		}
	}()
	return c
}

func GenerateCombinationsLengthN[T comparable](items []T, n int) [][]T {
	length := len(items)
	itemsCopy := make([]T, length)
	copy(itemsCopy, items)

	if length == 0 || n > length || n == 0 {
		return [][]T{{}}
	} else if n == length {
		initial := make([]T, length)
		copy(initial, itemsCopy)
		return [][]T{initial}
	}

	if n == length {
		combinations := [][]T{}
		for _, element := range itemsCopy {
			combinations = append(combinations, []T{element})
			return combinations
		}
	}

	first := itemsCopy[0]
	nMinusOneCombinations := GenerateCombinationsLengthN(itemsCopy[1:], n-1)
	for i := range nMinusOneCombinations {
		nMinusOneCombinations[i] = append([]T{first}, nMinusOneCombinations[i]...)
	}
	return append(nMinusOneCombinations, GenerateCombinationsLengthN(itemsCopy[1:], n)...)
}

func GenerateCombinationsLengthNGeneric[T comparable](items []T, n int) [][]T {
	length := len(items)
	itemsCopy := make([]T, length)
	copy(itemsCopy, items)

	if length == 0 || n > length || n == 0 {
		return [][]T{{}}
	} else if n == length {
		initial := make([]T, length)
		copy(initial, itemsCopy)
		return [][]T{initial}
	}

	if n == length {
		combinations := [][]T{}
		for _, element := range itemsCopy {
			combinations = append(combinations, []T{element})
			return combinations
		}
	}

	first := itemsCopy[0]
	nMinusOneCombinations := GenerateCombinationsLengthNGeneric(itemsCopy[1:], n-1)
	for i := range nMinusOneCombinations {
		nMinusOneCombinations[i] = append([]T{first}, nMinusOneCombinations[i]...)
	}
	return append(nMinusOneCombinations, GenerateCombinationsLengthNGeneric(itemsCopy[1:], n)...)
}

func GenerateAllCombinations(items []int) [][]int {
	combinations := [][]int{}
	for n := 0; n <= len(items); n++ {
		combinations = append(combinations, GenerateCombinationsLengthN(items, n)...)
	}
	return combinations
}

// IndexOf returns the index of the selected item or -1 if not present
func IndexOf[T comparable](item T, slice []T) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// IndexOfStr returns the index of the selected item or -1 if not present
func IndexOfStr(item string, slice []string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// IndexOfInt returns the index of the selected item or -1 if not present
func IndexOfInt(item int, slice []int) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Reverse returns a slice in the reversed order of the input
func Reverse[T comparable](input []T) []T {
	reversed := make([]T, len(input))

	for i, value := range input {
		reversed[len(reversed) - i - 1] = value
	}

	return reversed
}

// TrimRight removes all trailing elements from the slice that are equal to the trimValue
func TrimRight[T comparable](input []T, trimValue T) []T {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != trimValue {
			return input[:i+1]
		}
	}
	return input
}

// IndexOfSubset returns the index of the subset in the slice or -1 if not present
func IndexOfSubset[T comparable](slice []T, subset []T) int {
	if len(slice) < len(subset) {
		return -1
	}
	for i := 0; i < len(slice)-len(subset)+1; i++ {
		if Equals(slice[i:i+len(subset)], subset) {
			return i
		}
	}
	return -1
}
