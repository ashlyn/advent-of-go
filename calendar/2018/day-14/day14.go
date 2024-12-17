package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

type recipes struct {
	scores []byte
}

func (r *recipes) Generate(i, j int) (int, int) {
	d1, d2 := int(r.scores[i]-'0'), int(r.scores[j]-'0')
	r.add(d1 + d2)
	return (i + d1 + 1) % len(r.scores), (j + d2 + 1) % len(r.scores)
}

func (r *recipes) String() string {
	return string(r.scores)
}

func (r *recipes) Len() int {
	return len(r.scores)
}

func (r * recipes) Last(n int) []byte {
	return r.scores[len(r.scores)-n:]
}

func newRecipes(input []int) recipes {
	r := recipes{scores: []byte{}}
	for _, digit := range input {
		r.add(digit)
	}
	return r
}

func (r *recipes) add(score int) {
	s := []byte(strconv.Itoa(score))
	r.scores = append(r.scores, s...)
}

func main() {
	input := files.ReadFile(14, 2018, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	target, _ := strconv.Atoi(input[0])
	rounds := 10
	recipes := newRecipes([]int{3, 7})
	elf1, elf2 := 0, 1
	for recipes.Len() < target+rounds {
		elf1, elf2 = recipes.Generate(elf1, elf2)
	}
	return recipes.String()[target:target+rounds]
}

func solvePart2(input []string) int {
	target  := input[0]
	recipes := newRecipes([]int{3, 7})
	elf1, elf2 := 0, 1
	for i := 0; i < 40000000; i++ {
		elf1, elf2 = recipes.Generate(elf1, elf2)
		// wanted to have a stopping condition here and couldn't get it working as expected
		// instead, cutting off at a "reasonable" number of iterations
	}
	return strings.Index(recipes.String(), target)
}
