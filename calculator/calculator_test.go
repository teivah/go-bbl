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
