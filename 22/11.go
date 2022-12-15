package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type MonkeyInfo [6]string

type Monkey struct {
	id              int
	inspectCount    int
	divisor         int
	divisorsProduct int
	items           []int
	op              func(int) int
	test            func(int) bool
	trueTo          int
	falseTo         int
}

func (m *Monkey) Operate(part int) (int, int) {
	if len(m.items) == 0 {
		return -1, -1
	}

	m.inspectCount++
	item := m.items[0]
	m.items = m.items[1:]

	var v int
	if part == 1 {
		v = int(float64(m.op(item) / 3))
	} else {
		v = m.op(item) % m.divisorsProduct
	}

	if m.test(v) {
		return m.trueTo, v
	}
	return m.falseTo, v
}

func (m *Monkey) Receive(v int) {
	m.items = append(m.items, v)
}

var split func(string, string) []string

func atoi(v string) int {
	n, _ := strconv.Atoi(v)
	return n
}

func buildMonkeyOperation(v1, op, v2 string) func(int) int {
	return func(old int) int {
		var n1, n2 int
		if v1 == "old" {
			n1 = old
		} else {
			n1 = atoi(v1)
		}
		if v2 == "old" {
			n2 = old
		} else {
			n2 = atoi(v2)
		}

		switch op {
		case "*":
			return n1 * n2
		default:
			// addition
			return n1 + n2
		}
	}
}

func buildMonkeyTest(v int) func(int) bool {
	return func(i int) bool {
		return i%v == 0
	}
}

func runMonkeyBussiness(monkeys []Monkey, part, ite int) {
	for i := 0; i < ite; i++ {
		for j := 0; j < len(monkeys); j++ {
			for {
				to, v := monkeys[j].Operate(part)
				if to == -1 {
					break
				}
				monkeys[to].Receive(v)
			}
		}
	}
}

func calculateMaxBussiness(monkeys []Monkey) (int, int) {
	maxIns1, maxIns2 := math.MinInt, math.MinInt
	for _, mm := range monkeys {
		if mm.inspectCount > maxIns1 {
			maxIns2 = maxIns1
			maxIns1 = mm.inspectCount
		} else if mm.inspectCount > maxIns2 {
			maxIns2 = mm.inspectCount
		}
	}
	return maxIns1, maxIns2
}

func parseMonkey(lines MonkeyInfo) Monkey {
	m := Monkey{}
	for i, l := range lines {
		if l == "" {
			break
		}

		switch i {
		case 0:
			// Monkey id
			idStr := split(l, " ")[1]
			m.id = atoi(idStr[:len(idStr)-1])
		case 1:
			// Starting items
			m.items = make([]int, 0)
			itemsStr := split(strings.Split(l, ": ")[1], ", ")
			for _, v := range itemsStr {
				m.items = append(m.items, atoi(v))
			}
		case 2:
			// Operation
			op := split(split(split(l, ": ")[1], "= ")[1], " ")
			m.op = buildMonkeyOperation(op[0], op[1], op[2])
		case 3:
			// Divisible by
			test := split(l, "by ")[1]
			div := atoi(test)
			m.divisor = div
			m.test = buildMonkeyTest(div)
		case 4:
			// if true
			tm := split(l, " ")
			m.trueTo = atoi(tm[len(tm)-1])
		case 5:
			// if false
			tm := split(l, " ")
			m.falseTo = atoi(tm[len(tm)-1])
		}
		i++
	}

	return m
}

func scanMonkeys(sc *bufio.Scanner, part int) []Monkey {
	monkeys := make([]Monkey, 0)
	for sc.Scan() {
		lines := MonkeyInfo{}
		lines[0] = sc.Text()
		sc.Scan()
		lines[1] = sc.Text()
		sc.Scan()
		lines[2] = sc.Text()
		sc.Scan()
		lines[3] = sc.Text()
		sc.Scan()
		lines[4] = sc.Text()
		sc.Scan()
		lines[5] = sc.Text()

		monkeys = append(monkeys, parseMonkey(lines))

		// Skip empty line
		sc.Scan()
	}

	if part == 2 {
		divisiorsProduct := 1
		for _, m := range monkeys {
			divisiorsProduct *= m.divisor
		}
		for i := range monkeys {
			monkeys[i].divisorsProduct = divisiorsProduct
		}
	}

	return monkeys
}

func Run11P1() {
	f := GetInputFile("./inputs/11.txt")
	sc := bufio.NewScanner(f)
	split = strings.Split
	monkeys := scanMonkeys(sc, 1)

	runMonkeyBussiness(monkeys, 1, 20)

	maxIns1, maxIns2 := calculateMaxBussiness(monkeys)
	fmt.Printf("monkey bussiness value: %d\n", maxIns1*maxIns2)
}

func Run11P2() {
	f := GetInputFile("./inputs/11.txt")
	sc := bufio.NewScanner(f)
	split = strings.Split
	monkeys := scanMonkeys(sc, 2)

	runMonkeyBussiness(monkeys, 2, 10000)

	maxIns1, maxIns2 := calculateMaxBussiness(monkeys)
	fmt.Printf("monkey bussiness value: %d\n", maxIns1*maxIns2)
}
