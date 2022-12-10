package main

import (
	"bufio"
	"fmt"
	"math"
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

func northScore(m [][]int, col, pos int) int {
	score := 0
	threshold := m[pos][col]
	for i := pos - 1; i >= 0; i-- {
		score++
		if m[i][col] >= threshold {
			break
		}
	}
	return score
}

func southScore(m [][]int, col, pos int) int {
	score := 0
	threshold := m[pos][col]
	for i := pos + 1; i < len(m); i++ {
		score++
		if m[i][col] >= threshold {
			break
		}
	}
	return score
}

func eastScore(m [][]int, row, pos int) int {
	score := 0
	threshold := m[row][pos]
	for i := pos + 1; i < len(m); i++ {
		score++
		if m[row][i] >= threshold {
			break
		}
	}
	return score
}

func westScore(m [][]int, row, pos int) int {
	score := 0
	threshold := m[row][pos]
	for i := pos - 1; i >= 0; i-- {
		score++
		if m[row][i] >= threshold {
			break
		}
	}
	return score
}

func Run08P2() {
	f := GetInputFile("./inputs/08.txt")
	sc := bufio.NewScanner(f)
	mat := fillMatrix(sc)
	ms := math.MinInt

	for i := 1; i < len(mat)-1; i++ {
		for j := 1; j < len(mat[0])-1; j++ {
			sc := southScore(mat, j, i) * northScore(mat, j, i) * eastScore(mat, i, j) * westScore(mat, i, j)
			if sc > ms {
				ms = sc
			}
		}
	}
	fmt.Printf("max scenic score: %d\n", ms)
}
