package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func getCycles(sc *bufio.Scanner) []int {
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

	return cycles
}

func Run10P1() {
	f := GetInputFile("./inputs/10.txt")
	sc := bufio.NewScanner(f)

	cycles := getCycles(sc)
	cc := func(pos int) int {
		return cycles[pos-1] * pos
	}

	fmt.Printf("sum of signal strengths is: %d\n", cc(20)+cc(60)+cc(100)+cc(140)+cc(180)+cc(220))
}

func Run10P2() {
	f := GetInputFile("./inputs/10.txt")
	sc := bufio.NewScanner(f)

	cycles := getCycles(sc)
	sprite := make([]string, 0)
	for i, v := range cycles {
		mi := i % 40
		if mi == v || mi-1 == v || mi+1 == v {
			sprite = append(sprite, "#")
			continue
		}
		sprite = append(sprite, ".")
	}

	fmt.Println(sprite[0:40])
	fmt.Println(sprite[40:80])
	fmt.Println(sprite[80:120])
	fmt.Println(sprite[120:160])
	fmt.Println(sprite[160:200])
	fmt.Println(sprite[200:240])
}
