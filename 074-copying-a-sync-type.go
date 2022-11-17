package main

import "sync"

/*
 * Sync types must not be copied, otherwise they do not protect against data races.
 * In this particular example, this is because Increment has a value receiver which means Customer is copied, along with the mutex.
 * go run -race will detect a data race here, go vet will complain about it too.
 */

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{
		counters: map[string]int{},
	}
}

func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func TestCounter() {
	c := NewCounter()
	go func() {
		c.Increment("foo")
	}()
	go func() {
		c.Increment("bar")
	}()
}
