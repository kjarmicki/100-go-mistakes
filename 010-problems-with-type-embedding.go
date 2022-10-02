package main

import (
	"io"
	"os"
	"sync"
)

type Foo struct {
	Bar // fields and methods of an embedded type are promoted - accessible within parent type
}

type Bar struct {
	Baz int
}

// embedding also works on interfaces
type MyReadWriter interface {
	io.Reader
	io.Writer
	~int
}

// misuse example
type InMem struct {
	sync.Mutex
	m map[string]int
}

func NewInMem() *InMem {
	return &InMem{
		m: make(map[string]int),
	}
}

func (i *InMem) Get(key string) (int, bool) {
	i.Lock()
	v, contains := i.m[key]
	i.Unlock()
	return v, contains
}

// correct usage example
type Logger struct {
	io.WriteCloser
}

func ProblemsWithTypeEmbedding() {
	m := NewInMem()
	m.Lock() // methods of sync.Mutex are promoted and therefore visible publicly, but they shouldn't be

	// this is correct usage - we want methods of WriteCloser to be exposed, we don't need to wrap and forward the calls
	log := Logger{WriteCloser: os.Stdout}
	log.Write([]byte("something happelend"))
	log.Close()
}

/*
 * Embedding vs OOP subclassing
 * is mostly about identity of the receiver of a method. Taking Logger as the example:
 * - with subclassing the receiver of .Write would be Logger
 * - with embedding the receiver of .Write is the WriteCloser (os.Stdout)
 * Embedding is more about composition than inheritance.
 */
