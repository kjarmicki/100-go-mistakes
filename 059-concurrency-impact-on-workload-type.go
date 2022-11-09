package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
 * Number of goroutines assigned for a task depends on type of that task.
 * If the workload is I/O-bound, it depends on the external system - how many concurrent accesses can the system cope with?
 * If the workload is CPU-bound, it's generally good idea to use as many OS threads as there are ones allocated to running goroutines.
 * The latter can be obtained with runtime.GOMAXPROCS(0)
 */

// simulate work
func task(b []byte) int {
	for i := 0; i < 10_000; i++ {
		if rand.Intn(100_000) == 42 {
			return 1
		}
	}
	return 0
}

func read(r io.Reader) (int, error) {
	var count int64
	var wg sync.WaitGroup
	n := runtime.GOMAXPROCS(0) // CPU-bound task: use GOMAXPROCS

	ch := make(chan []byte, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}
	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		ch <- b
	}
	close(ch)
	wg.Wait()
	return int(count), nil
}

func ReadSunset() {
	file, err := os.Open("./assets/sunset.jpeg")
	if err != nil {
		panic(err)
	}
	result, _ := read(file)
	fmt.Println(result)
}
