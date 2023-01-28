package main

import (
	"bufio"
	"io"
	"os"

	_ "embed"
	"fmt"
)

// //go:embed tests_E/35
// var input []byte

// //go:embed tests_E/35.a
// var output []byte

func main() {
	in := bufio.NewReader(os.Stdin)
	// in := bufio.NewReader(bytes.NewBuffer(input))

	var count int
	fmt.Fscanln(in, &count)

	for i := 0; i < count; i++ {
		flag := ReportVerification(in)
		PrintResult(flag)
	}
}

func PrintResult(b bool) {
	if b {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func ReportVerification(in io.Reader) bool {
	var countTasks int
	fmt.Fscanln(in, &countTasks)

	m := make(map[int]int)
	var taskN int
	var flag bool
	for i := 0; i < countTasks; i++ {
		fmt.Fscan(in, &taskN)
		if !flag {
			if v, ok := m[taskN]; ok && v+1 != i {
				flag = true
			}
			m[taskN] = i
		}
	}
	fmt.Fscanln(in)
	return flag
}
