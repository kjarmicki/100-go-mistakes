package main

// nil channels block:

func NilRead() {
	var ch chan int
	<-ch // blocks forever
}

func NilWrite() {
	var ch chan int
	ch <- 0 // also blocks forever
}

/*
 * Case: merge two int channels into one channel
 */

// this won't work: close(ch) is unreachable.
// if either ch1 or ch2 is closed, the loop will keep on reading zero value from that channel.
func mergeSimpleSelect(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

// this is better, but wasteful.
// if either ch1 or ch2 is closed, the loop will keep on checking the condition, wasting CPU resources.
func mergeFlags(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case v, open := <-ch1:
				if !open {
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()

	return ch
}

// optimal.
// when either ch1 or ch2 are closed, they're replaced with nil channel.
// this means that the loop is not going to be wasting resources by reading from this channel and wait for others instead.
func mergeNilChannels(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}
