package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getInputFile() *os.File {
	f, err := os.Open("./inputs/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func Run01P1() {
	f := getInputFile()
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

func Run01P2() {
	f := getInputFile()
	defer f.Close()

	sc := bufio.NewScanner(f)
	i, sum := 0, 0
	fm, sm, tm := math.MinInt, math.MinInt, math.MinInt
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			sum = 0
			i++
		} else {
			v, _ := strconv.Atoi(t)
			sum += v
		}

		if sum > fm {
			tm = sm
			sm = fm
			fm = sum
		} else if sum > sm {
			tm = sm
			sm = sum
		} else if sum > tm {
			tm = sum
		}
	}

	fmt.Printf("max 3 calories: %d, %d, %d\n", fm, sm, tm)
	fmt.Printf("max 3 calories sum: %d\n", fm+sm+tm)
}
