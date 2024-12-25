package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(24, 2024, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int64 {
	gates, wires, zBits := parseInput(input)

	completed := make([]bool, len(gates))
	zValues := make([]string, zBits)

	for slices.Contains(completed, false) {
		for i, gate := range gates {
			if completed[i] {
				continue
			}
			if _, ok := wires[gate.operand1]; !ok {
				continue
			}
			if _, ok := wires[gate.operand2]; !ok {
				continue
			}
			wires[gate.destinationWire] = gate.execute(wires[gate.operand1], wires[gate.operand2])
			completed[i] = true
			if gate.destinationWire[0] == 'z' {
				bit, _ := strconv.Atoi(gate.destinationWire[1:])
				zValues[bit] = strconv.Itoa(wires[gate.destinationWire])
			}
		}
	}

	z, _ := strconv.ParseInt(strings.Join(slices.Reverse(zValues), ""), 2, 64)
	return z
}

func solvePart2(input []string) string {
	// analyzed full ripple-carry adder logic to identify rules for wires and find violations
	// https://www.electronics-tutorials.ws/combination/comb_7.html
	// https://www.101computing.net/binary-additions-using-logic-gates/
	/*
		rules:
			wire1 xor wire2 -> zn (z wires only get xors) (three bad wires)
			notx and/or noty -> notz (can't xor to not a z without x and y) (three bad wires)
			x xor y -> wire, wire xor wire2 -> wire3 (must use x/y xor wires with another xor to carry) (one bad wire)
			wire1 and wire2 -> wire3, wire3 or wire4 -> wire5 (must use and/or to carry) (one bad wire)
	*/

	// could code these rules and run over the gates but regex + manual analysis is similar

	/* 
		sus instructions
		* x06 AND y06 -> z06
		06 xor -> sfm

		* gbd OR  fjv -> z13
		goes with x13 XOR y13 -> nmm

		* njc AND ngk -> z38
		goes with y38 AND x38 -> cpg

		* wvm OR dhs -> z45
		fine, last/most significant bit

		cjt XOR sfm -> jmq
		nmm XOR kwb -> gmh
		ngk XOR njc -> qrh
	*/

	/*
		xor outputs
		fht, kng, mqt, rvk, fns, jkn, gnj, qqd, cbd, qdt, mms, ptj, fgr, tsh, pvt
		jsp, njc, qsw, qkp, thb, dbv, fhk, hrh, dvj, kkg, wrj, hvf, nvr, mpv, hdk
		nmm, hch, cgj, ptw, hjc, tfn, jtk, ndr, gvt, wqn, sfm, tmh, rrt, z00, gts

		x25 XOR y25 -> cbd
		mrj OR cbd -> bjr
		missing xor
	*/

	/*
		and outputs
		bgc, dhs, qmd, svm, bch, cvm, mrj, fvk, dmk, bvb, djv, cjf, jgf, qdp, cbh
		z06, jpt, qkh, jgt, cpg, fcd, nsb, dfj, qtd, nqp, mmm, vfv, fmd, ftc, bwb
		rkw, chv, kqf, z38, rds, hct, bhs, pww, rbf, ghp, stn, cpb, dwn, gqg, vwp
		nbk, mwp, rjp, mkc, bbb, fjm, mfr, wwt, ckd, bnt, rjj, fjv, hnr, cqb, sbj
		gbd, wtb, swm, wdp, jrg, qqv, qpp, pbb, kjv, qmn, hrf, vmr, wgw, hmc, qqj
		mtn, rqf, fcb, gqf, kqm, msb, nmh, knk, mvf, smt, wvm, wjb, kgw, ntf

		kdt AND rqf -> mrj
		rqf XOR kdt -> z25
		x25 AND y25 -> rqf
		rqf missing or
	*/

	// no clue which pairs should be swapped but these are the invalid wires
	gates := []string{"cbd","gmh","jmq","qrh","rqf","z06","z13","z38"}
	sort.Slice(gates, func(i, j int) bool {
		return gates[i] < gates[j]
	})
	return strings.Join(gates, ",")
}

var (
	and string = "AND"
	or string = "OR"
	xor string = "XOR"
)
type gateFunction func(int, int) int

func andFunction(a, b int) int {
	return a & b
}

func orFunction(a, b int) int {
	return a | b
}

func xorFunction(a, b int) int {
	return a ^ b
}

type gate struct {
	operand1 string
	operand2 string
	execute gateFunction
	destinationWire string
}

func parseInput(input []string) ([]gate, map[string]int, int) {
	wireStrs := strings.Split(input[0], "\n")
	wires := map[string]int{}
	for _, wireStr := range wireStrs {
		wireName, wireValue := parseWire(wireStr)
		wires[wireName] = wireValue
	}

	zBits := 0
	gateStrs := strings.Split(input[1], "\n")
	gates := make([]gate, len(gateStrs))
	for i, gateStr := range gateStrs {
		gates[i] = parseGate(gateStr)
		if gates[i].destinationWire[0] == 'z' {
			zBits++
		}
	}

	return gates, wires, zBits
}

func parseWire(wire string) (string, int) {
	parts := strings.Split(wire, ": ")
	startingValue, _ := strconv.Atoi(parts[1])
	return parts[0], startingValue
}

func parseGate(gateStr string) gate {
	parts := strings.Split(gateStr, " ")
	var execute gateFunction
	switch parts[1] {
	case and:
		execute = andFunction
	case or:
		execute = orFunction
	case xor:
		execute = xorFunction
	}
	return gate{operand1: parts[0], operand2: parts[2], execute: execute, destinationWire: parts[4]}
}
