package main

import (
	"context"
	"io"
	"os"
	"time"
)

/*
 * Deadline means that an activity should be stopped once time.Duration has passed or time.Time is now
 */

func publish(ctx context.Context, position int) error {
	for {
		select {
		// one case for processing
		case <-ctx.Done(): // context has timed out
			return ctx.Err()
		}
	}
}

func ContextDeadline() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // internally, context.WithTimeout creates a goroutine that will be retined in memory for 4 seconds or until cancel is called
	// so not calling it would mean that even if we return, we're leaving objects in memory
	_ = publish(ctx, 60)
}

/*
 * Cancellation means that an activity should be stopped once a signal is received
 */

func watcher(ctx context.Context, reader io.Reader) {
	for {
		select {
		// one case for processing
		case <-ctx.Done(): // context was cancelled
			return
		}
	}
}

func ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // when main returns, it calls cancel function to cancel the context passed to a watcher function

	go func() {
		file, _ := os.Open("./assets/sunset.jpeg")
		watcher(ctx, file)
	}()
}
