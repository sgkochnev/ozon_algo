package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed tests_H/00
var input []byte

//go:embed tests_H/01.a
var output []byte

func main() {
	in := bufio.NewReader(bytes.NewBuffer(input))

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		// sourceMap := make(map[byte][]position)
		var n, m int
		fmt.Fscanln(in, &n, &m)

		var groups []map[byte]map[position]struct{}
		groups[0] = make(map[byte]map[position]struct{})

		for k := 0; k < n; k++ {
			str, _, err := in.ReadLine()
			if err != nil {
				panic(err)
			}
			str = bytes.Trim(str, ".")
			s := bytes.Split(str, []byte("."))

			for j, v := range s {
				// sourceMap[v[0]] = append(sourceMap[v[0]], position{k, j})
				
			}
		}

		// qwe := make(map[int][]int)
		// for _, v := range sourceMap[66] {
		// 	qwe[v.r] = append(qwe[v.r], v.c)
		// }
		// fmt.Println(qwe)

		// fmt.Println(sourceMap)
	}
}

type field struct {
	color byte
	position
}

type position struct {
	row int
	col int
}
