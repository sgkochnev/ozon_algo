package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
)

// //go:embed tests_D/09
// var input []byte

// //go:embed tests_D/09.a
// var output []byte

// func check(out io.Reader, table [][]int, n int) {
// 	var flag bool
// 	var flagRow bool
// 	for i, row := range table {
// 		for _, el := range row {
// 			var ref int
// 			fmt.Fscan(out, &ref)
// 			if el != ref {
// 				flagRow = true
// 			}
// 		}
// 		if flagRow {
// 			flag = flag || flagRow
// 			flagRow = false
// 			fmt.Printf("table - %d; row - %d;\n", n, i)
// 		}
// 		fmt.Fscanf(out, "\n")
// 	}
// 	fmt.Fscanf(out, "\n")
// 	if flag {
// 		printTable(table)
// 	}
// }

func main() {
	in := bufio.NewReader(os.Stdin)
	var countTable int
	fmt.Fscanf(in, "%d\n\n", &countTable)

	for k := 0; k < countTable; k++ {
		table := scanData(in)
		table = clickOnColumn(in, table)
		// check(out, table, k)
		printTable(table)
	}
}

func scanData(in io.Reader) [][]int {
	var n, m int
	fmt.Fscanln(in, &n, &m)
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &table[i][j])
		}
		fmt.Fscanln(in)
	}

	return table
}

func clickOnColumn(in io.Reader, table [][]int) [][]int {
	var countClick int
	var prevClick int
	fmt.Fscanln(in, &countClick)
	for i := 0; i < countClick; i++ {
		var click int
		fmt.Fscan(in, &click)
		if click == prevClick {
			continue
		}
		sortTable(&table, click-1)
		prevClick = click
	}
	fmt.Fscanf(in, "\n\n")
	return table
}

func printTable(table [][]int) {
	for _, row := range table {
		for _, el := range row {
			fmt.Printf("%d ", el)
		}
		fmt.Println()
	}
	fmt.Println()
}

type minimum struct {
	el  int
	idx int
}

func sortTable(tab *[][]int, idx int) {
	table := *tab
	t := make([][]int, len(table))
	m := make(map[int]struct{})
	for i := 0; i < len(table); i++ {
		min := minimum{el: 101, idx: 31}
		for j := 0; j < len(table); j++ {
			if _, ok := m[j]; ok {
				continue
			}
			if min.el > table[j][idx] {
				min.el = table[j][idx]
				min.idx = j
			}
		}
		t[i] = table[min.idx]
		m[min.idx] = struct{}{}
	}
	*tab = t
}
