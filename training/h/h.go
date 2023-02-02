package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
)

// //go:embed tests_H/24
// var input []byte

// //go:embed tests_H/24.a
// var output []byte

type position struct {
	row int
	col int
}

func main() {
	// in := bufio.NewReader(bytes.NewBuffer(input))
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		sourceMap := make(map[byte]map[position]struct{})
		var n, m int
		fmt.Fscanln(in, &n, &m)

		for k := 0; k < n; k++ {
			str, _, err := in.ReadLine()
			if err != nil {
				panic(err)
			}
			str = bytes.Trim(str, ".")
			s := bytes.Split(str, []byte("."))

			for j, v := range s {
				if _, ok := sourceMap[v[0]]; !ok {
					sourceMap[v[0]] = make(map[position]struct{})
				}
				sourceMap[v[0]][position{k, j}] = struct{}{}
			}
		}

		res := true
		for _, positions := range sourceMap {
			for pos := range positions {
				count := 0
				ref := len(positions)
				res = res && chain(positions, pos, &count, ref)
				break
			}
		}

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func chain(poss map[position]struct{}, pos position, count *int, ref int) bool {
	if ref == *count {
		return true
	}

	if len(poss) == 1 {
		*count++
		delete(poss, pos)
		return true
	}

	listCoords := allNeighbors(pos)
	res := false
	for _, v := range listCoords {
		delete(poss, pos)
		if _, ok := poss[v]; ok {
			res = res || chain(poss, v, count, ref)
			*count++
		}
	}

	return res
}

func allNeighbors(pos position) []position {
	l := position{row: pos.row, col: pos.col - 1}
	r := position{row: pos.row, col: pos.col + 1}
	u1 := position{row: pos.row - 1, col: pos.col}
	d1 := position{row: pos.row + 1, col: pos.col}
	var u2, d2 position
	if pos.row%2 == 0 {
		u2 = position{row: pos.row - 1, col: pos.col - 1}
		d2 = position{row: pos.row + 1, col: pos.col - 1}
	} else {
		u2 = position{row: pos.row - 1, col: pos.col + 1}
		d2 = position{row: pos.row + 1, col: pos.col + 1}
	}

	return []position{u2, u1, l, r, d2, d1}
}
