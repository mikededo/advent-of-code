package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Run10P1() {
	f := GetInputFile("./inputs/10.txt")
	sc := bufio.NewScanner(f)
	cycles := make([]int, 1)
	cycles[0] = 1

	for sc.Scan() {
		t := strings.Split(sc.Text(), " ")
		c := cycles[len(cycles)-1]
		if t[0] == "noop" {
			cycles = append(cycles, c)
			continue
		}
		v, _ := strconv.Atoi(t[1])
		cycles = append(cycles, c, c+v)
	}

	cc := func(pos int) int {
		return cycles[pos-1] * pos
	}

	fmt.Printf("sum of signal strengths is: %d\n", cc(20)+cc(60)+cc(100)+cc(140)+cc(180)+cc(220))
}
