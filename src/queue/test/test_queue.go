package main

import (
	"os"
	"runtime/pprof"

	"github.com/mr-lasoma/datastructslib/src/queue"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	q := queue.NewQueue[int]()
	popolate(q)
	print(q)
}

func popolate(q *queue.NewQueue[int]) {
	for i := 0; i < 10_000; i++ {
		q.Enqueue(i)
	}
}

func print(q *queue.NewQueue[int]) {
	for i := 0; i < 10_000; i++ {
		fmt.println(q.Dequeue())
		if err != nil {
			return
		}
	}
}
