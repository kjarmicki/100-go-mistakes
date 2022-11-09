package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
 * Data races occur when multiple goroutines access the same memory location simultaneously
 * and at least one of them is writing.
 */

func race(verifyDeterminism func() bool) {
	for i := 0; i < 100_000; i++ {
		if !verifyDeterminism() {
			panic(fmt.Sprintf("kaboom, race at %d attempt\n", i))
		}
	}
	fmt.Println("no race detected")
}

// non-deterministic
// i++ does 3 things: read i, increment value, write new value
// so it could be that i is read in both goroutines before increment happen in either
// or it could be that i is read in one after increment already happened in another
// hence result is 1 or 2
func plainInc() int {
	i := 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		i++
		wg.Done()
	}()

	go func() {
		i++
		wg.Done()
	}()

	wg.Wait()
	return i
}

func PlainIncDataRace() {
	race(func() bool {
		return plainInc() == 2
	})
}

// deterministic
// atomic operations happen as a single unit
func atomicInc() int64 {
	var i int64
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		atomic.AddInt64(&i, 1)
		wg.Done()
	}()

	go func() {
		atomic.AddInt64(&i, 1)
		wg.Done()
	}()

	wg.Wait()
	return i
}

func AtomicIncDataRace() {
	race(func() bool {
		return atomicInc() == int64(2)
	})
}

// deterministic
// mutex locks a critical path so only one goroutine can be operating on this path at a time
func mutexInc() int {
	i := 0
	mutex := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
		wg.Done()
	}()

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
		wg.Done()
	}()

	wg.Wait()
	return i
}

func MutexIncDataRace() {
	race(func() bool {
		return mutexInc() == 2
	})
}

// deterministic
// increment happens in the main goroutine
func channelInc() int {
	i := 0
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch
	return i
}

func ChannelIncDataRace() {
	race(func() bool {
		return channelInc() == 2
	})
}
