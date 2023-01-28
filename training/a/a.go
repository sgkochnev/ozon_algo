package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	var count int
	fmt.Fscanln(in, &count)

	var a, b int
	for i := 0; i < count; i++ {
		fmt.Fscanf(in, "%d %d\n", &a, &b)
		fmt.Println(Sum(a, b))
	}
}

func Sum(a, b int) int {
	return a + b
}
