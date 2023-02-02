package main

import (
	"bufio"
	"container/heap"
	_ "embed"
	"fmt"
	"os"
)

////go:embed tests_I/01
//var input []byte
//
////go:embed tests_I/01.a
//var output []byte

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Task struct {
	breakFree int64
	proc      int
}

type TasksHeap []Task

func (t TasksHeap) Len() int {
	return len(t)
}

func (t TasksHeap) Less(i, j int) bool {
	return t[i].breakFree < t[j].breakFree
}

func (t TasksHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *TasksHeap) Push(v any) {
	*t = append(*t, v.(Task))
}

func (t *TasksHeap) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[:n-1]
	return x
}

func main() {

	freeProc := &IntHeap{}
	heap.Init(freeProc)

	//	in := bufio.NewReader(bytes.NewBuffer(input))
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(in, &n, &m)

	var proc int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &proc)
		heap.Push(freeProc, proc)
	}
	in.ReadLine()

	tasks := &TasksHeap{}
	heap.Init(tasks)

	//    var task Task

	var res int64
	var t, l int64
	for i := 0; i < m; i++ {
		fmt.Fscanln(in, &t, &l)

		freeWorkProc(freeProc, tasks, t)

		if freeProc.Len() == 0 {
			continue
		}
		task := Task{
			breakFree: t + l,
			proc:      heap.Pop(freeProc).(int),
		}
		heap.Push(tasks, task)
		res += l * int64(task.proc)
	}
	fmt.Println(res)
}

func freeWorkProc(freeProc, tasks heap.Interface, t int64) {
	for {
		if tasks.Len() == 0 {
			break
		}
		task := heap.Pop(tasks).(Task)
		if task.breakFree <= t {
			heap.Push(freeProc, task.proc)
			continue
		} else {
			heap.Push(tasks, task)
			break
		}
	}
}
