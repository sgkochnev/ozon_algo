package main

import (
	"bufio"
	"bytes"
	"fmt"

	_ "embed"
)

//go:embed tests_J/09
var input []byte

//go:embed tests_J/09.a
var output []byte

type Map map[rune]*Node

type Node struct {
	data     rune
	children Map
	endWord  bool
}

func NewNode(val rune) *Node {
	var c Map
	return &Node{
		data:     val,
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

func (t *Tree) search(h *Node, word string) string {
	var resWords []string
	var res []rune

	var endChar rune
	var preEnd *Node
	for _, val := range word {
		endChar = val
		if h.children != nil {
			if node, ok := h.children[val]; ok {
				res = append(res, node.data)
				if node.endWord {
					resWords = append(resWords, string(res))
				}
			} else {
				break
			}
		} else {
			break
		}
		preEnd = h
		h, _ = h.Next(val)
	}

	// fmt.Println("THIS:", string(res), word)
	if string(res) == word && len(resWords) > 0 && resWords[len(resWords)-1] == word {
		if v, ok := h.children[endChar]; ok {
			if v.children == nil {
				for r, node := range h.children {
					if r != endChar {
						res = append(res, node.data)
						if node.endWord {
							resWords = append(resWords, string(res))
						}
						h = node
						break
					}
				}
			}
		}
		if 
	}

	for h.children != nil {
		if h.endWord && word != string(res) {
			break
		}
		for k, v := range h.children {
			res = append(res, k)
			h = v
			if v.endWord {
				resWords = append(resWords, string(res))
			}
			break
		}
	}

	if len(resWords) > 1 && resWords[len(resWords)-1] == word {
		return resWords[len(resWords)-2]
	}
	if string(res) == word {
		return ""
	}

	return string(res)
}

func (t *Tree) Search(word string) string {
	h := t.head
	res := t.search(h, word)
	// for res == "" || res == word {
	// 	for k, v := range h.children {
	// 		nres := t.search(v, word)
	// 		res = string(k) + nres
	// 	}
	// }
	if res == "" {
		var w []rune
		for r, node := range h.children {
			w = append(w, r)
			h = node
		}
		res = string(w)
	}

	return res
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
	in := bufio.NewReader(bytes.NewBuffer(input))
	out := bufio.NewReader(bytes.NewBuffer(output))
	// in := bufio.NewReader(os.Stdin)

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
		ol, _, _ := out.ReadLine()
		fmt.Println(Reverse(rhyme), string(ol), string(line))
		// fmt.Println(Reverse(rhyme))
	}
}
