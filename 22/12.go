package main

import (
	"bufio"
	"fmt"
)

type HillPos [2]int

type HillSolver struct {
	// we keep references
	nodes    map[HillPos]byte
	visited  map[HillPos]bool
	currDist map[HillPos]int
	start    HillPos
	end      HillPos
}

func (s *HillSolver) FindAllPaths() {
	bfsQueue := []HillPos{s.start}
	s.currDist[s.start] = 1

	for len(bfsQueue) > 0 {
		curr := bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		if s.isFinalPath(curr) {
			s.currDist[s.end] = s.currDist[curr] + 1
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

func Run12P1() {
	f := GetInputFile("./inputs/12.txt")
	sc := bufio.NewScanner(f)
	solver := HillSolver{
		nodes:    make(map[HillPos]byte),
		visited:  make(map[HillPos]bool),
		currDist: make(map[HillPos]int),
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

	solver.FindAllPaths()
	fmt.Printf("minimum path found: %d\n", solver.currDist[solver.end])
}
