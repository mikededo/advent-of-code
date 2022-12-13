package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Pos [2]float64

func posDistance(p1, p2 Pos) float64 {
	return math.Sqrt(math.Pow(p1[0]-p2[0], 2) + math.Pow(p1[1]-p2[1], 2))
}

func Run09P1() {
	f := GetInputFile("./inputs/09.txt")
	sc := bufio.NewScanner(f)

	visited := make(map[Pos]bool)
	h, ph := Pos{0, 0}, Pos{0, 0}
	for sc.Scan() {
		t := strings.Split(sc.Text(), " ")
		c, _ := strconv.Atoi(t[1])
		for i := 0; i < c; i++ {
			switch t[0] {
			case "R":
				if posDistance(Pos{h[0] + 1, h[1]}, ph) >= 2 {
					ph[0] = h[0]
					ph[1] = h[1]
				}
				h[0]++
			case "U":
				if posDistance(Pos{h[0], h[1] + 1}, ph) >= 2 {
					ph[1] = h[1]
					ph[0] = h[0]
				}
				h[1]++
			case "D":
				if posDistance(Pos{h[0], h[1] - 1}, ph) >= 2 {
					ph[1] = h[1]
					ph[0] = h[0]
				}
				h[1]--
			default:
				if posDistance(Pos{h[0] - 1, h[1]}, ph) >= 2 {
					ph[0] = h[0]
					ph[1] = h[1]
				}
				h[0]--
			}

			visited[ph] = true
		}
	}

	fmt.Printf("total visited spots: %d\n", len(visited))
}
