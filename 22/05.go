package main

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	SPACE_B         = 32
	OPEN_BRACKET_B  = 91
	CLOSE_BRACKET_B = 93
)

func reverseMatrix[S ~[][]E, E any](mtx [][]E) {
	for k := range mtx {
		s := mtx[k]
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
}

func makeMatrixOfSize(size int) [][]string {
	mtx := make([][]string, size)
	for i := range mtx {
		mtx[i] = make([]string, 0)
	}
	return mtx
}

func parseLineIntoSlice(line []byte, sliceSize int) []string {
	res := make([]string, sliceSize)
	pos := 0
	for i, b := range line {
		switch b {
		case SPACE_B:
			if i%4 == 3 {
				pos++
			}
		case OPEN_BRACKET_B:
		case CLOSE_BRACKET_B:
			continue
		default:
			res[pos] = string(b)
		}
	}
	return res
}

func addBytesToMatrix(mtx *[][]string, bts []byte) {
	size := len(bts)/4 + 1
	if len(*mtx) == 0 {
		*mtx = makeMatrixOfSize(size)
	}
	if len(bts) != 0 {
		for i, c := range parseLineIntoSlice(bts, size) {
			if c != "" {
				(*mtx)[i] = append((*mtx)[i], c)
			}
		}
	}
}

func removeNumberLineFromMatrix(mtx [][]string) {
	for i := range mtx {
		mtx[i] = mtx[i][:len(mtx[i])-1]
	}
}

func applyStackMoveToMatrix(mtx *[][]string, os, ds, q int) {
	for i := 0; i < q; i++ {
		l := len((*mtx)[os-1])
		(*mtx)[ds-1] = append((*mtx)[ds-1], (*mtx)[os-1][l-1])
		(*mtx)[os-1] = (*mtx)[os-1][:l-1]
	}
}

func applyListMoveToMatrix(mtx *[][]string, os, ds, q int) {
	osl := &(*mtx)[os-1]
	dsl := &(*mtx)[ds-1]
	*dsl = append(*dsl, (*osl)[len(*osl)-q:]...)
	*osl = (*osl)[:len(*osl)-q]
}

func matrixTopValues(mtx [][]string) []string {
	res := make([]string, 0)
	for i := range mtx {
		res = append(res, mtx[i][len(mtx[i])-1])
	}
	return res
}

func run05(cb func(*[][]string, int, int, int)) {
	f := GetInputFile("./inputs/05.txt")
	sc := bufio.NewScanner(f)
	parsingMtx := true

	var mtx [][]string
	for sc.Scan() {
		if parsingMtx {
			bts := sc.Bytes()
			addBytesToMatrix(&mtx, bts)
			if len(bts) == 0 {
				removeNumberLineFromMatrix(mtx)
				reverseMatrix[[][]string](mtx)
				parsingMtx = false
			}
			continue
		}

		var q, os, ds int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &q, &os, &ds)
		cb(&mtx, os, ds, q)
	}

	fmt.Printf("top strings are: %s\n", strings.Join(matrixTopValues(mtx), ""))
}

func Run05P1() {
	run05(applyStackMoveToMatrix)
}

func Run05P2() {
	run05(applyListMoveToMatrix)
}
