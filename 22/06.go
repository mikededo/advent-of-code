package main

import (
	"bufio"
	"fmt"
)

func areAllDifferent(b1, b2, b3, b4 byte) bool {
	return b1 != b2 && b1 != b3 && b1 != b4 && b2 != b3 && b2 != b4 && b3 != b4
}

func Run06P1() {
	f := GetInputFile("./inputs/06.txt")
	sc := bufio.NewScanner(f)

	sc.Scan()
	bts := sc.Bytes()
	fmt.Println(len(bts))
	i := 3
	for i < len(bts) {
		if areAllDifferent(bts[i-3], bts[i-2], bts[i-1], bts[i]) {
			break
		}
		i++
	}

	fmt.Printf("first marker after: %d\n", i+1)
}
