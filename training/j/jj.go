// bad solution
package main

import (
	"bufio"
	"fmt"
	"os"

	_ "embed"
)

// //go:embed tests_J/14
// var input []byte

// //go:embed tests_J/14.a
// var output []byte

type Map map[rune]*Node

type Node struct {
	value    rune
	children Map
	endWord  bool
	parrent  *Node
}

func NewNode(val rune) *Node {
	var c Map
	return &Node{
		value:    val,
		children: c,
	}
}

func (n *Node) Add(val rune) *Node {
	if n.children == nil {
		n.children = make(Map)
	}
	if _, ok := n.children[val]; ok {
		return n.children[val]
	}
	node := NewNode(val)
	n.children[val] = node
	node.parrent = n
	return node
}

func (n *Node) Next(val rune) (*Node, bool) {
	v, ok := n.children[val]
	if ok {
		return v, ok
	}
	return n, !ok
}

type Tree struct {
	head *Node
}

func NewTree() *Tree {
	return &Tree{
		head: NewNode('\n'),
	}
}

func (t *Tree) Add(word string) {
	h := t.head
	for _, r := range word {
		h = h.Add(r)
	}
	h.endWord = true
}

func (t *Tree) search(h *Node, word string) []rune {
	var res []rune
	k := 0

	var f func(*Node, []rune) *Node

	f = func(h *Node, chars []rune) *Node {
		if len(chars) == 0 {
			if h.children != nil && string(res) == word {
				return h
			}

			for h.children == nil || len(h.children) <= 1 {
				h = h.parrent
				k++
				if h.endWord {
					break
				}
			}

			res = res[:len(res)-k]
			return h
		}

		char := chars[0]

		if h.children != nil {
			if node, ok := h.children[char]; ok {
				res = append(res, char)
				h = f(node, chars[1:])
			}
		}

		return h
	}

	h = f(h, []rune(word))

	if h.endWord {
		if string(res) != string(word) {
			return res
		}
		for val, node := range h.children {
			h = node
			res = append(res, val)
			break
		}
	}

	var char rune
	if k != 0 {
		char = []rune(word)[len(word)-k]
	} else {
		char = []rune(word)[0]
	}

	for val, node := range h.children {
		if val != char {
			h = node
			res = append(res, val)
			break
		}
	}

	for !h.endWord {
		for val, node := range h.children {
			h = node
			res = append(res, val)
			break
		}
	}

	return res
}
func (t *Tree) Search(word string) string {
	return string(t.search(t.head, word))
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	t := NewTree()
	// in := bufio.NewReader(bytes.NewBuffer(input))
	// out := bufio.NewReader(bytes.NewBuffer(output))
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		line, _, _ := in.ReadLine()
		word := Reverse(string(line))
		t.Add(word)
	}

	var q int
	fmt.Fscanln(in, &q)

	for i := 0; i < q; i++ {
		line, _, _ := in.ReadLine()
		word := Reverse(string(line))
		rhyme := t.Search(word)
		// ol, _, _ := out.ReadLine()
		// fmt.Println(Reverse(rhyme), "\t", string(ol), "\t", string(line))
		fmt.Println(Reverse(rhyme))
	}
}
