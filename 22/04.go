package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	l, r int
}

func isPairInBetween(p1, p2 Pair) bool {
	return p1.l >= p2.l && p1.l <= p2.r && p1.r >= p2.l && p1.r <= p2.r
}

func isPairOverlapping(p1, p2 Pair) bool {
	return (p1.l >= p2.l && p1.l <= p2.r) || (p1.r >= p2.l && p1.r <= p2.r)
}

func csvToPairs(csvText string) (Pair, Pair) {
	csv := strings.Split(csvText, ",")
	f, s := Pair{}, Pair{}
	fs, ss := strings.Split(csv[0], "-"), strings.Split(csv[1], "-")
	f.l, _ = strconv.Atoi(fs[0])
	f.r, _ = strconv.Atoi(fs[1])
	s.l, _ = strconv.Atoi(ss[0])
	s.r, _ = strconv.Atoi(ss[1])

	return f, s
}

func run04(comparator func(Pair, Pair) bool) {
	f := GetInputFile("./inputs/04.txt")

	sc := bufio.NewScanner(f)
	c := 0
	for sc.Scan() {
		f, s := csvToPairs(sc.Text())

		if comparator(f, s) || comparator(s, f) {
			c++
		}
	}

	fmt.Printf("result: %d\n", c)
}

func Run04P1() {
	run04(isPairInBetween)
}

func Run04P2() {
	run04(isPairOverlapping)
}
