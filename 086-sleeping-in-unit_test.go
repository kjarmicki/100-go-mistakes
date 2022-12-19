package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SleepHandler struct {
	n         int
	publisher sleepPublisher
}

type SleepFoo struct {
	value int
}

type sleepPublisher interface {
	Publish([]SleepFoo)
}

func getFoos(inputs []int) []SleepFoo {
	foos := make([]SleepFoo, 0, len(inputs))
	for _, input := range inputs {
		foos = append(foos, SleepFoo{
			value: input,
		})
	}
	return foos
}

func (h SleepHandler) getBestFoo(inputs []int) SleepFoo {
	foos := getFoos(inputs)
	best := foos[0]

	go func() {
		if len(foos) > h.n {
			foos = foos[:h.n]
		}
		h.publisher.Publish(foos)
	}()

	return best
}

/* testing */

type mutexPublisherMock struct {
	mu  sync.RWMutex
	got []SleepFoo
}

func (p *mutexPublisherMock) Publish(got []SleepFoo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.got = got
}

func (p *mutexPublisherMock) Get() []SleepFoo {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.got
}

type channelPublisherMock struct {
	ch chan []SleepFoo
}

func (p *channelPublisherMock) Publish(got []SleepFoo) {
	p.ch <- got
}

func retry(t *testing.T, retries int, wait time.Duration, assertion func(t *testing.T) bool) {
	emptyTesting := &testing.T{}
	for i := 0; i < retries-1; i++ {
		if assertion(emptyTesting) {
			assertion(t)
			return
		}
		time.Sleep(wait)
	}
	assertion(t)
}

func TestGetBestFooWithMutexPublisher(t *testing.T) {
	mock := mutexPublisherMock{}
	h := SleepHandler{
		publisher: &mock,
		n:         2,
	}

	foo := h.getBestFoo([]int{42})
	expected := SleepFoo{value: 42}
	assert.Equal(t, expected, foo)

	// this isn't bad, but it's also not entirely optimal
	// assertion will be retried, but won't necessarily happen as soon as the result is ready
	retry(t, 3, time.Microsecond*10, func(t *testing.T) bool {
		published := mock.Get()
		return assert.Equal(t, []SleepFoo{expected}, published)
	})

	// this is a bad idea, it makes the test flaky
	// there's no guarantee that published results will be ready after 10ms
	time.Sleep(10 * time.Millisecond)
	published := mock.Get()
	assert.Equal(t, []SleepFoo{expected}, published)
}

func TestGetBestFooWithChannelPublisher(t *testing.T) {
	mock := channelPublisherMock{
		ch: make(chan []SleepFoo),
	}
	h := SleepHandler{
		publisher: &mock,
		n:         2,
	}

	foo := h.getBestFoo([]int{42})
	expected := SleepFoo{value: 42}
	assert.Equal(t, expected, foo)

	// an optimal implementation
	// ready as soon as the result is ready
	published := <-mock.ch
	assert.Equal(t, []SleepFoo{expected}, published)
}
