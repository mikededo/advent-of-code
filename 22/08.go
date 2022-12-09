package main

import (
	"bufio"
	"fmt"
)

func fillMatrix(sc *bufio.Scanner) [][]int {
	f := make([][]int, 0)
	for sc.Scan() {
		t := sc.Text()
		r := make([]int, len(t))
		for i, c := range t {
			r[i] = int(c - '0')
		}
		f = append(f, r)
	}
	return f
}

func makeVisibleMap(rows, cols int) [][]bool {
	res := make([][]bool, rows)
	for i := 0; i < cols; i++ {
		res[i] = make([]bool, cols)
	}
	return res
}

func makeMaxsMaps(cols int) [4][]int {
	res := [4][]int{
		make([]int, cols),
		make([]int, cols),
		make([]int, cols),
		make([]int, cols),
	}
	for i := range res {
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	return res
}

func Run08P1() {
	f := GetInputFile("./inputs/08.txt")
	sc := bufio.NewScanner(f)
	mat := fillMatrix(sc)
	mm := makeMaxsMaps(len(mat[0]))
	vm := makeVisibleMap(len(mat), len(mat[0]))
	total := 0

	visit := func(r, c int, m *int) {
		if mat[r][c] > *m {
			*m = mat[r][c]
			if !vm[r][c] {
				vm[r][c] = true
				total++
			}
		}
	}

	for north := range mat {
		south := len(mat) - north - 1
		for east := range mat[north] {
			west := len(mat[north]) - east - 1
			visit(north, east, &mm[0][east])
			visit(north, east, &mm[1][north])
			visit(south, west, &mm[2][west])
			visit(south, west, &mm[3][south])
		}
	}

	fmt.Printf("visible trees: %d\n", total)
}
