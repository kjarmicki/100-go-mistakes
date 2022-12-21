package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

/*
 * Benchmarks are run for 1 second by default.
 * b.N value starts at 1 and is adjusted to match the time.
 */

// the skeleton of a benchmark is as follows:
func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
	/* Example output:
	 * cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	 * BenchmarkFoo-12    	 6394740	       158.0 ns/op	       0 B/op	       0 allocs/op
	 *                       ^ executions    ^ avg execution time
	 */
}

func foo() {
	if time.Now() == time.Now().Add(time.Hour) {
		fmt.Println("the universe is in a weird state")
	}
}

// mistake 1: not resetting or pausing the timer for expensive setup

// bad:
func BenchmarkWithSetupNoTimerReset(b *testing.B) {
	expensiveSetup()
	for i := 0; i < b.N; i++ {
		foo()
	}
}

// good:
func BenchmarkWithSetupTimerReset(b *testing.B) {
	expensiveSetup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo()
	}
}

// if the setup neeeds to be performed for each loop iteration:
/* commented out because it takes ages to run
func BenchmarkWithSetupTimerPause(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		expensiveSetup()
		b.StartTimer()
		foo()
	}
}
*/

func expensiveSetup() {
	time.Sleep(time.Second * 3)
}

// mistake 2: making wrong assumptions about micro-benchmarks

func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}

// ^ in these cases, a result can be affected by many things: machine activity, power management, caching etc.
// it's advisable to use -benchtime option to increase the time of the benchmark and thus lower the probability of inaccuracy

// mistake 3: not taking compiler optimizations into consideration

// this benchmark will yield result of about 0.30ns/op which is rougly one clock cycle.
// this is because once compiler inlines popcnt and realizes it has no side effects,
// the benchmark is effectively testing and empty loop
func BenchmarkPopcntInlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}

// one way to approach it is to cause a side effect, like global variable assignment
var global uint64

func BenchmarkPopcntSideEffects(b *testing.B) {
	var v uint64
	for i := 0; i < b.N; i++ {
		v = popcnt((uint64(i)))
	}
	// assignment to global is done once per benchmark instead of every time because it's expensive and can skew the results
	global = v
}

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}
