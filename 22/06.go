package main

import (
	"bufio"
	"fmt"
)

var void struct{}

func areFourDifferent(b1, b2, b3, b4 byte) bool {
	return b1 != b2 && b1 != b3 && b1 != b4 && b2 != b3 && b2 != b4 && b3 != b4
}

func areAllDifferent(sl []byte) bool {
	s := make(map[byte]struct{}, 0)
	for _, e := range sl {
		s[e] = void
	}
	return len(sl) == len(s)
}

func Run06P1() {
	f := GetInputFile("./inputs/06.txt")
	sc := bufio.NewScanner(f)

	sc.Scan()
	bts := sc.Bytes()
	i := 3
	for i < len(bts) {
		if areFourDifferent(bts[i-3], bts[i-2], bts[i-1], bts[i]) {
			break
		}
		i++
	}

	fmt.Printf("first marker after: %d\n", i+1)
}

func Run06P2() {
	f := GetInputFile("./inputs/06.txt")
	sc := bufio.NewScanner(f)

	sc.Scan()
	bts := sc.Bytes()
	i := 14
	for i < len(bts) {
		if areAllDifferent(bts[i-14 : i]) {
			break
		}
		i++
	}

	fmt.Printf("first message marker after: %d\n", i)
}
