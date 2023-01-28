package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		var count int
		fmt.Fscanln(in, &count)
		employees := make(map[int][]int)
		arr := make([]int, 0, count)
		for j := 0; j < count; j++ {
			var skill int
			fmt.Fscan(in, &skill)
			arr = append(arr, skill)
			employees[skill] = append(employees[skill], j+1)
		}
		fmt.Fscanln(in)

		PrintRes(arr, employees)
	}
}

type pair struct {
	first  int
	second int
}

func (p *pair) println() {
	fmt.Println(p.first, p.second)
}

func PrintRes(arr []int, employees map[int][]int) {
	pairs := make([]pair, 0, len(arr)/2)
	for i, v := range arr {
		if indexes, ok := employees[v]; ok && indexes[0] == i+1 {
			pairs = append(pairs, pair{first: indexes[0]})
			employees[v] = indexes[1:]
			if len(employees[v]) == 0 {
				delete(employees, v)
			}
		} else {
			continue
		}

		for j := 0; j < 100; j++ {
			if v+j > 100 && v-j < 0 {
				break
			}
			var (
				indexes1, indexes2 []int
				ok1, ok2           bool
			)
			if v+j <= 100 {
				indexes1, ok1 = employees[v+j]
			}
			if v-j >= 1 {
				indexes2, ok2 = employees[v-j]
			}

			if ok1 && ok2 && len(indexes1) > 0 && len(indexes2) > 0 {
				if indexes1[0] < indexes2[0] {
					f(len(pairs)-1, v+j, &pairs, employees)
				} else {
					f(len(pairs)-1, v-j, &pairs, employees)
				}
				break
			}

			if ok1 && len(indexes1) > 0 {
				f(len(pairs)-1, v+j, &pairs, employees)
				break
			}

			if ok2 && len(indexes2) > 0 {
				f(len(pairs)-1, v-j, &pairs, employees)
				break
			}
		}
	}

	for _, p := range pairs {
		p.println()
	}
	fmt.Println()
}

func f(i, v int, pairs *[]pair, employees map[int][]int) {
	(*pairs)[i].second, employees[v] = employees[v][0], employees[v][1:]
	if len(employees[v]) == 0 {
		delete(employees, v)
	}
}
