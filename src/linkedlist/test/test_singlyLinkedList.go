package main

import (
	"fmt"
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
		l.PushFront(i)
	}
}

func print(l *ll.SinglyLinkedList[int]) {
	for i := 0; i < 10_000; i++ {
		v, err := l.PopBack()
		if err != nil {
			return
		}
		fmt.Println(v)
	}
}
