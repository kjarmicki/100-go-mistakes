package main

import "context"

/*
 * Goroutines can hold up resources. Expose methods to close these resources once they're not needed anymore.
 */

type configWatcher struct {
	// holds some reources
}

func (cw *configWatcher) watch(ctx context.Context) {
	// starts watching the resource
}

func (cw *configWatcher) close() {
	// closes the resource handlers
}

func newWatcher() *configWatcher {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := &configWatcher{}
	go w.watch(ctx)
	return w
}

func StoppingGoroutines() {
	w := newWatcher()
	defer w.close() // close the resource handlers when the app is exiting

	// run the application
}
