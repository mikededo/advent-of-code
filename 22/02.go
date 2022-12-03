package main

import (
	"bufio"
	"fmt"
)

const (
	fr = "A"
	fp = "B"
	fs = "C"

	sr = "X"
	sp = "Y"
	ss = "Z"

	l = "X"
	d = "Y"
	w = "Z"
)

func getPoints(f, s string) int {
	switch f {
	case fr:
		switch s {
		case sr:
			return 4
		case sp:
			return 8
		case ss:
			return 3
		}
	case fp:
		switch s {
		case sr:
			return 1
		case sp:
			return 5
		case ss:
			return 9
		}
	default:
		// fs
		switch s {
		case sr:
			return 7
		case sp:
			return 2
		case ss:
			return 6
		}
	}

	// this should never happen
	return 0
}

func Run02P1() {
	f := GetInputFile("./inputs/02.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	sum := 0
	for sc.Scan() {
		t := sc.Text()
		var f, s string

		fmt.Sscanf(t, "%s %s", &f, &s)
		sum += getPoints(f, s)
	}

	fmt.Printf("your score is: %d\n", sum)
}

func Run02P2() {
	f := GetInputFile("./inputs/02.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	sum := 0
	for sc.Scan() {
		t := sc.Text()
		var f, s string

		fmt.Sscanf(t, "%s %s", &f, &s)
		switch f {
		case fr:
			switch s {
			case l:
				sum += getPoints(fr, ss)
			case d:
				sum += getPoints(fr, sr)
			case w:
				sum += getPoints(fr, sp)
			}
		case fp:
			switch s {
			case l:
				sum += getPoints(fp, sr)
			case d:
				sum += getPoints(fp, sp)
			case w:
				sum += getPoints(fp, ss)
			}
		default:
			// fs
			switch s {
			case l:
				sum += getPoints(fs, sp)
			case d:
				sum += getPoints(fs, ss)
			case w:
				sum += getPoints(fs, sr)
			}
		}
	}

	fmt.Printf("your score is: %d\n", sum)
}
