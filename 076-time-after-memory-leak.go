package main

import (
	"context"
	"fmt"
	"time"
)

func TimeAfterVsSleep() {
	<-time.After(time.Second)
	// is the same thing as
	time.Sleep(time.Second)
	// the only difference is in output: channel vs void
	// so it makes no sense to use time.After inline instead of time.Sleep
}

type Event struct{}

// there is a problem with this function
// resources associated with time.After (like the channel) won't be released until the associated time has passed
// if the time is really long and the loop is called many times this might mean significant amount of memory allocated
func consumerWithTimeAfter(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			fmt.Println(event)
		case <-time.After(time.Hour):
			fmt.Println("no messages received")
		}
	}
}

// this is somewhat better, but still not optimal
// creating a new context for each loop iteration isn't exactly lightweight
// also this feels like a hack, that's not the purpose of a context
func consumerWithConstext(ch <-chan Event) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel()
			fmt.Println(event)
		case <-ctx.Done():
			fmt.Println("no messages received")
		}
	}
}

// optimal solution
// uses time.NewTimer (which time.After also uses under the hood) and resets it for each loop iteration
// which means that there's minimal amount of new resources allocated with each iteration and none are lingering afterwards
func consumerWithTimer(ch <-chan Event) {
	timerDuration := time.Hour
	timer := time.NewTimer(timerDuration)

	for {
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			fmt.Println(event)
		case <-timer.C:
			fmt.Println("no messages received")
		}
	}
}
