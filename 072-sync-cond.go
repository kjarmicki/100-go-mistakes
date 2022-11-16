package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Case: there is a donation object for accepting donations
 * and there are listener goroutines notifying when a particular goal was met
 */

// mutex-based solution
type donationMutex struct {
	mu      sync.RWMutex
	balance int
}

func DonationMutex() {
	donation := &donationMutex{}
	var wg sync.WaitGroup

	// inefficient - each listener goroutine keeps looping until its goal is met
	// which wastes a lot of CPU cycles
	listener := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
		wg.Done()
	}

	wg.Add(2)
	go listener(10)
	go listener(15)

	go func() {
		for {
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()

	wg.Wait()
}

// channel-based solution
type donationChannels struct {
	balance int
	ch      chan int
}

func DonationChannels() {
	donation := &donationChannels{ch: make(chan int)}

	listener := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("%d goal reached\n", donation.balance)
				return
			}
		}
	}

	// there is a problem with these listeners, since there is one producer they will be competing for values
	// so first one can get 1, 3, 4, 6, 8, 9, 11, 12, 14
	// and the second one 2, 5, 7, 10, 13, 15
	// so they might report higher goal than anticipated
	go listener(10)
	go listener(15)

	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
		if donation.balance == 15 {
			close(donation.ch)
			time.Sleep(time.Second)
			return
		}
	}
}

// cond-based implementation

type donationCond struct {
	cond    *sync.Cond
	balance int
}

func DonationCond() {
	donation := &donationCond{cond: sync.NewCond(&sync.Mutex{})}

	listener := func(goal int) {
		donation.cond.L.Lock() // start the critical section

		// this section means "wait until donation.balance < goal"
		// but unlike mutex solution it doesn't eat up CPU resources
		// and unlike channels solution can be observed by multiple goroutines
		for donation.balance < goal {
			donation.cond.Wait()
		}

		fmt.Printf("%d goal reached\n", donation.balance)
		donation.cond.L.Unlock() // end the critical section
	}

	go listener(10)
	go listener(15)

	for {
		time.Sleep(time.Second)

		// this section updates the state (condition) and then broadcasts the fact that it was updated
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}
