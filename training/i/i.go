package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"sort"
)

// //go:embed tests_I/01
// var input []byte

// //go:embed tests_I/01.a
// var output []byte

func main() {
	// in := bufio.NewReader(bytes.NewBuffer(input))
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(in, &n, &m)

	p := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	in.ReadLine()

	var res int64

	mm := make(map[int64]int64)
	for _, proc := range p {
		mm[proc] = 0
	}

	// fmt.Println(p)
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })
	// fmt.Println(p)
	var t, l int64
	for i := 0; i < m; i++ {
		fmt.Fscanln(in, &t, &l)
		for _, proc := range p {
			if v, ok := mm[proc]; ok && v <= t {
				mm[proc] = t + l
				res += proc * l

				// fmt.Println(res, mm)
				break
			}
		}

	}
	fmt.Println(res)
}
