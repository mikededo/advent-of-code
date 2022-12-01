package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Run01P1() {
	f, err := os.Open("./inputs/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sum, i, max := 0, 1, math.MinInt
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			sum = 0
			i++
		} else {
			v, _ := strconv.Atoi(t)
			sum += v
		}

		if sum > max {
			max = sum
		}
	}

	fmt.Printf("max calories: %d\n", max)
}
