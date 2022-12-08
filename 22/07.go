package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type FileTreeManager struct {
	cdStack []*FileNode // this works as a stack, always has one element
}

type FileNode struct {
	childs map[string]*FileNode
	name   string
	size   int
}

func NewFileTreeManager(name string) *FileTreeManager {
	return &FileTreeManager{
		cdStack: []*FileNode{NewDirectoryNode(name)},
	}
}

func (m *FileTreeManager) CdIn(name string) {
	f, e := m.cdStack[len(m.cdStack)-1].FindFolderOfName(name)
	if e != nil {
		fmt.Println(e.Error())
		panic(-1)
	}
	m.cdStack = append(m.cdStack, f)
}

func (m *FileTreeManager) CdOut() {
	if len(m.cdStack) != 1 {
		m.cdStack = m.cdStack[:len(m.cdStack)-1]
	}
}

func (m *FileTreeManager) AddFolderToCurrent(name string) {
	c := m.cdStack[len(m.cdStack)-1]
	c.AddFolderNode(name)
}

func (m *FileTreeManager) AddFileToCurrent(name string, size int) {
	c := m.cdStack[len(m.cdStack)-1]
	c.AddFileNode(name, size)
}

func (m *FileTreeManager) TotalSumOfAtMost100k() int {
	r := make([]int, 0)
	m.cdStack[0].TotalSumOfAtMost100k(&r)
	return sumOfSlice(r)
}

func NewDirectoryNode(name string) *FileNode {
	return &FileNode{
		childs: make(map[string]*FileNode, 0),
		name:   name,
		size:   0,
	}
}

func NewFileNode(name string, size int) *FileNode {
	return &FileNode{
		childs: nil,
		name:   name,
		size:   size,
	}
}

func (n *FileNode) AddFolderNode(name string) {
	if n.childs[name] == nil {
		n.childs[name] = NewDirectoryNode(name)
	}
}

func (n *FileNode) AddFileNode(name string, size int) {
	if n.childs[name] == nil {
		n.childs[name] = NewFileNode(name, size)
	}
}

func (n *FileNode) FindFolderOfName(name string) (*FileNode, error) {
	cn := n.childs[name]
	if cn != nil && cn.name == name && cn.childs != nil {
		return cn, nil
	}

	return nil, fmt.Errorf("node with name %s not found in %s", name, n.name)
}

func (m *FileNode) TotalSumOfAtMost100k(r *[]int) int {
	if len(m.childs) == 0 {
		return m.size
	}

	// sum each directory & file
	sum := 0
	for _, v := range m.childs {
		res := v.TotalSumOfAtMost100k(r)
		sum += res
		// add only directories
		if len(v.childs) != 0 && res <= 100000 {
			*r = append(*r, res)
		}
	}

	return sum + m.size
}

func (n *FileNode) ToString() string {
	return n.toString(0)
}

func (n *FileNode) toString(lvl int) string {
	res := ""
	if lvl == 0 {
		res += fmt.Sprintf("- %s\n", n.name)
	} else {
		indent := strings.Repeat(" ", lvl)
		res += fmt.Sprintf("%s - %s\n", indent, n.name)
	}

	if len(n.childs) != 0 {
		for _, v := range n.childs {
			res += v.toString(lvl + 1)
		}
	}
	return res
}

func sumOfSlice(sl []int) int {
	sum := 0
	for _, n := range sl {
		sum += n
	}
	return sum
}

func Run07P1() {
	f := GetInputFile("./inputs/07.txt")
	sc := bufio.NewScanner(f)

	var tm *FileTreeManager = nil
	for sc.Scan() {
		t := strings.Split(sc.Text(), " ")
		if len(t) == 3 {
			// cd command
			f := t[len(t)-1]
			if tm == nil {
				tm = NewFileTreeManager(f)
			} else if f == ".." {
				tm.CdOut()
			} else {
				tm.CdIn(f)
			}
		} else if len(t) == 2 {
			// ls command or ls output
			if t[1] == "ls" {
				continue
			}

			if t[0] == "dir" {
				tm.AddFolderToCurrent(t[1])
			} else {
				sz, _ := strconv.Atoi(t[0])
				tm.AddFileToCurrent(t[1], sz)
			}
		}
	}

	fmt.Println(tm.cdStack[0].ToString())
	fmt.Println(tm.TotalSumOfAtMost100k())
}
