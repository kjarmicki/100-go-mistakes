package main

import "sync"

/*
 * A race condition occurs when the behavior depends on the sequence or
 * the timing of events that can't be controlled.
 */

// there is no data race here, the mutex is guarding access to i variable
// still, this code is not dererministic - there is no way of knowing whether the value returned will be 1 or 2
func raceCondition() int {
	i := 0
	mutex := sync.Mutex{}
	var wg sync.WaitGroup

	go func() {
		mutex.Lock()
		i = 1
		mutex.Unlock()
		wg.Done()
	}()

	go func() {
		mutex.Lock()
		i = 2
		mutex.Unlock()
		wg.Done()
	}()

	wg.Wait()
	return i
}

func RaceCondition() {
	race(func() bool {
		return raceCondition() == 2
	})
}
