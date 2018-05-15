package calculator

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

// go test .
func TestSum(t *testing.T) {
	n := add(5, 7)
	assert.Equal(t, n, 12)
}

// go test -bench .
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(5, 7)
	}
}

// go test -race .
func TestRaceDetector(t *testing.T) {
	race()
}

/*
CPU profiling:
	go test -run=XXX -bench=. -cpuprofile=cpu.out
	go tool pprof cpuprofile cpu.out

Tracing:
	go test -run=XXX -bench=. -trace=trace.out
	go tool trace trace.out

Memory profiling, network latency (Go 1.11) etc.
 */
func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// ...
	}
}
