package main

import (
	"advent-of-go/utils/conv"
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(22, 2023, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	bricks, supportedBy := stackBlocks(input)
	return getSafeBricks(bricks, supportedBy).Size()
}

func solvePart2(input []string) int {
	bricks, supportedBy := stackBlocks(input)
	return calculateChainReactions(bricks, supportedBy)
}

func getZValue(input string) int {
	z, _ := strconv.Atoi(strings.Split(input, ",")[4])
	return z
}

type blockSlice struct {
	z int
	block string
}
func stackBlocks(input []string) (*sets.Set, map[string]*sets.Set) {
	stacks, bricks := map[grid.Coords][]blockSlice{}, sets.New()
	sort.Slice(input, func(i, j int) bool {
		iZ, jZ := getZValue(input[i]), getZValue(input[j])
		return iZ < jZ
	})
	for i, line := range input {
		b := fmt.Sprintf("%d", i) // string(rune(i + 65))
		bricks.Add(b)
		values := conv.ToIntSlice(strings.Split(strings.Replace(line, "~", ",", -1),","))
		coords, zValue := []grid.Coords{}, 0
		for x := values[0]; x <= values[3]; x++ {
			for y := values[1]; y <= values[4]; y++ {
				c := grid.Coords{ X: x, Y: y }
				coords = append(coords, c)
				if stack, ok := stacks[c]; !ok {
					stacks[c] = []blockSlice{}
				} else if len(stack) != 0 {
					zValue = maths.Max(stack[len(stack) - 1].z + 1, zValue)
				}
			}
		}
		for _, c := range coords {
			height := values[5] - values[2] + 1
			for z := 0; z < height; z++ {
				stacks[c] = append(stacks[c], blockSlice{zValue + z, b })
			}
		}
	}
	supportedBy := map[string]*sets.Set{}
	for _, v := range stacks {
		for i := len(v) - 1; i > 0; i-- {
			current, under := v[i], v[i - 1]
			if current.z == under.z + 1 && current.block != under.block {
				if _, ok := supportedBy[current.block]; !ok {
					s := sets.New()
					supportedBy[current.block] = &s
				}
				supportedBy[current.block].Add(under.block)
			}
		}
	}
	return &bricks, supportedBy
}

func getSafeBricks(bricks *sets.Set, supportedBy map[string]*sets.Set) *sets.Set {
	safeToDisintigrateBricks := bricks.Copy()
	for _, v := range supportedBy {
		if v.Size() == 1 {
			safeToDisintigrateBricks.Remove(v.Random())
		}
	}
	return &safeToDisintigrateBricks
}

func calculateChainReactions(bricks *sets.Set, supportedBy map[string]*sets.Set) int {
	safe := getSafeBricks(bricks, supportedBy)
	unsafe := bricks.Copy()
	unsafe.RemoveSet(*safe)
	chainReactions := 0
	for unsafe.Size() > 0 {
		b := unsafe.Random()
		unsafe.Remove(b)
		willFall := calculateChainReaction(b, supportedBy)
		chainReactions += willFall.Size()
	}
	return chainReactions
}

func calculateChainReaction(brick string, supportedBy map[string]*sets.Set) sets.Set {
	supportedCopy := map[string]*sets.Set{}
	for k, v := range supportedBy {
		s := v.Copy()
		supportedCopy[k] = &s
	}

	willFall := sets.New()
	queue := sets.New()
	queue.Add(brick)
	for queue.Size() > 0 {
		b := queue.Random()
		queue.Remove(b)
		willFall.Add(b)
		for k, v := range supportedCopy {
			if v.Size() == 1 && v.Has(b) {
				queue.Add(k)
			}
			if v.Has(b) {
				v.Remove(b)
			}
		}
	}

	willFall.Remove(brick)

	return willFall
}
