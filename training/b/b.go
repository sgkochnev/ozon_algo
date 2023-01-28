package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

type Goods = map[int]int

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		var count int
		fmt.Fscanln(in, &count)
		goods := make(Goods)
		for j := 0; j < count; j++ {
			var price int
			fmt.Fscan(in, &price)
			goods[price]++
		}
		fmt.Fscanln(in)
		fmt.Println(Result(goods))
	}
}

func Result(goods Goods) int {
	var res int
	for k, v := range goods {
		res += k * (goods[k] - (v / 3))
	}
	return res
}
