package main

import (
	"bufio"
	"fmt"
	"math"
)

type HillPos [2]int

type HillSolver struct {
	// we keep references
	nodes    map[HillPos]byte
	visited  map[HillPos]bool
	currDist map[HillPos]int
	minDist  int
	start    HillPos
	end      HillPos
}

func (s *HillSolver) FindAllPaths(positions []HillPos) {
	for _, start := range positions {
		// restart everything
		bfsQueue := []HillPos{start}
		s.currDist = make(map[HillPos]int)
		s.visited = make(map[HillPos]bool)

		s.currDist[start] = 1

		for len(bfsQueue) > 0 {
			curr := bfsQueue[0]
			bfsQueue = bfsQueue[1:]

			if s.isFinalPath(curr) {
				dist := s.currDist[curr] + 1
				s.currDist[s.end] = dist
				if s.minDist > dist {
					s.minDist = dist
				}
				break
			}

			adjs := [4]HillPos{
				{curr[0], curr[1] + 1},
				{curr[0], curr[1] - 1},
				{curr[0] + 1, curr[1]},
				{curr[0] - 1, curr[1]},
			}
			for _, adj := range adjs {
				if s.isPossibleNextPath(curr, adj) {
					s.visited[adj] = true
					bfsQueue = append(bfsQueue, adj)
					s.currDist[adj] = s.currDist[curr] + 1
				}
			}
		}
	}
}

func (s *HillSolver) isFinalPath(pos HillPos) bool {
	return s.end[0] == pos[0] && s.end[1] == pos[1]
}

func (s *HillSolver) isPossibleNextPath(curr, next HillPos) bool {
	return !s.visited[next] && s.isValidNextByte(curr, next)
}

func (s *HillSolver) isValidNextByte(curr, next HillPos) bool {
	cv := s.nodes[curr]
	nv, exists := s.nodes[next]

	if cv == 'S' {
		return true
	} else if nv == 'E' && cv == 'z' {
		return true
	}

	return exists && int(nv)-int(cv) <= 1
}

func scanGrid(sc *bufio.Scanner) HillSolver {
	solver := HillSolver{
		nodes:    make(map[HillPos]byte),
		visited:  make(map[HillPos]bool),
		currDist: make(map[HillPos]int),
		minDist:  math.MaxInt,
	}

	i := 0
	for sc.Scan() {
		for j, b := range sc.Bytes() {
			pos := HillPos{-j, i}
			if b == 'S' {
				solver.start = pos
			} else if b == 'E' {
				solver.end = pos
			}
			solver.nodes[pos] = b
			solver.visited[pos] = false
		}
		i--
	}

	return solver
}

func Run12P1() {
	f := GetInputFile("./inputs/12.txt")
	sc := bufio.NewScanner(f)
	solver := scanGrid(sc)

	solver.FindAllPaths([]HillPos{solver.start})
	fmt.Printf("minimum path found: %d\n", solver.currDist[solver.end])
}

// 345
func Run12P2() {
	f := GetInputFile("./inputs/12.txt")
	sc := bufio.NewScanner(f)
	solver := scanGrid(sc)

	// find all a points
	points := make([]HillPos, 0)
	for k, v := range solver.nodes {
		if v == 'a' {
			points = append(points, k)
		}
	}

	solver.FindAllPaths(points)
	fmt.Printf("minimum path found of all 'a': %d\n", solver.minDist)
}
