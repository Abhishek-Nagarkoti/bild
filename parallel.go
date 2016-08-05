package bild

import (
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func parallelize(size int, fn func(start, end int)) {
	procs := runtime.GOMAXPROCS(0)
	if procs <= 1 {
		fn(0, size)
	} else {
		var wg sync.WaitGroup
		counter := size
		partSize := size / procs

		for counter > 0 {
			start := counter - partSize
			end := counter
			if start < 0 {
				start = 0
			}
			counter -= partSize
			wg.Add(1)
			go func() {
				defer wg.Done()
				fn(start, end)
			}()
		}

		wg.Wait()
	}
}
