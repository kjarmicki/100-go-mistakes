package main

import (
	"fmt"
	"runtime"
)

/*
 * Be careful when slicing large arrays or slices.
 * In some cases the large backing array may still be kept in memory.
 */

func SliceLeakingCapacity() {
	get1000SlicesFromMessagesLeaking()
	get1000SlicesFromMessages()
}

// generate a message with the size of million bytes
func randomMessage() []byte {
	msg := make([]byte, 1_000_000)
	for i := 0; i < len(msg); i++ {
		msg[i] = byte(i)
	}
	return msg
}

// this is going to consume more memory than expected
func get1000SlicesFromMessagesLeaking() {
	slices := make([][]byte, 1000)
	defer func() {
		_ = slices
	}()
	for i := 0; i < 1000; i++ {
		msg := randomMessage()
		// here a slice of first 5 bytes is taken
		// but slices[i] has capacity of entire msg, not only 5!
		// hence the backing array is still referenced by capacity and takes up all the memory required by msg
		slices[i] = msg[:5]
	}
	runtime.GC()
	printAlloc() // about 1GB
}

// this is going to consume expected small amount of memory
func get1000SlicesFromMessages() {
	slices := make([][]byte, 1000)
	defer func() {
		_ = slices
	}()
	for i := 0; i < 1000; i++ {
		msg := randomMessage()
		// here slice is explicitly declared as having size of 5
		// copy sets slice's length and capacity at min(src, dst), so it's going to be 5
		// hence the backing array can be discarded and collected by GC
		slices[i] = make([]byte, 5)
		copy(slices[i], msg)
	}
	runtime.GC()
	printAlloc() // less than 1MB
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MB\n", m.Alloc/1024/1024)
}
