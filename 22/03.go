package main

import (
	"bufio"
	"fmt"
)

const (
	ASCII_LOW_UPPER  = 65
	ASCII_HIGH_UPPER = 90

	ASCII_LOW_LOWER  = 97
	ASCII_HIGH_LOWER = 122
)

func between(u, l int, b int) bool {
	return u <= b && b <= l
}

func byteSliceToScore(s []byte) []int {
	res := make([]int, 0)
	for _, b := range s {
		if between(ASCII_LOW_LOWER, ASCII_HIGH_LOWER, int(b)) {
			res = append(res, int(b)-ASCII_LOW_LOWER+1)
			continue
		}

		res = append(res, int(b)-ASCII_LOW_UPPER+27)
	}

	return res
}

func findEqualInTwoSlices(f, s []byte) int {
	var val int
	fm := make(map[int]bool)
	for _, v := range byteSliceToScore(f) {
		fm[v] = true
	}
	for _, v := range byteSliceToScore(s) {
		if fm[v] {
			val = v
			break
		}
	}

	return val
}

func findEqualInThreeSlices(f, s, t []byte) int {
	var val int
	fm := make(map[int]bool)
	sm := make(map[int]bool)

	for _, v := range byteSliceToScore(f) {
		fm[v] = true
	}
	for _, v := range byteSliceToScore(s) {
		if fm[v] {
			sm[v] = true
		}
	}
	for _, v := range byteSliceToScore(t) {
		if sm[v] {
			val = v
			break
		}
	}

	return val
}

func Run03P1() {
	f := GetInputFile("./inputs/03.txt")

	sc := bufio.NewScanner(f)
	sum := 0
	for sc.Scan() {
		t := sc.Text()

		fh := []byte(t[0 : len(t)/2])

		sh := []byte(t[(len(t) / 2):])
		sum += findEqualInTwoSlices(fh, sh)
	}

	fmt.Printf("sum of priorities: %d\n", sum)
}

func Run03P2() {
	f := GetInputFile("./inputs/03.txt")

	sc := bufio.NewScanner(f)
	sum, g := 0, 1
	var fh, sh, th []byte
	for sc.Scan() {
		t := sc.Text()

		switch g {
		case 1:
			fh = []byte(t)
		case 2:
			sh = []byte(t)
		case 3:
			th = []byte(t)
			sum += findEqualInThreeSlices(fh, sh, th)
			g = 0
		}

		g++
	}

	fmt.Printf("sum of priorities: %d\n", sum)
}
