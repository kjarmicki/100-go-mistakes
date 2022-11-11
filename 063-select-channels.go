package main

import "fmt"

/*
 * `select` statement picks any available (ready to read or ready to write) channel at random.
 * This can lead to surprising behavior when channels are buffered.
 */

func consumeChannel(done <-chan struct{}, ch <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			return
		}
	}
}

func consumeBufferedChannel(done <-chan struct{}, ch <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			for {
				select {
				case v := <-ch:
					fmt.Println(v)
				default:
					return
				}
			}
		}
	}
}

func SelectChannels() {
	ch := make(chan int, 10)
	done := make(chan struct{})

	// for buffered channel, this consumption is not deterministic
	// it might happen that both done and ch will be available
	// hence ending consumption prematurely
	go consumeChannel(done, ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- struct{}{}
	fmt.Println("----")

	ch = make(chan int, 10)
	done = make(chan struct{})

	// even for buffered channel, this consumption is deterministic
	// if it happens that both done and ch will be available, consumer will drain ch before returning
	go consumeBufferedChannel(done, ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- struct{}{}
}
