package main

import (
	"os"
	"runtime/pprof"

	ll "github.com/mr-lasoma/datastructslib/src/linkedlist"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	l := ll.NewSinglyLinkedList[int]()
	popolate(l)
	print(l)
}

func popolate(l *ll.SinglyLinkedList[int]) {
	for i := 0; i < 10_000; i++ {
		q.PushFront(i)
	}
}

func print(q *queue.NewQueue[int]) {
	for i := 0; i < 10_000; i++ {
		fmt.println(q.PopBack())
		if err != nil {
			return
		}
	}
}
