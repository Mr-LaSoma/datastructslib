package main

import (
	"os"
	"runtime/pprof"

	hm "github.com/mr-lasoma/datastructslib/src/hashmap"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	hm := hm.NewHashMap[int, int]()
	popolate(hm)
	print(hm)
}

func popolate(hm *hm.HashMap[int, int]) {
	for i := 0; i < 10_000; i++ {
		hm.Put(i, i)
	}
}

func print(hm *hm.HashMap[int, int]) {
	for i := 0; i < 10_000; i++ {
		_, _ = hm.Get(i)
	}
}
