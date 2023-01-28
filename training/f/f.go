package main

import (
	"bufio"
	"os"

	// _ "embed"
	"fmt"
	"io"
	"sort"
)

// //go:embed tests_F/01
// var input []byte

// //go:embed tests_F/01.a
// var output []byte

type myTime struct {
	hh uint8
	mm uint8
	ss uint8
}

func (t *myTime) isValid() bool {
	return t.hh < 24 && t.mm < 60 && t.ss < 60
}

func (t *myTime) lt(tt *myTime) bool {
	if t.hh < tt.hh {
		return true
	} else if t.hh == tt.hh && t.mm < tt.mm {
		return true
	} else if t.hh == tt.hh && t.mm == tt.mm && t.ss < tt.ss {
		return true
	}
	return false
}

type timeInterval struct {
	t1 myTime
	t2 myTime
}

func (ti *timeInterval) isCorrectInterval() bool {
	if ti.t1.hh < ti.t2.hh {
		return true
	} else if ti.t1.hh == ti.t2.hh && ti.t1.mm < ti.t2.mm {
		return true
	} else if ti.t1.hh == ti.t2.hh && ti.t1.mm == ti.t2.mm && ti.t1.ss <= ti.t2.ss {
		return true
	}
	return false
}

func main() {
	// in := bufio.NewReader(bytes.NewBuffer(input))

	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		PrintResult(f(in))
	}
}

func PrintResult(b bool) {
	if b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

type timeIntervals []timeInterval

func (tis timeIntervals) sort() {
	sort.Slice(tis, func(i, j int) bool {
		return tis[i].t1.lt(&tis[j].t1)
	})
}

func f(in io.Reader) bool {
	var n int
	fmt.Fscanln(in, &n)

	intervals := make(timeIntervals, 0, n)

	var t1, t2 myTime
	flag := true
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d:%d:%d-%d:%d:%d\n", &t1.hh, &t1.mm, &t1.ss, &t2.hh, &t2.mm, &t2.ss)
		if flag {
			newInterval := timeInterval{t1, t2}
			if !t1.isValid() || !t2.isValid() || !newInterval.isCorrectInterval() {
				flag = false
			}
			intervals = append(intervals, newInterval)

		}
	}
	if flag && len(intervals) >= 1 {
		intervals.sort()
		interval1 := intervals[0]
		for _, interval2 := range intervals[1:] {
			if !interval1.t2.lt(&interval2.t1) {
				return false
			}
			interval1 = interval2
		}
	}

	return flag
}
