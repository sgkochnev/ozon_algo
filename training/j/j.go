package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"sort"
)

// //go:embed tests_J/01
// var input []byte

// //go:embed tests_J/01.a
// var output []byte

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Dict []string

func (d *Dict) Add(word string) {
	*d = append(*d, word)
}

func (d *Dict) Sort() {
	sort.Slice(*d, func(i, j int) bool { return (*d)[i] < (*d)[j] })
}

func (d *Dict) BinSearch(word string) string {
	dict := *d

	m := 0
	l, r := 0, len(dict)-1
	for l < r {
		m = (l + r) / 2
		if dict[m] <= word {
			l = m + 1
		} else if dict[m] > word {
			r = m
		}
	}

	res := ""
	max := 0
	for _, v := range []int{-2, -1, 0} {
		if l+v >= 0 && l+v < len(dict) && dict[l+v] != word {
			if m := compare(dict[l+v], word); max <= m {
				max = m
				res = dict[l+v]
			}
		}
	}
	return res
}

func compare(str1, str2 string) int {
	count := 0
	s := []rune(str1)
	for i, v := range str2 {
		if !(len(s) > i && v == s[i]) {
			break
		}
		count++
	}
	return count
}

func main() {
	// in := bufio.NewReader(bytes.NewBuffer(input))
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(in, &n)

	dict := new(Dict)
	for i := 0; i < n; i++ {
		line, _, _ := in.ReadLine()
		word := Reverse(string(line))
		dict.Add(word)
	}
	dict.Sort()

	var q int
	fmt.Fscanln(in, &q)
	// out := bufio.NewReader(bytes.NewBuffer(output))
	for i := 0; i < q; i++ {
		line, _, _ := in.ReadLine()
		word := Reverse(string(line))
		word = Reverse(dict.BinSearch(word))
		// ol, _, _ := out.ReadLine()
		// fmt.Println(word, "\t", string(ol), "\t", string(line))
		fmt.Println(word)
	}
}
