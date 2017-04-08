package gorandom

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

const COUNT_BYTES int = 10000

func TestBytes_RandomBytes(t *testing.T) {
	results := make(map[string][]byte)
	for i := 0; i < COUNT_BYTES; i++ {
		b := RandomBytes(12)
		s := string(b)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = b
		assert.Equal(t, len(results), i+1)
	}
}

func TestBytes_RandomBytesByTime(t *testing.T) {
	assert.Equal(t, RandomBytesByTime(-1), EmptyBytes)
	assert.Equal(t, RandomBytesByTime(0), EmptyBytes)
	assert.Equal(t, RandomBytesByTime(3), EmptyBytes)
	assert.Equal(t, RandomBytesByTime(5), EmptyBytes)

	results := make(map[string][]byte, COUNT_BYTES)
	for i := 0; i < COUNT_BYTES; i++ {
		b := RandomBytesByTime(16)
		s := string(b)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = b
		assert.Equal(t, len(results), i+1)
	}
}
